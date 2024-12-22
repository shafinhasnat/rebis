package main

import (
	"github.com/shafinhasnat/rebis/protocol"
	"github.com/shafinhasnat/rebis/rebis"
)

func main() {
	rebis.InitRebis()
	protocol.RebisProtocol()
	// rebis.MakeQuery("SET Hello World")
	// rebis.MakeQuery("SET Magfu World")
	// res := rebis.MakeQuery("GET Hello")
	// fmt.Println(res)
}
