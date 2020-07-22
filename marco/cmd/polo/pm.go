package main

import "github.com/spf13/cobra"

var pmCmd = &cobra.Command{
	Use:   "pm",
	Short: "manage nginx package",
	Long:  "manage nginx package",
}

func init() {
	rootCmd.AddCommand(pmCmd)
}
