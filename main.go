package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
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
func IdentifyCase(input string) ([]string, error) {
	snakeCase := regexp.MustCompile("(.*?)_([a-zA-Z]+)")
	kebabCase := regexp.MustCompile("^([a-z][a-z0-9]*)(-[a-z0-9]+)*$")
	pascalCase := regexp.MustCompile("[A-Z]([A-Z0-9]*[a-z][a-z0-9]*[A-Z]|[a-z0-9]*[A-Z][A-Z0-9]*[a-z])[A-Za-z0-9]*")
	camalCase := regexp.MustCompile("^[a-z][a-zA-Z0-9]+$")
	switch {
	case snakeCase.MatchString(input):
		return strings.Split(input, "_"), nil
	case kebabCase.MatchString(input):
		return strings.Split(input, "-"), nil
	case pascalCase.MatchString(input):
		var words []string = make([]string, 0)
		word := ""
		for i, char := range input {
			if unicode.IsUpper(char) && i > 0 {
				words = append(words, word)
				word = ""
			}
			word += string(char)
		}
		words = append(words, word)
		return words, nil
	case camalCase.MatchString(input):
		var words []string = make([]string, 0)
		word := ""
		for _, char := range input {
			if unicode.IsUpper(char) {
				words = append(words, word)
				word = ""
			}
			word += string(char)
		}
		words = append(words, word)
		return words, nil
	}
	return nil, errors.New("Unable to identify case")
}
func cli() {
	tokenOutput := new([]string)

	var rootCmd = &cobra.Command{
		Use:     "csc",
		Short:   "convert cases of a string",
		Example: "csc camal my_string",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if *tokenOutput, err = IdentifyCase(args[0]); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true}) // no help command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "snake",
		Short: "convert to `snake_case`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(SnakeCase(*tokenOutput))
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "kebab",
		Short: "convert to `kebab-case`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(KebabCase(*tokenOutput))
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "pascal",
		Short: "convert to `PascalCase`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(PascalCase(*tokenOutput))
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "camal",
		Short: "convert to `camalCase`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(CamalCase(*tokenOutput))
		},
	})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	cli()
}
