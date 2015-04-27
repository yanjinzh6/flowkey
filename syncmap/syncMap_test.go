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
	var key interface{}
	var value interface{}
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
	var key interface{}
	var value interface{}
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
	var key interface{}
	var value interface{}
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
	var key interface{}
	var value interface{}
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
	var key interface{}
	var value interface{}
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

func TestUpdateM(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	m.Put(key, value, 0)
	me.Put(key, value, 0)
	value = rand.Intn(9)
	t.Log(m.Update(key, value))
	t.Log(me.Update(key, value))
}

func TestIsEmpty(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	t.Log(m.IsEmpty())
	t.Log(me.IsEmpty())
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	m.Put(key, value, 0)
	me.Put(key, value, 0)
	t.Log(m.IsEmpty())
	t.Log(me.IsEmpty())
}

func TestClear(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	m.Put(key, value, 0)
	me.Put(key, value, 0)
	t.Log(m.IsEmpty())
	t.Log(me.IsEmpty())
	m.Clear()
	me.Clear()
	t.Log(m.IsEmpty())
	t.Log(me.IsEmpty())
}

func TestClearUp(t *testing.T) {
	//m := NewSyncMap()
	//me := NewSyncMapEnt()
}

func TestSize(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	t.Log(m.Size())
	t.Log(me.Size())
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	m.Put(key, value, 0)
	me.Put(key, value, 0)
	t.Log(m.Size())
	t.Log(me.Size())
}

func TestSync(t *testing.T) {
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	me.PutSimple(key, value)
	t.Log(me.UpdateTime(key))
	me.Sync(key)
	t.Log(me.UpdateTime(key))
}

func TestGetfreqM(t *testing.T) {
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	me.PutSimple(key, value)
	t.Log(me.Getfreq(key))
	me.Get(key)
	t.Log(me.Getfreq(key))
}

func TestChgfreqM(t *testing.T) {
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	me.PutSimple(key, value)
	t.Log(me.Chgfreq(key))
	value = rand.Intn(9)
	me.Update(key, value)
	t.Log(me.Chgfreq(key))
}

func TestCreateTime(t *testing.T) {
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	me.PutSimple(key, value)
	t.Log(me.CreateTime(key))
}

func TestUpdateTime(t *testing.T) {
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	me.PutSimple(key, value)
	t.Log(me.UpdateTime(key))
	value = rand.Intn(9)
	me.Update(key, value)
	t.Log(me.UpdateTime(key))
}

func TestKeyList(t *testing.T) {
	m := NewSyncMap()
	me := NewSyncMapEnt()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	m.PutSimple(key, value)
	me.PutSimple(key, value)
	t.Log(m.KeyList())
	t.Log(me.KeyList())
}
