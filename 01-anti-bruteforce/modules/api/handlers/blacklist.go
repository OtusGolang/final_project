package handlers

import (
	"database/sql"
	"log"
)

func AddIPBlackList(subnet string, db *sql.DB) bool {
	_, err := db.Exec("INSERT INTO blacklist (subnet) VALUES ($1)", subnet)
	if err != nil {
		log.Println("Error INSERT INTO blacklist")
	}
	return true
}

func DelIPBlackList(subnet string, db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM blacklist WHERE subnet=$1", subnet)
	if err != nil {
		log.Println("Error DELETE FROM blacklist")
	}
	return true
}
