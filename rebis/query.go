package rebis

import (
	"strings"
)

func (r *Rebis) MakeQuery(query string) string {
	query_verbs := strings.Split(query, " ")
	switch verb := query_verbs[0]; verb {
	case "SET":
		if len(query_verbs) != 3 {
			return "ERROR"
		}
		return r.Set(query_verbs[1], query_verbs[2][:len(query_verbs[2])-1])
	case "GET":
		if len(query_verbs) != 2 {
			return "ERROR"
		}
		return r.Get(query_verbs[1][:len(query_verbs[1])-1])
	case "DELETE":
		if len(query_verbs) != 2 {
			return "ERROR"
		}
		return r.Delete(query_verbs[1][:len(query_verbs[1])-1])
	default:
		return "ERROR"
	}
}
