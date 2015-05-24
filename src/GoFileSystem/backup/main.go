package main

import (
	"fmt"
	"net"
	"os"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "5000"
	CONN_TYPE = "tcp"
)

var sig_end chan int = make(chan int, 2)

func main() {
	go TcpServer()
	go HttpServer()
	<- sig_end
	<- sig_end
}

func TcpServer() {
	defer sigEnd()

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

func HttpServer() {
	defer sigEnd()

	r := mux.NewRouter()
	r.HandleFunc("/kvman/countkey", HandleCountKey).Methods("GET")
	r.HandleFunc("/kvman/dump", HandleDump).Methods("GET")
	r.HandleFunc("/kvman/shutdown", HandleShutdown).Methods("GET")
	http.ListenAndServe("localhost:8000", r)
}

func sigEnd() {
	sig_end <- 1
}
