package main

import (
	"flag"
	"fmt"
	"os"
)

const appVersion = "0.0.1"

func stringIsEmpty(s *string) bool {
	return s == nil || *s == ""
}

func main() {

	version := flag.Bool("version", false, "print version")
	templateFile := flag.String("template-file", "", "template-file")
	templateMap := flag.String("template-map", "", "template-map")

	flag.Parse()
	if version != nil && *version {
		fmt.Fprintf(os.Stderr, "%s\n", appVersion)
		return
	}

	if stringIsEmpty(templateFile) || stringIsEmpty(templateMap) {
		fmt.Fprintln(os.Stderr, "template file or template map is empty!")
		return
	}
	GenTemplate(*templateFile, *templateMap)
}
