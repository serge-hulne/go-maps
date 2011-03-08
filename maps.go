package maps


type Map map[interface{}]interface{}

type Pair struct {
	Key   interface{}
	Value interface{}
}

type Keyer interface {
	Key() interface{}
}

func key(k interface{}) interface{} {
	if k, ok := k.(Keyer); ok {
		return k.Key()
	}
	return k
}

func (m Map) Get(k interface{}) (interface{}, bool) {
	v, Ok := m[key(k)]
	return v.(Pair).Value, Ok
}

func (m Map) Delete(k interface{}) {
	m[key(k)] = nil, false
}

func (m Map) Insert(k interface{}, v interface{}) {
	m[key(k)] = Pair{k, v}
}

func (m Map) Len() int {
	return len(m)
}

func (m Map) Do(f func(Pair)) {
	for _, p := range m {
		f(p.(Pair))
	}
}
