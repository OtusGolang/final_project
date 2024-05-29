package main

import (
	"01-anti-bruteforce/modules/api/handlers"
	"01-anti-bruteforce/modules/db"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("BL-add", func(t *testing.T) {
		db := db.DbConnect("127.0.0.1", 5432, "postgres", "postgres", "postgres")
		defer db.Close()
		status_add_bl := handlers.AddIPBlackList("8.8.8.8/24", db)
		rows, err := db.Query("SELECT subnet FROM blacklist WHERE subnet=$1", "8.8.8.8/24")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var subnet string
		for rows.Next() {
			err = rows.Scan(&subnet)
			if err != nil {
				log.Fatal(err)
			}
		}
		require.Equal(t, status_add_bl, true)
		require.Equal(t, subnet, "8.8.8.8/24")
	})

	t.Run("BL-delete", func(t *testing.T) {
		db := db.DbConnect("127.0.0.1", 5432, "postgres", "postgres", "postgres")
		defer db.Close()
		status_add_bl := handlers.DelIPBlackList("8.8.8.8/24", db)
		rows, err := db.Query("DELETE FROM blacklist WHERE subnet=$1", "8.8.8.8/24")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var subnet string
		for rows.Next() {
			err = rows.Scan(&subnet)
			if err != nil {
				log.Fatal(err)
			}
		}
		require.Equal(t, status_add_bl, true)
		require.Equal(t, subnet, "")
	})

	t.Run("WL-add", func(t *testing.T) {
		db := db.DbConnect("127.0.0.1", 5432, "postgres", "postgres", "postgres")
		defer db.Close()
		status_add_bl := handlers.AddIPWhiteList("8.8.8.8/24", db)
		rows, err := db.Query("SELECT subnet FROM whitelist WHERE subnet=$1", "8.8.8.8/24")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var subnet string
		for rows.Next() {
			err = rows.Scan(&subnet)
			if err != nil {
				log.Fatal(err)
			}
		}
		require.Equal(t, status_add_bl, true)
		require.Equal(t, subnet, "8.8.8.8/24")
	})

	t.Run("WL-delete", func(t *testing.T) {
		db := db.DbConnect("127.0.0.1", 5432, "postgres", "postgres", "postgres")
		defer db.Close()
		status_add_bl := handlers.AddIPWhiteList("8.8.8.8/24", db)
		rows, err := db.Query("DELETE FROM whitelist WHERE subnet=$1", "8.8.8.8/24")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var subnet string
		for rows.Next() {
			err = rows.Scan(&subnet)
			if err != nil {
				log.Fatal(err)
			}
		}
		require.Equal(t, status_add_bl, true)
		require.Equal(t, subnet, "")
	})

	t.Run("Login", func(t *testing.T) {
		db := db.DbConnect("127.0.0.1", 5432, "postgres", "postgres", "postgres")
		defer db.Close()
		status_add_bl := handlers.AddIPWhiteList("8.8.8.8/24", db)
		rows, err := db.Query("DELETE FROM whitelist WHERE subnet=$1", "8.8.8.8/24")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var subnet string
		for rows.Next() {
			err = rows.Scan(&subnet)
			if err != nil {
				log.Fatal(err)
			}
		}
		require.Equal(t, status_add_bl, true)
		require.Equal(t, subnet, "")
	})
}
