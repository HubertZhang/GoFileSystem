package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "5000"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if (err != nil) {
		fmt.Println("Main::LISTEN::Error: " + err.Error())
		os.Exit(1)
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if (err != nil) {
			fmt.Println("Main::ACCEPT::Error: " + err.Error())
			os.Exit(1)
		}

		go HandleConnection(conn)
	}
}
