package ipcheck

import (
	"log"
	"net"
	"net/http"
)

func ParseIP(r *http.Request) string {

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		log.Println("Invalid IP address")
	}
	return parsedIP.String()
}
