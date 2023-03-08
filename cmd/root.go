package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shelfit",
	Short: "A simple way to manage lists (or 'shelves') of items.",
	Long: `
A simple way to manage lists (or 'shelves') of items.

Create and manage a shelf in which you can add and remove
items.

For more info, check: https://github.com/sanghunjlee/shelfit
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
