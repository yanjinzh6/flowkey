package syncmap

import (
	"math/rand"
	"testing"
)

func TestNewSyncMap(t *testing.T) {
	m := NewSyncMap()
	if m != nil {
		t.Log("NewSyncMap(): ", m)
	} else {
		t.Error("nil")
	}
	myEnt := NewTimeEntity("value", 0)
	m.Put(1, myEnt, 0)
	ent, _ := m.Get(1)
	if val, ok := ent.(TimeEntity); ok {
		t.Log(val)
	}
}

func TestGet(t *testing.T) {
	m := NewSyncMap()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	m.Put(key, value, 0)
	v, err := m.Get(key)
	if err == nil && value == v {
		t.Log("Get(): ", v)
	} else {
		t.Error("value: ", value, " v: ", v, "err: ", err)
	}
	maps := make(map[interface{}]interface{})
	v, err = m.Get(maps)
	if err != nil {
		t.Log("Get(): test used map key: ", err)
	} else {
		t.Error("err: ")
	}
}
