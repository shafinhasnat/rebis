package main

import (
	"github.com/shafinhasnat/rebis/protocol"
	"github.com/shafinhasnat/rebis/rebis"
)

func main() {
	rebis := rebis.InitRebis()
	protocol.RebisProtocol(rebis)
}
