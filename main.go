package main

import (
	"strings"
)

func basecase(words []string, delimiter string) string {
	var s = ""
	for i, word := range words {
		if i > 0 {
			s += delimiter
		}
		s += strings.ToLower(word)
	}
	return s
}

func SnakeCase(words []string) string {
	return basecase(words, "_")
}

func KebabCase(words []string) string {
	return basecase(words, "-")
}
func PascalCase(words []string) string {
	var s = ""
	for _, word := range words {
		s += strings.Title(word)
	}
	return s
}
func CamalCase(words []string) string {
	var s = ""
	for i, word := range words {
		if i == 0 {
			s += strings.ToLower(word)
		} else if i > 0 {
			s += strings.Title(word)
		}
	}
	return s
}
func main() {

}
