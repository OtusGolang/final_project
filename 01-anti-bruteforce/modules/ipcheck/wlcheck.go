package ipcheck

import (
	"database/sql"
	"log"
	"net"
)

type WhitelistEntry struct {
	Subnet *net.IPNet
}

func CheckIpinWL(request_ip string, db *sql.DB) bool {
	entries, err := fetchWhitelistEntries(db)
	if err != nil {
		log.Println(err)
	}

	isAllowed := false
	for _, entry := range entries {
		if entry.Subnet.Contains(net.ParseIP(request_ip)) {
			isAllowed = true
			break
		}
	}
	return isAllowed
}

func fetchWhitelistEntries(db *sql.DB) ([]WhitelistEntry, error) {
	rows, err := db.Query("SELECT subnet FROM whitelist")
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
