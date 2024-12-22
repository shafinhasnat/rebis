package rebis

import (
	"strings"
)

func InitRebis() *Rebis {
	return &Rebis{Map: make(map[string]string)}
}

func (r *Rebis) MakeQuery(query string) string {
	query_verbs := strings.Split(query, " ")
	if query_verbs[0] == "SET" {
		res := r.Set(query_verbs[1], query_verbs[2][:len(query_verbs[2])-1])
		return res
	} else if query_verbs[0] == "GET" {
		res := r.Get(query_verbs[1][:len(query_verbs[1])-1])
		return res
	} else if query_verbs[0] == "DELETE" {
		res := r.Delete(query_verbs[1][:len(query_verbs[1])-1])
		return res
	}
	return "ERROR"
}
