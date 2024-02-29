/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/sanghunjlee/shelfit/shelfit"
	"github.com/spf13/cobra"
)

var (
	category string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "list books",
	Long: `
list books in the shelf
	`,
	Run: func(cmd *cobra.Command, args []string) {
		shelfit.NewApp().ListBooks(strings.Join(args, " "), category)
	},
}

func init() {
	listCmd.Flags().StringVarP(&category, "category", "c", "", "Filter by category")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
