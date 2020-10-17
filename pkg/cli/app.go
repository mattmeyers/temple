package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	htemplate "html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	ttemplate "text/template"

	"github.com/fsnotify/fsnotify"
	"github.com/mattmeyers/temple"
)

// App represents the temple application. To populate the struct with values
// from the command line as described in the documentation, instantiate the
// struct with the New() function.
type App struct {
	Templates  []string
	DataFile   string
	OutputFile string

	HTMLFuncMap temple.FuncMap
	TextFuncMap temple.FuncMap

	HTML  bool
	Watch bool

	logger *logger
}

func usage() {
	fmt.Fprint(os.Stderr, "Name:\n\ttemple - compile Go templates from the command line\n\n")
	fmt.Fprint(os.Stderr, "Usage:\n\ttemple [OPTION]... <BASE TEMPLATE> [TEMPLATE]...\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
}

// New creates a new App. The default values are populated with values from the
// command line. The list of command line args will be used as the list of
// template files. At least one must be provided. The available command line
// flags are:
//	 -o string: The output filename
//	 -d string: The data file
//	 -w: 		Indicates that the input files should be watched for changes
//	 -v: 		Indicates that additional logging should be displayed
//	 -html:		Indicates that the html/template parser should be used
func New() *App {
	flag.Usage = usage
	flagHTML := flag.Bool("html", false, "use html/template for template parsing")
	flagWatch := flag.Bool("w", false, "watch input files for changes")
	flagVerbose := flag.Bool("v", false, "show extra log info")
	flagData := flag.String("d", "", "a JSON file containing the template data")
	flagOutput := flag.String("o", "", "the output filename")
	flag.Parse()

	return &App{
		Templates:   flag.Args(),
		DataFile:    *flagData,
		OutputFile:  *flagOutput,
		TextFuncMap: make(temple.FuncMap),
		HTMLFuncMap: make(temple.FuncMap),
		Watch:       *flagWatch,
		HTML:        *flagHTML,
		logger:      newLogger(*flagVerbose),
	}
}

// WithFuncMap merges the provided FuncMap into the App's text and HTML FuncMaps.
// This method can be called multiple times to merge in multiple FuncMaps.
func (a *App) WithFuncMap(f temple.FuncMap) *App {
	a.HTMLFuncMap = temple.MergeFuncMaps(a.HTMLFuncMap, f)
	a.TextFuncMap = temple.MergeFuncMaps(a.TextFuncMap, f)
	return a
}

// ClearFuncMaps resets both the text and HTML FuncMaps for the App.
func (a *App) ClearFuncMaps() *App {
	a.HTMLFuncMap.Clear()
	a.TextFuncMap.Clear()
	return a
}

// WithHTMLFuncMap merges the provided FuncMap into the App's HTML FuncMap.
// This method can me called multipled times to merge in multiple FuncMaps.
func (a *App) WithHTMLFuncMap(f temple.FuncMap) *App {
	a.HTMLFuncMap = temple.MergeFuncMaps(a.HTMLFuncMap, f)
	return a
}

// ClearHTMLFuncMap resets the App's HTML FuncMap.
func (a *App) ClearHTMLFuncMap() *App {
	a.HTMLFuncMap.Clear()
	return a
}

// WithTextFuncMap merges the provided FuncMap into the App's text FuncMap.
// This method can me called multipled times to merge in multiple FuncMaps.
func (a *App) WithTextFuncMap(f temple.FuncMap) *App {
	a.TextFuncMap = temple.MergeFuncMaps(a.TextFuncMap, f)
	return a
}

// ClearTextFuncMap resets the App's text FuncMap.
func (a *App) ClearTextFuncMap() *App {
	a.TextFuncMap.Clear()
	return a
}

// Run runs the temple application. If App.Watch is set to true, then this
// method will never return.
func (a *App) Run() error {

	if flag.NArg() == 0 {
		a.logger.Fatal("temple: at least one input file required")
	}

	var f parseFunc
	if a.HTML {
		f = a.parseHTML
	} else {
		f = a.parseText
	}

	if a.Watch {
		a.watch(f)
		return nil
	}

	data, err := readDataFile(a.DataFile)
	if err != nil {
		return err
	}

	w, err := getWriter(a.OutputFile)
	if err != nil {
		return err
	}
	defer w.Close()

	err = f(flag.Args(), data, w)
	if err != nil {
		return err
	}

	return nil
}

type parseFunc func([]string, interface{}, io.Writer) error

func (a *App) watch(parse parseFunc) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		a.logger.Fatal(err.Error())
	}
	defer watcher.Close()

	go func() {
		var err error
		data, err := readDataFile(a.DataFile)
		if err != nil {
			a.logger.Fatal("error reading data file: %v", err)
		}

		err = a.update(parse, data)
		if err != nil {
			a.logger.Error(err.Error())
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

				if event.Name == a.DataFile {
					data, err = readDataFile(a.DataFile)
					if err != nil {
						a.logger.Error("error reading data file: %v", err)
						continue
					}
				}

				a.logger.Debug("Detected change in %s, rebuilding...\n", event.Name)
				err = a.update(parse, data)
				if err != nil {
					a.logger.Error(err.Error())
				} else {
					a.logger.Debug("Successful rebuild!")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				a.logger.Error("error while watching: %v", err)
			}
		}
	}()

	for _, f := range a.Templates {
		a.logger.Info("Watching %s for changes...\n", f)
		err = watcher.Add(f)
		if err != nil {
			a.logger.Fatal(err.Error())
		}
	}

	if a.DataFile != "" {
		a.logger.Info("Watching %s for changes...\n", a.DataFile)
		err = watcher.Add(a.DataFile)
		if err != nil {
			a.logger.Fatal(err.Error())
		}
	}

	<-(make(chan struct{}))
}

func (a *App) update(parse parseFunc, data interface{}) error {
	var err error
	var f *os.File
	if a.OutputFile == "" {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(a.OutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			a.logger.Fatal("error opening outfile: %v", err)
		}
	}

	// We don't immediately check for this error. Even if an error occurs,
	// execution can continue. We'll return the error and log it.
	err = parse(a.Templates, data, f)

	if f != os.Stdout {
		err = f.Close()
		if err != nil {
			a.logger.Fatal("error closing outfile: %v", err)
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

func (a *App) parseHTML(infiles []string, data interface{}, w io.Writer) error {
	_, name := filepath.Split(infiles[0])

	t, err := htemplate.New(name).Funcs(a.HTMLFuncMap.HTML()).ParseFiles(infiles...)
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) parseText(infiles []string, data interface{}, w io.Writer) error {
	_, name := filepath.Split(infiles[0])

	t, err := ttemplate.New(name).Funcs(a.TextFuncMap.Text()).ParseFiles(infiles...)
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
