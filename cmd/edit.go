/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/sanghunjlee/shelfit/shelfit"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit <book>",
	Short: "edit an existing book",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

There are four main key flags to describe the "book" that you want to add:
'@': Category 
'.': Genre - a tag to describe the book (applies to its subitems, volumes)
'+': Volume - a sub-item that are related to the book
'!': Status - [unread, started, finished] - describes the status of the book (or the volume)`,
	Run: func(cmd *cobra.Command, args []string) {
		shelfit.NewApp().EditBook(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
