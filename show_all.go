package pongo2

import "sort"

func GetAllTags() []string {
	allTags := make([]string, 0, len(tags))
	for k := range tags {
		allTags = append(allTags, k)
	}
	sort.Strings(allTags)
	return allTags
}

func GetAllFilters() []string {
	allFilters := make([]string, 0, len(filters))
	for k := range filters {
		allFilters = append(allFilters, k)
	}
	sort.Strings(allFilters)
	return allFilters
}
