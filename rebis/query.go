package rebis

import (
	"regexp"
	"strings"
)

func (r *Rebis) MakeQuery(query string) string {
	query_parts := strings.SplitN(query, " ", 3)
	if len(query_parts) > 1 {
		if matched, _ := regexp.MatchString(`\s|^.{17,}$`, query_parts[1]); matched {
			return "ERROR"
		}
	}
	switch verb := query_parts[0]; verb {
	case "SET":
		if len(query_parts) != 3 {
			return "ERROR"
		}
		return r.Set(query_parts[1], query_parts[2])
	case "GET":
		if len(query_parts) != 2 {
			return "ERROR"
		}
		return r.Get(query_parts[1])
	case "DELETE":
		if len(query_parts) != 2 {
			return "ERROR"
		}
		return r.Delete(query_parts[1])
	case "RESET":
		if len(query_parts) != 1 {
			return "ERROR"
		}
		return r.Reset()
	default:
		return "ERROR"
	}
}
