package main

import (
	"github.com/shafinhasnat/rebis/protocol"
	"github.com/shafinhasnat/rebis/rebis"
)

func main() {
	db := rebis.InitRebis()
	protocol.RebisProtocol(db)
}
