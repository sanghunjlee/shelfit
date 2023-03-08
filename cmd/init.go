package cmd

import (
	"github.com/sanghunjlee/shelfit/shelfit"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new 'shelf' in the current directory",
	Long: `
Initializes a new 'shelf' in the current directory.

This will generate a '.shelf.json' file in the current directory.
After which, you can start adding 'books' into it. 

For more info, check: https://github.com/sanghunjlee/shelfit`,
	Run: func(cmd *cobra.Command, args []string) {
		shelfit.NewApp().Initialize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
