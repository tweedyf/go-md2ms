package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/russross/blackfriday/v2"
)

var outFilePath = flag.String("o", "", "Path to output processed file (default: stdout)")
// Note: md2ms takes input from stdin

// Render converts a markdown document into a roff/ms formatted document.
func Render(doc []byte) []byte {
	renderer := NewRoffRenderer()

	return blackfriday.Run(doc,
		[]blackfriday.Option{blackfriday.WithRenderer(renderer),
			blackfriday.WithExtensions(renderer.GetExtensions())}...)
}

func main() {
	var err error
	flag.Parse()

	inFile := os.Stdin
	inFilePath := flag.Arg(0)
	if inFilePath != "" {
		inFile, err = os.Open(inFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	defer inFile.Close() // nolint: errcheck

	doc, err := ioutil.ReadAll(inFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out := Render(doc)

	outFile := os.Stdout
	if *outFilePath != "" {
		outFile, err = os.Create(*outFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer outFile.Close() // nolint: errcheck
	}
	_, err = outFile.Write(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
