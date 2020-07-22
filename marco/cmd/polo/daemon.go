package main

import "github.com/spf13/cobra"

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "manage daemon instance",
	Long:  "manage daemon instance",
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
