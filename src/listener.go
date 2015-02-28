package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

func startServer(listenPort int, backends *Backends) {
	port := strconv.Itoa(listenPort)
	fmt.Println("Starting server on port ", port)

	addr, _ := net.ResolveTCPAddr("tcp", ":"+port)

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Could not listen on port because:", err.Error())
		return
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("Error occured accepting a connection", err.Error())
		}

		go handleConnection(con, backends.NextAddress())
	}

}

func handleConnection(cli_conn net.Conn, srv_addr string) {
	srv_conn, err := net.Dial("tcp", srv_addr)
	if err != nil {
		fmt.Printf("Could not connect to server (%q), connection dropping\n", srv_addr)
		return
	}

	// close the conections when done
	defer func() {
		srv_conn.Close()
		cli_conn.Close()
	}()

	go io.Copy(cli_conn, srv_conn)
	io.Copy(srv_conn, cli_conn)
}
