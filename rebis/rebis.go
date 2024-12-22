package rebis

import (
	"fmt"
)

type Rebis struct {
	Map map[string]string
}

func (r *Rebis) Set(key string, value string) string {
	r.Map[key] = value
	return "OK"
}

func (r *Rebis) Get(key string) string {
	res := r.Map[key]
	return res
}

func (r *Rebis) Delete(key string) string {
	delete(r.Map, key)
	fmt.Println("DELETE", r)
	return "OK"
}
