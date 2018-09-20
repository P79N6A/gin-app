package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	checkError(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		conn.Write([]byte("had receive client message"))

	}
}
func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal error: %s", err.Error())
		os.Exit(1)
	}
}
