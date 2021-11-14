package main
import (
	"strings"
)
func SnakeCase(words []string) string {
	var s = ""
	for i, word := range words {
		if i > 0 {
			s += "_"
		}
		s += strings.ToLower(word)
	}
	return s
}
func main() {

}