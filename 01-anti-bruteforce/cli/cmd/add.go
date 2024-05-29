package cmd

import (
	"01-anti-bruteforce/modules/api/handlers"
	"01-anti-bruteforce/modules/db"
	"fmt"

	"github.com/spf13/cobra"
)

// ResetCmd represents the add command.
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add to black/white list",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.DbConnect(HOST, PORT, USER, POSTGRES_PASSWORD, POSTGRES_DBNAME)
		defer db.Close()

		list := cmd.PersistentFlags().Lookup("kind").Value.String()
		if list == "whitelist" {
			status_del_bl := handlers.AddIPWhiteList(cmd.PersistentFlags().Lookup("network").Value.String(), db)
			fmt.Println("cli status_del_bl: ", status_del_bl)
		} else if list == "blacklist" {
			status_del_bl := handlers.AddIPBlackList(cmd.PersistentFlags().Lookup("network").Value.String(), db)
			fmt.Println("cli status_del_bl: ", status_del_bl)

		}
	},
}

func init() {
	AddCmd.PersistentFlags().StringP("network", "n", "", "network")
	RemoveCmd.PersistentFlags().StringP("kind", "k", "blacklist", fmt.Sprintf("%s or %s", "blacklist", "whitelist"))
	AddCmd.PersistentFlags().StringP("mask", "m", "", "mask")
}
