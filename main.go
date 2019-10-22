package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var isVerbose bool
	flag.BoolVar(&isVerbose, "verbose", false, "Output detail log to the stderr")
	flag.Parse()

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	var m map[string]string
	if err := yaml.NewDecoder(f).Decode(&m); err != nil {
		log.Fatal(err)
	}

	for k, v := range m {
		e := fmt.Sprintf("export %s='%s';", k, v)
		if isVerbose {
			fmt.Fprintln(os.Stderr, e)
		}

		fmt.Print(e)
	}
}
