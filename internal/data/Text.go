package data

import "strings"

var (
	categories = []string{"Basketball", "Soccer", "Running", "Golf", "Skate", "Tennis", "Baseball"}
)

func CleanCategory(raw string) string {
	category := "Men's shoes"
	for _, i := range categories {
		if strings.Contains(raw, i) {
			category = i
		}
	}
	return category
}
