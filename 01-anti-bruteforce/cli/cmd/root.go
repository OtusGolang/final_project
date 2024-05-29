package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// var (
// 	whiteListKey = "whitelist"
// 	blackListKey = "blacklist"
// )

var mainCmd = &cobra.Command{
	Use:   "cli-client",
	Short: "CLI client for antibrute force service",
}

func Execute() {
	mainCmd.AddCommand(AddCmd)
	mainCmd.AddCommand(RemoveCmd)
	mainCmd.AddCommand(ResetCmd)

	err := mainCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
