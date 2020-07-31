package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	htmpl "html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	ttmpl "text/template"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"github.com/mattmeyers/temple"
)

func main() {
	flag.Usage = usage
	flagHTML := flag.Bool("html", false, "use html/template for template parsing")
	flagWatch := flag.Bool("watch", false, "watch input files for changes")
	flagFuncs := flag.String("funcs", "", "load a Go file to add to the FuncMap")
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

	funcs, err := readFuncsFile(*flagFuncs)
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

	var f func([]string, temple.FuncMap, interface{}, io.Writer) error
	if *flagHTML {
		f = parseHTML
	} else {
		f = parseText
	}

	if *flagWatch {
		fmt.Println("file watching not implement, try again later")
		os.Exit(1)
	} else {
		err = f(flag.Args(), funcs, data, w)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
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

func readFuncsFile(filename string) (temple.FuncMap, error) {
	if filename == "" {
		return nil, nil
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	_, err = i.Eval(string(buf))
	if err != nil {
		return nil, err
	}

	fMap := temple.FuncMap{}

	fset := token.NewFileSet()
	af, err := parser.ParseFile(fset, "", buf, 0)
	if err != nil {
		log.Fatal(err)
	}

	pkgName := af.Name.Name

	for _, decl := range af.Decls {
		if v, ok := decl.(*ast.FuncDecl); ok {
			funcName := v.Name.Name
			fullName := pkgName + "." + funcName
			fn, err := i.Eval(fullName)
			if err != nil {
				return nil, err
			}
			fMap[funcName] = fn.Interface()
		}
	}

	return fMap, nil
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

func parseHTML(infiles []string, funcs temple.FuncMap, data interface{}, w io.Writer) error {
	fMap := temple.MergeFuncMaps(temple.FullFuncMap(), funcs)

	t, err := htmpl.New(infiles[0]).Funcs(fMap.HTML()).ParseFiles(infiles...)
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func parseText(infiles []string, funcs temple.FuncMap, data interface{}, w io.Writer) error {
	fMap := temple.MergeFuncMaps(temple.FullFuncMap(), funcs)

	t, err := ttmpl.New(infiles[0]).Funcs(fMap.Text()).ParseFiles(infiles...)
	if err != nil {
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
