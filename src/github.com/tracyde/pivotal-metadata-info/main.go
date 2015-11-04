package main

import (
	"github.com/spf13/cobra"
)

var fileName string

var rootCmd = &cobra.Command{
	Use:   "pivotal-metadata-info",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&fileName, "fname", "f", "", "Name of the file to read")
}

func main() {
	AddCommands()
	rootCmd.Execute()
}

func AddCommands() {
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(listCmd)
}
