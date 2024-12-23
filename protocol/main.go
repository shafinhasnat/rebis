package protocol

import (
	"fmt"
	"net"

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
		buffer := make([]byte, 128)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed:", err)
			break
		}
		res := rebis.MakeQuery(string(buffer[:n]))
		r := fmt.Sprintf("> %s\n", res)
		_, err = conn.Write([]byte(r))
		if err != nil {
			fmt.Println("Error writing:", err)
		}
	}
}
