package network

import (
	"api-authenticator-proxy/util/log"
	"fmt"
	"net"
	"strconv"
)

func IsPortValid(port string) bool {
	intPort, err := strconv.Atoi(port)
	if err != nil {
		log.Warning("Invalid port number: %s", port)
		return false
	}
	address := fmt.Sprintf(":%d", intPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false // Port is in use
	}
	listener.Close()
	return true // Port is not in use
}
