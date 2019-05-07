package medutils

import "strings"

func ConvertToLower(tags []string) []string {
	var tagsLower []string
	for _, tag := range tags {
		tag = strings.ToLower(strings.TrimSpace(tag))
		tagsLower = append(tagsLower, tag)
	}

	return tagsLower
}
