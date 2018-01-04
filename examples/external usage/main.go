package main

import (
	"fmt"
	"os"

	"git.3lab.re/marahin/goyam/goyam"
)

/* Keep in mind paths used expect you to run this program from goyam/ directory */
func main() {
	firstFile, err := goyam.LoadYAMLFile("examples/yaml files/swagger_1.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	secondFile, err := goyam.LoadYAMLFile("examples/yaml files/swagger_2.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Mutual keys: ")
	for _, key := range goyam.YAMLFindMutualKeys(firstFile, secondFile) {
		fmt.Println(key)
	}
}
