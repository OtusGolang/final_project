package handlers

import (
	"database/sql"
	"log"
)

func AddIPWhiteList(subnet string, db *sql.DB) bool {
	_, err := db.Exec("INSERT INTO whitelist (subnet) VALUES ($1)", subnet)
	if err != nil {
		log.Println("Error INSERT INTO whitelist")
	}
	return true
}

func DelIPWhiteList(subnet string, db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM whitelist WHERE subnet=$1", subnet)
	if err != nil {
		log.Println("Error DELETE FROM whitelist")
	}
	return true
}
