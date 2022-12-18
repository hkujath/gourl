package main

import (
	"flag"
	"fmt"
	"gourl/httphandling"
	"io"
	"os"
	"path/filepath"
)

var (
	flagOutput = flag.String("o", "", "output file")
	flagHeader = flag.Bool("header", false, "print HTTP-header")
)

// init is used to initialize the command line arguments.
func init() {
	flag.Parse()
}

// main function
// can be called from commandline with "gourl http://www.google.de"
func main() {
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Please enter just one URL.")
		os.Exit(1)
	}

	url := args[0]
	if !httphandling.ValidateURL(url) {
		fmt.Println("Non valid URL found.")
		os.Exit(1)
	}

	respBody := httphandling.GetHttpRequest(url, *flagHeader)
	defer respBody.Close()

	var writer io.Writer = os.Stdout

	writeOutput(writer, respBody)
}

// writeOutput writes website to output (can be a file)
func writeOutput(writer io.Writer, body io.ReadCloser) {
	tmpWriter := writer

	if *flagOutput != "" {

		// Create complete path, if not existent
		err := os.MkdirAll(filepath.Dir(*flagOutput), 0755)
		if err != nil {
			fmt.Printf("Couldn't create output path for file %s\n%s", *flagOutput, err)
			os.Exit(1)
		}

		f, err := os.OpenFile(*flagOutput, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Couldn't write output file %s\n%s", *flagOutput, err)
			os.Exit(1)
		}
		defer f.Close()
		tmpWriter = f
	}

	io.Copy(tmpWriter, body)

}
