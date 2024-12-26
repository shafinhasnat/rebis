package rebis

type Rebis struct {
	Map map[string]string
}

func InitRebis() *Rebis {
	return &Rebis{Map: make(map[string]string)}
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
	return "OK"
}

func (r *Rebis) Reset() string {
	r.Map = make(map[string]string)
	return "OK"
}
