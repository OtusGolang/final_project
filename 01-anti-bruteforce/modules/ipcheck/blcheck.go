package ipcheck

import (
	"database/sql"
	"fmt"
	"log"
	"net"
)

type BlacklistEntry struct {
	Subnet *net.IPNet
}

func CheckIpinBL(request_ip string, db *sql.DB) bool {
	entries, err := fetchBlacklistEntries(db)
	if err != nil {
		log.Println(err)
	}

	isAllowed := false
	for _, entry := range entries {
		fmt.Println("request_ip: ", request_ip)
		if entry.Subnet.Contains(net.ParseIP(request_ip)) {
			isAllowed = true
			break
		}
	}
	return isAllowed
}

func fetchBlacklistEntries(db *sql.DB) ([]WhitelistEntry, error) {
	rows, err := db.Query("SELECT subnet FROM blacklist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []WhitelistEntry
	for rows.Next() {
		var subnetStr string
		if err := rows.Scan(&subnetStr); err != nil {
			return nil, err
		}

		_, subnet, err := net.ParseCIDR(subnetStr)
		if err != nil {
			log.Println("Invalid subnet:", subnetStr)
			continue
		}

		entry := WhitelistEntry{
			Subnet: subnet,
		}
		entries = append(entries, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
