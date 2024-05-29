package cmd

import (
	"01-anti-bruteforce/modules/limiter"
	"fmt"

	"github.com/spf13/cobra"
)

var Newlimiter = limiter.NewRateLimiter()

// ResetCmd represents the add command.
var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset bucket",
	Run: func(cmd *cobra.Command, args []string) {

		status_clear_bucket := Newlimiter.ClearBucket(Newlimiter, cmd.PersistentFlags().Lookup("login").Value.String(), cmd.PersistentFlags().Lookup("ip").Value.String())
		fmt.Println("cli status_clear_bucket: ", status_clear_bucket)
	},
}

func init() {
	ResetCmd.PersistentFlags().StringP("login", "l", "", "login for reset bucket")
	ResetCmd.PersistentFlags().StringP("ip", "i", "", "ip for reset bucket")
}
