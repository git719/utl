// iplib.go

package utl

import (
	"net"
	"net/http"
	"time"
)

func ValidIpStr(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil
}

func IsIpPortStrReachable(ipPortStr string) bool {
	// Checks if IP_Address:Port string is reachable
	conn, err := net.DialTimeout("tcp", ipPortStr, time.Second*3)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func InternetIsAvailable() bool {
	_, err := http.Get("http://httpbin.org/ip")
	return err == nil
}

func InternetNotAvailable() bool {
	return !InternetIsAvailable()
}
