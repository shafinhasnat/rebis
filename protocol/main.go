package protocol

import (
	"fmt"
	"net"
	"strings"

	"github.com/shafinhasnat/rebis/rebis"
)

func RebisProtocol(rebis *rebis.Rebis) {
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on port 6666")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		handler(conn, rebis)
	}
}
func handler(conn net.Conn, rebis *rebis.Rebis) {
	fmt.Println("Connection started")
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed:", err)
			break
		}
		query := string(buffer[:n])
		query = strings.TrimRight(query, "\n")
		query = strings.TrimRight(query, "\r")
		res := rebis.MakeQuery(query)
		r := fmt.Sprintf("> %s\n", res)
		_, err = conn.Write([]byte(r))
		if err != nil {
			fmt.Println("Error writing:", err)
		}
	}
}
