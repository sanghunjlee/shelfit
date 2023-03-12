/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/sanghunjlee/shelfit/shelfit"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <book>",
	Aliases: []string{"a"},
	Example: `
shelfit add @anime chainsawman .good .action !started
shelfit add @manga dededede .favorite .drama +v1-v8 !finished`,
	Short: "Adds books",
	Long: `
Adds books

There are four main key flags to describe the "book" that you want to add:
'@': Category 
'.': Genre - a tag to describe the book (applies to its subitems, volumes)
'+': Volume - a sub-item that are related to the book
'!': Status - [unread, started, finished] - describes the status of the book (or the volume)`,

	Run: func(cmd *cobra.Command, args []string) {
		shelfit.NewApp().AddBook(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
