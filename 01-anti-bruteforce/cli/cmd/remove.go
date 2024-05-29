package cmd

import (
	"01-anti-bruteforce/modules/api/handlers"
	"01-anti-bruteforce/modules/db"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var HOST = os.Getenv("HOST")
var PORT = 5432
var USER = os.Getenv("USER")
var POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var POSTGRES_DBNAME = os.Getenv("POSTGRES_DBNAME")

// ResetCmd represents the add command.
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove from black/white list",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.DbConnect(HOST, PORT, USER, POSTGRES_PASSWORD, POSTGRES_DBNAME)
		//db := db.DbConnect("127.0.0.1", 5432, "postgres", "postgres", "postgres")
		defer db.Close()

		list := cmd.PersistentFlags().Lookup("kind").Value.String()
		if list == "whitelist" {
			status_del_bl := handlers.DelIPWhiteList(cmd.PersistentFlags().Lookup("network").Value.String(), db)
			fmt.Println("cli status_del_bl: ", status_del_bl)
		} else if list == "blacklist" {
			status_del_bl := handlers.DelIPBlackList(cmd.PersistentFlags().Lookup("network").Value.String(), db)
			fmt.Println("cli status_del_bl: ", status_del_bl)

		}
	},
}

func init() {
	RemoveCmd.PersistentFlags().StringP("network", "n", "", "network")
	RemoveCmd.PersistentFlags().StringP("kind", "k", "blacklist", "blacklist or whitelist")
}
