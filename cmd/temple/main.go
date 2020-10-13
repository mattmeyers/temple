package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	htmpl "html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	ttmpl "text/template"

	"github.com/fsnotify/fsnotify"
	"github.com/mattmeyers/temple"
)

func main() {
	watch()
	flag.Usage = usage
	flagHTML := flag.Bool("html", false, "use html/template for template parsing")
	flagWatch := flag.Bool("watch", false, "watch input files for changes")
	flagData := flag.String("d", "", "a JSON file containing the template data")
	flagOutput := flag.String("o", "", "the output filename")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("temple: at least one input file required")
		os.Exit(1)
	}

	data, err := readDataFile(*flagData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w, err := getWriter(*flagOutput)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer w.Close()

	var f func([]string, interface{}, io.Writer) error
	if *flagHTML {
		f = parseHTML
	} else {
		f = parseText
	}

	if *flagWatch {
		fmt.Println("file watching not implement, try again later")
		os.Exit(1)
	} else {
		err = f(flag.Args(), data, w)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	files := []string{"testdata/template1.tmpl", "testdata/template2.tmpl"}
	go func() {
		checksumCache := map[string]string{}
		outfile, err := os.Create("out.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer outfile.Close()

	loop:
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				for _, f := range files {
					if f != event.Name {
						continue
					}
					checksum, err := md5sum(f)
					if err != nil {
						log.Println("error:", err)
						continue loop
					}
					fmt.Println(checksum)

					if checksumCache[f] != checksum {
						fmt.Println("Parsing and printing")
						err = parseText(files, nil, outfile)
						if err != nil {
							fmt.Println("error:", err)
							continue loop
						}
						if checksum != "" {
							checksumCache[f] = checksum
						}
						outfile.Sync()

						_, err = outfile.Seek(0, 0)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	for _, f := range files {
		err = watcher.Add(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
}

func md5sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [Options] template1 template2...\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Compile Go templates from the command line\n\n")
	flag.PrintDefaults()
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

// passthrough represents a pipeline function that does nothing.
// Parsing templates containing functions not defined in the
// tmeple library will cause parsing to fail. This function can
// be used in the FuncMap to nullify these functions.
//
// Note: Since this is removes pipeline functions, the resulting
// output will likely have the incorrect data inserted if it
// compiles at all.
var passthrough = func(args ...interface{}) interface{} {
	for _, a := range args {
		fmt.Printf("%v: %T\n", a, a)
	}
	if len(args) == 0 {
		return nil
	}
	return args[len(args)-1]
}

func parseHTML(infiles []string, data interface{}, w io.Writer) error {
	t, err := htmpl.New(infiles[0]).Funcs(temple.FullFuncMap().HTML()).ParseFiles(infiles...)
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
