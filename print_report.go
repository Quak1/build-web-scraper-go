package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	sortedPages := sortPages(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.name)
	}
}

type page struct {
	name  string
	count int
}

func sortPages(pagesMap map[string]int) []page {
	pages := []page{}

	for name, count := range pagesMap {
		pages = append(pages, page{name: name, count: count})
	}

	sort.Slice(pages, func(i, j int) bool {
		if pages[i].count == pages[j].count {
			return pages[i].name < pages[j].name
		}
		return pages[i].count > pages[j].count
	})

	return pages
}
