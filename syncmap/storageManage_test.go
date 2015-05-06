package syncmap

import (
	"fmt"
	"math/rand"
	// "runtime"
	"testing"
	"time"
)

func TestNewStorageManage(t *testing.T) {
	// runtime.GOMAXPROCS(2)
	s := NewStorageManageSUD(time.Second * 5)
	t.Log(s)
	go func() {
		s2 := NewStorageManageSUD(time.Second * 5)
		t.Log(s2)
		for i := 0; i < 100; i++ {
			fmt.Println(i, s2.Size())
			if i == 20 {
				s2.ChTick(time.Second * 8)
			}
			time.Sleep(time.Millisecond * 300)
		}
	}()
	go func() {
		s3 := NewStorageManageSUD(time.Second * 5)
		t.Log(s3)
		var key interface{}
		var value interface{}
		for i := 0; i < 100; i++ {
			key = rand.Intn(999)
			value = rand.Intn(999)
			time.Sleep(time.Millisecond * 200)
			s3.Put(key, value, time.Second*10)
		}
	}()
	time.Sleep(time.Second * 30)
}

func TestNewStorageManageUD(t *testing.T) {
	s := NewStorageManageUD(time.Second * 5)
	s2 := NewStorageManageUD(time.Second * 5)
	t.Log(s)
	t.Log(s2)
	// time.Sleep(time.Second * 10)
}

func TestGetS(t *testing.T) {
	s := NewStorageManage()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	s.Put(key, value, 0)
	t.Log(s.Get(key))
}

func TestPutS(t *testing.T) {
	s := NewStorageManage()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	t.Log(s.Put(key, value, 0))
	t.Log(s.Size())
}

func TestPutSimpleS(t *testing.T) {
	s := NewStorageManage()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	t.Log(s.PutSimple(key, value))
	t.Log(s.Size())
}

func TestPutNormalS(t *testing.T) {
	s := NewStorageManage()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	t.Log(s.PutNormal(key, value))
	t.Log(s.Size())
}

func TestPutIfAbsentS(t *testing.T) {
	s := NewStorageManage()
	var key interface{}
	var value interface{}
	key = rand.Intn(9)
	value = rand.Intn(9)
	t.Log(s.PutIfAbsent(key, value, 0))
	t.Log(s.Size())
}

func TestPutAllS(t *testing.T) {
	s := NewStorageManage()
	var key interface{}
	var value interface{}
	maps := make(map[interface{}]interface{})
	for i := 0; i < 10; i++ {
		key = rand.Intn(999)
		value = rand.Intn(999)
		maps[key] = value
	}
	t.Log(s.PutAll(maps, 0))
}

func TestRemoveS(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestRemoveEntryS(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestUpdateS(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestIsEmptyS(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
	t.Log(s.IsEmpty())
}

func TestClearS(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.IsEmpty())
	s.Clear()
	t.Log(s.IsEmpty())
}

func TestClearUpS(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
	s.ClearUp()
	t.Log(s.Size())
}

func TestSizeS(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestSyncM(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestAddStorage(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestDelStorage(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestStorageRule(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}

func TestClearStorage(t *testing.T) {
	s := NewStorageManage()
	t.Log(s.Size())
}
