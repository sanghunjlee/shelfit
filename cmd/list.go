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
	expand  bool
	groupBy string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		shelfit.NewApp().ListBooks(strings.Join(args, " "), expand, groupBy)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&expand, "expand", "e", false, "Expand the list to include volumes (if exist)")
	listCmd.Flags().StringVarP(&groupBy, "groupBy", "g", "", "Group books and list accordingly")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
