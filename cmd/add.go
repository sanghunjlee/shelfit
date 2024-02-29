/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/sanghunjlee/shelfit/shelfit"
	"github.com/spf13/cobra"
)

// variables for flags
var (
	note string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <item>",
	Aliases: []string{"a"},
	Example: `
shelfit add @anime chainsawman .good .action !started
shelfit add @manga dededede .favorite .drama +v1-v8 !finished`,
	Short: "Add an item to the shelf",
	Long: `
Add an item to the shelf

There are inline flags to describe the item that you want to add:
'!': Category (required)
'#': Tag
`,

	Run: func(cmd *cobra.Command, args []string) {
		shelfit.NewApp().AddBook(strings.Join(args, " "), note)
	},
}

func init() {
	noteFlagUsage := "Add a short note (Please encapsulate with \"\")\nUse [addNote] command for more detailed notes"

	addCmd.Flags().StringVarP(&note, "note", "n", "", noteFlagUsage)
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
