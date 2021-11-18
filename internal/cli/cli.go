package cli

import (
	"fmt"
	"os"

	cc "pkg/case_converter"

	"github.com/spf13/cobra"
)

func Commands() {
	tokenOutput := new([]string)

	var rootCmd = &cobra.Command{
		Use:     "csc",
		Short:   "convert cases of a string",
		Example: "csc camal my_string",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if *tokenOutput, err = cc.IdentifyCase(args[0]); err != nil {
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
			fmt.Println(cc.SnakeCase(*tokenOutput))
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "kebab",
		Short: "convert to `kebab-case`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cc.KebabCase(*tokenOutput))
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "pascal",
		Short: "convert to `PascalCase`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cc.PascalCase(*tokenOutput))
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "camal",
		Short: "convert to `camalCase`",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cc.CamalCase(*tokenOutput))
		},
	})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
