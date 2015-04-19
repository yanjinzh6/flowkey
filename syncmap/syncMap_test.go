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

func TestNewSyncMapEnt(t *testing.T) {
	m := NewSyncMapEnt()
	if m != nil {
		t.Log("NewSyncMapEnt(): ", m)
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
func initMap(s SyncMap) {
	var key interface{}
	var value interface{}
	for i := 0; i < 10; i++ {
		key = rand.Intn(999)
		value = rand.Intn(999)
		s.Put(key, value, 0)
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
	me := NewSyncMapEnt()
	me.Put(key, value, 0)
	v, err = me.Get(key)
	if err == nil && value == v {
		t.Log("Get(): ", v)
	} else {
		t.Error("value: ", value, " v: ", v, "err: ", err)
	}
	v, err = me.Get(maps)
	if err != nil {
		t.Log("Get(): test used map key: ", err)
	} else {
		t.Error("err: ")
	}
}

func TestPut(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	for i := 0; i < 10; i++ {
		key = rand.Intn(999)
		value = rand.Intn(999)
		t.Log(m.Put(key, value, 0))
		t.Log(me.Put(key, value, 0))
	}
	t.Log(m, me)
}

func TestPutSimple(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	for i := 0; i < 10; i++ {
		key = rand.Intn(999)
		value = rand.Intn(999)
		t.Log(m.PutSimple(key, value))
		t.Log(me.PutSimple(key, value))
	}
	t.Log(m, me)
}

func TestPutNormal(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	for i := 0; i < 10; i++ {
		key = rand.Intn(999)
		value = rand.Intn(999)
		t.Log(m.PutNormal(key, value))
		t.Log(me.PutNormal(key, value))
	}
	t.Log(m, me)
}

func TestPutIfAbsent(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	for i := 0; i < 10; i++ {
		key = 1
		value = rand.Intn(999)
		t.Log(m.PutIfAbsent(key, value, 0))
		t.Log(me.PutIfAbsent(key, value, 0))
	}
	t.Log(m, me)
}

func TestPutAll(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	maps := make(map[interface{}]interface{})
	for i := 0; i < 10; i++ {
		key = rand.Intn(999)
		value = rand.Intn(999)
		maps[key] = value
	}
	t.Log(m.PutAll(maps, 0))
	t.Log(me.PutAll(maps, 0))
}

func TestRemove(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	t.Log(m.Put(key, value, 0))
	t.Log(me.Put(key, value, 0))
	t.Log(m.Remove(key))
	t.Log(me.Remove(key))
	t.Log(m.Remove(11))
	t.Log(me.Remove(11))
}

func TestRemoveEntry(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	t.Log(m.Put(key, value, 0))
	t.Log(me.Put(key, value, 0))
	t.Log(m.RemoveEntry(key, 11))
	t.Log(me.RemoveEntry(key, 11))
	t.Log(m.RemoveEntry(11, 11))
	t.Log(me.RemoveEntry(11, 11))
	t.Log(m.RemoveEntry(key, value))
	t.Log(me.RemoveEntry(key, value))
}

func TestUpdate(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestIsEmpty(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestClear(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestClearUp(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestSize(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestSync(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestGetfreq(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestChgfreq(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestCreateTime(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestUpdateTime(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}

func TestKeyList(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
}
