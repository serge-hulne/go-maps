package maps


//---
// generic interface{} Key() based Map:
//---

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


//---
// Specialized String() key based SMap:
//---

type Stringer interface {
	String() string
}

type SMap map[string]Pair

func NewSMap() SMap {
	return make(map[string]Pair)
}

func (m SMap) Insert(key Stringer, value interface{}) {
	m[key.String()] = Pair{key, value}
}

func (m SMap) Do(f func(Pair)) {
	for key, value := range m {
		f(Pair{key, value})
	}
}

func (m SMap) Get(key Stringer) (interface{}, bool) {
	v, t := m[key.String()]
	return v, t
}

func (m SMap) Delete(key Stringer) {
	m[key.String()] = Pair{nil, nil}, false
}

func (m SMap) Len() int {
	return len(m)
}


//---
// Specialized Int() key based IMap :
//---

type Inter interface {
	Int() int
}

type IMap map[int]Pair

func NewIMap() IMap {
	return make(map[int]Pair)
}

func (m IMap) Insert(key Inter, value interface{}) {
	m[key.Int()] = Pair{key, value}
}

func (m IMap) Do(f func(Pair)) {
	for key, value := range m {
		f(Pair{key, value})
	}
}

func (m IMap) Get(key Inter) (interface{}, bool) {
	v, t := m[key.Int()]
	return v, t
}

func (m IMap) Delete(key Inter) {
	m[key.Int()] = Pair{nil, nil}, false
}

func (m IMap) Len() int {
	return len(m)
}
