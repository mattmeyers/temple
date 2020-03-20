package main

import (
	"encoding/json"
	"flag"
	"fmt"
	htmltmpl "html/template"
	"io"
	"io/ioutil"
	"os"
	texttmpl "text/template"

	"github.com/mattmeyers/temple"
)

func main() {
	// temple --html --watch -d <datafile> -o <outfile> infiles...
	flagHTML := flag.Bool("html", false, "use html/template for template parsing")
	// flagWatch := flag.Bool("watch", false, "watch input files for changes")
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

	if *flagHTML {
		parseHTML(flag.Args(), data, w)
	} else {
		parseText(flag.Args(), data, w)
	}
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
	t, err := htmltmpl.New(infiles[0]).Funcs(temple.HTMLFuncMap).ParseFiles(infiles...)
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
	t, err := texttmpl.New(infiles[0]).Funcs(temple.TextFuncMap).ParseFiles(infiles...)
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
