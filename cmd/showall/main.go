package main

import (
	"fmt"

	"github.com/flosch/pongo2/v6"
)

func main() {
	fmt.Printf("Printing all tags and filters for version %s\n\n", pongo2.Version)

	allTags := pongo2.GetAllTags()
	allFilters := pongo2.GetAllFilters()

	fmt.Println("TAGS")
	fmt.Println("----")
	for _, tkey := range allTags {
		fmt.Printf("\"%s\",\n", tkey)
	}
	fmt.Println("----")

	fmt.Println("FILTERS")
	fmt.Println("----")
	for _, fkey := range allFilters {
		fmt.Printf("\"%s\",\n", fkey)
	}
	fmt.Println("----")
}
