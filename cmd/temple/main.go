package main

import (
	"encoding/json"
	"flag"
	"fmt"
	htmpl "html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	ttmpl "text/template"

	"github.com/fsnotify/fsnotify"
	"github.com/mattmeyers/temple"
)

type app struct {
	l *logger
}

type parseFunc func([]string, interface{}, io.Writer) error

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... <BASE TEMPLATE> [TEMPLATE]...\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Compile Go templates from the command line\n\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flagHTML := flag.Bool("html", false, "use html/template for template parsing")
	flagWatch := flag.Bool("w", false, "watch input files for changes")
	flagVerbose := flag.Bool("v", false, "show extra log info")
	flagData := flag.String("d", "", "a JSON file containing the template data")
	flagOutput := flag.String("o", "", "the output filename")
	flag.Parse()

	app := &app{l: newLogger(*flagVerbose)}

	if flag.NArg() == 0 {
		app.l.Fatal("temple: at least one input file required")
	}

	var f parseFunc
	if *flagHTML {
		f = parseHTML
	} else {
		f = parseText
	}

	if *flagWatch {
		app.watch(f, files{templates: flag.Args(), data: *flagData, outfile: *flagOutput})
		return
	}

	data, err := readDataFile(*flagData)
	if err != nil {
		app.l.Fatal(err.Error())
	}

	w, err := getWriter(*flagOutput)
	if err != nil {
		app.l.Fatal(err.Error())
	}
	defer w.Close()

	err = f(flag.Args(), data, w)
	if err != nil {
		app.l.Fatal(err.Error())
	}
}

type files struct {
	templates []string
	data      string
	outfile   string
}

func (a *app) watch(parse parseFunc, fs files) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		a.l.Fatal(err.Error())
	}
	defer watcher.Close()

	go func() {
		var err error
		data, err := readDataFile(fs.data)
		if err != nil {
			a.l.Fatal("error reading data file: %v", err)
		}

		err = a.update(parse, fs, data)
		if err != nil {
			a.l.Error(err.Error())
		}

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Write != fsnotify.Write {
					continue
				}

				if event.Name == fs.data {
					data, err = readDataFile(fs.data)
					if err != nil {
						a.l.Error("error reading data file: %v", err)
						continue
					}
				}

				a.l.Debug("Detected change in %s, rebuilding...\n", event.Name)
				err = a.update(parse, fs, data)
				if err != nil {
					a.l.Error(err.Error())
				} else {
					a.l.Debug("Successful rebuild!")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				a.l.Error("error while watching: %v", err)
			}
		}
	}()

	for _, f := range fs.templates {
		a.l.Info("Watching %s for changes...\n", f)
		err = watcher.Add(f)
		if err != nil {
			a.l.Fatal(err.Error())
		}
	}

	if fs.data != "" {
		a.l.Info("Watching %s for changes...\n", fs.data)
		err = watcher.Add(fs.data)
		if err != nil {
			a.l.Fatal(err.Error())
		}
	}

	<-(make(chan struct{}))
}

func (a *app) update(parse parseFunc, fs files, data interface{}) error {
	var err error
	var f *os.File
	if fs.outfile == "" {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(fs.outfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			a.l.Fatal("error opening outfile: %v", err)
		}
	}

	// We don't immediately check for this error. Even if an error occurs,
	// execution can continue. We'll return the error and log it.
	err = parse(fs.templates, data, f)

	if f != os.Stdout {
		err = f.Close()
		if err != nil {
			a.l.Fatal("error closing outfile: %v", err)
		}
	}

	return err
}

func readDataFile(filename string) (interface{}, error) {
	if filename == "" {
		return nil, nil
	}

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(f, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getWriter(filename string) (io.WriteCloser, error) {
	if filename == "" {
		return os.Stdout, nil
	}

	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func parseHTML(infiles []string, data interface{}, w io.Writer) error {
	_, name := filepath.Split(infiles[0])

	t, err := htmpl.New(name).Funcs(temple.FullFuncMap().HTML()).ParseFiles(infiles...)
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func parseText(infiles []string, data interface{}, w io.Writer) error {
	_, name := filepath.Split(infiles[0])

	t, err := ttmpl.New(name).Funcs(temple.FullFuncMap().Text()).ParseFiles(infiles...)
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
