package protocol

import (
	"fmt"
	"net"

	"github.com/shafinhasnat/rebis/rebis"
)

func RebisProtocol() {
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
		handler(conn)
	}
}
func handler(conn net.Conn) {
	for {
		var buffer [1024]byte
		_, err := conn.Read(buffer[:])
		if err != nil {
			fmt.Println("Error reading:", err)
		}
		res := rebis.MakeQuery(string(buffer[:]))
		r := fmt.Sprintf("> %s\n", res)
		_, err = conn.Write(([]byte(r)))
		if err != nil {
			fmt.Println("Error writing:", err)
		}
	}
}
