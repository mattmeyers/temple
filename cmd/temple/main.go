package main

import (
	"fmt"
	"os"

	"github.com/mattmeyers/temple"
	"github.com/mattmeyers/temple/pkg/cli"
)

func main() {
	if err := cli.New().WithFuncMap(temple.FullFuncMap()).ClearTextFuncMap().Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
