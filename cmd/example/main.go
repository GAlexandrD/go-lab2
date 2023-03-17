package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/GAlexandrD/go-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File path to the expression data")
	outputFile      = flag.String("o", "", "File path to write the output")
)

func main() {
	flag.Parse()
	var reader io.Reader = nil
	var writer io.Writer = os.Stdout

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.OpenFile(*inputFile, os.O_RDONLY, 0777)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		reader = file
		defer file.Close()
	}

	if *outputFile != "" {
		file, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		writer = file
		defer file.Close()
	}
	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

}
