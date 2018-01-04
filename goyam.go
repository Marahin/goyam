package main

import (
	"fmt"
	"os"
	"strconv"

	. "git.3lab.re/marahin/goyam/goyam"
)

func main() {
	firstFile, err := LoadYAMLFile(os.Args[1:][0])
	if err != nil {
		fmt.Println("Error loading first file...")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	secondFile, err := LoadYAMLFile(os.Args[1:][1])

	summary := YAMLCompare(firstFile, secondFile, "root >")

	if summary.ErrorsCount() == 0 && summary.WarningsCount() == 0 {
		fmt.Printf(Green("No differences found!\n"))
		os.Exit(0)
	} else {
		for id, err := range summary.Errors {
			fmt.Printf("%s %s\n", Red("(error "+strconv.Itoa(id)+")"), err)
		}
		for id, warn := range summary.Warnings {
			fmt.Printf("%s %s\n", Yellow("(warning "+strconv.Itoa(id)+")"), warn)
		}
		fmt.Printf("%s errors, %s warnings found.\n",
			Red(strconv.Itoa(summary.ErrorsCount())),
			Yellow(strconv.Itoa(summary.WarningsCount())),
		)
		os.Exit(1)
	}
}
