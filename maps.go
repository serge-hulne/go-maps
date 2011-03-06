package maps


//---
type Pair struct {
	Key   interface{}
	Value interface{}
}

type ValueOk struct {
	Value interface{}
	Ok    bool
}

//---
type Stringer interface {
	String() string
}

type SMap map[string]Pair

func NewSMap() SMap {
	return make(map[string]Pair)
}

func (m *SMap) Insert(key Stringer, value interface{}) {
	(*m)[key.String()] = Pair{key, value}
}

func (m *SMap) Do(f func(Pair)) {
	for key, value := range *m {
		f(Pair{key, value})
	}
}

func (m *SMap) Get(key Stringer) ValueOk {
	v, t := (*m)[key.String()]
	return ValueOk{v, t}
}

func (m *SMap) Delete(key Stringer) {
	(*m)[key.String()] = Pair{nil, nil}, false
}

//---
type Inter interface {
	Int() int
}

type IMap map[int]Pair

func NewIMap() IMap {
	return make(map[int]Pair)
}

func (m *IMap) Insert(key Inter, value interface{}) {
	(*m)[key.Int()] = Pair{key, value}
}

func (m *IMap) Do(f func(Pair)) {
	for key, value := range *m {
		f(Pair{key, value})
	}
}

func (m *IMap) Get(key Inter) ValueOk {
	v, t := (*m)[key.Int()]
	return ValueOk{v, t}
}

func (m *IMap) Delete(key Inter) {
	(*m)[key.Int()] = Pair{nil, nil}, false
}
