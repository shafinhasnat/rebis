package rebis

import (
	"strings"
)

var GlobalRebis Rebis

func InitRebis() {
	GlobalRebis = Rebis{Map: make(map[string]string)}
}

func MakeQuery(query string) string {
	query_verbs := strings.Split(query, " ")
	if query_verbs[0] == "SET" {
		res := GlobalRebis.Set(query_verbs[1], query_verbs[2])
		return res
	} else if query_verbs[0] == "GET" {
		res := GlobalRebis.Get(query_verbs[1])
		return res
	} else if query_verbs[0] == "DELETE" {
		res := GlobalRebis.Delete(query_verbs[1])
		return res
	}
	return "OK"
}
