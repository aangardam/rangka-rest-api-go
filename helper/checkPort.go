package helper

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func IsPortAvalaible(port int) bool {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	conn.Close()
	return true
}

func findAvailablePort(startPort int) int {
	for port := startPort; port <= startPort+100; port++ {
		if IsPortAvalaible(port) {
			return port
		}
	}
	return 0
}

func GetAvalaiblePort(portSrt string) int {
	port, err := strconv.Atoi(portSrt)
	if err != nil {
		panic(err)
	}
	if !IsPortAvalaible(port) {
		port = findAvailablePort(port)
		if port == 0 {
			log.Fatalf("Port Not Found")
		}
	}
	return port
}
