// Get ip utils
// Author：foolin

package webutil

import (
	"encoding/binary"
	"net"
	"net/http"
	"strings"
)

//Ip get ip address
func Ip(req *http.Request) string {
	ips := proxy(req)
	if ips != nil && ips[0] != "" {
		pos := strings.LastIndex(ips[0], ":")
		if pos == -1 {
			return ips[0]
		}
		return ips[0][:pos]
	}

	remoteAddr := req.Header.Get("Remote_addr")
	if remoteAddr == "" {
		remoteAddr = req.Header.Get("X-Real-IP")
		if remoteAddr == "" {
			remoteAddr = req.RemoteAddr
		}
	}

	if remoteAddr == "" {
		return "127.0.0.1"
	}

	pos := strings.LastIndex(remoteAddr, ":")
	if pos == -1 {
		return remoteAddr
	}
	return remoteAddr[:pos]
}

//proxy 
func proxy(req *http.Request) []string {
	if ips := req.Header.Get("X-Forwarded-For"); ips != "" {
		return strings.Split(ips, ",")
	}
	return nil
}

//Ip2long iptooogong
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}