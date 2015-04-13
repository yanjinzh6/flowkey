//package syncmap is a map with sync.RWMutex.
package syncmap

import (
	"errors"
	. "github.com/yanjinzh6/flowkey/tools"
	"sync"
	"time"
)

const (
	DEFAULT_DURATION_TIME = time.Minute * 30
)

var (
	NilKeyError = errors.New("nil key error")
	NilError    = errors.New("text")
)

type syncMap struct {
	m      map[interface{}]interface{}
	rwlock sync.RWMutex
}

type SyncMap interface {
	Get(key interface{}) (val interface{}, err error)
	Put(key, value interface{}) (val interface{}, err error)
	PutIfAbsent(key, value interface{}) (b bool, err error)
	PutAll(child map[interface{}]interface{}) (err error)
	Remove(key interface{}) (val interface{}, err error)
	RemoveEntry(key, value interface{}) (b bool, err error)
	IsEmpty() (b bool)
	Clear() (err error)
	Size() (size int)
}

func NewSyncMap() SyncMap {
	return &syncMap{
		m: make(map[interface{}]interface{}),
	}
}

func (s *syncMap) Get(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.RLock()
	val = s.m[key]
	if ok, ent := chTimeEntity(val); ok {
		val, err = ent.Value()
	}
	s.rwlock.RUnlock()
	return
}

func (s *syncMap) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	val = s.m[key]
	if val == nil {
		ent := NewTimeEntity(value, d)
		s.m[key] = ent
	} else {
		if ok, ent := chTimeEntity(val); ok {
			val, err = ent.Value()
		} else {
			s.m[key] = value
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMap) PutSimple(key, value interface{}) (val interface{}, err error) {
	Put(key, value, DEFAULT_DURATION_TIME)
}

func (s *syncMap) PutNormal(key, value interface{}) (val interface{}, err error) {
	Put(key, value, DEFAULT_DURATION_TIME)
}

func (s *syncMap) PutIfAbsent(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	if s.m[key] == nil {
		b = true
		s.m[key] = value
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMap) PutAll(child map[interface{}]interface{}) (err error) {
	if child != nil {
		s.rwlock.Lock()
		for k, v := range child {
			s.m[k] = v
		}
		s.rwlock.Unlock()
	}
	return
}

func (s *syncMap) Remove(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	val = s.m[key]
	if val != nil {
		delete(s.m, key)
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMap) RemoveEntry(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.m[key]
	if val != nil && val == value {
		b = true
		delete(s.m, key)
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMap) IsEmpty() (b bool) {
	s.rwlock.RLock()
	if s.m == nil || len(s.m) == 0 {
		b = true
	} else {
		b = false
	}
	s.rwlock.RUnlock()
	return
}

func (s *syncMap) Clear() (err error) {
	s.rwlock.Lock()
	for k := range s.m {
		delete(s.m, k)
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMap) Size() (size int) {
	s.rwlock.RLock()
	size = len(s.m)
	s.rwlock.RUnlock()
	return
}

func chTimeEntity(val interface{}) (ok bool, ent TimeEntity) {
	if val == nil {
		return false
	}
	switch value := val.(type) {
	case TimeEntity:
		return true, value
	default:
		return false, nil
	}
}
