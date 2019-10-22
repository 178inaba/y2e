package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var m map[string]string
	if err := yaml.NewDecoder(f).Decode(&m); err != nil {
		log.Fatal(err)
	}

	for k, v := range m {
		fmt.Printf("export %s='%s';", k, v)
	}
}
