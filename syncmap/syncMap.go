//package syncmap is a map with sync.RWMutex.
package syncmap

import (
	"errors"
	. "github.com/yanjinzh6/flowkey/tools"
	"sync"
	"time"
)

var (
	NilKeyError   = errors.New("nil key error")
	TimeOutError  = errors.New("the entity is die")
	HasEntError   = errors.New("old data is erased")
	NotEntError   = errors.New("not entity remove")
	NotEqualError = errors.New("map[key] and value are not equal")
)

type syncMap struct {
	m      map[interface{}]interface{}
	rwlock sync.RWMutex
}

type SyncMap interface {
	Get(key interface{}) (val interface{}, err error)
	Put(key, value interface{}, d time.Duration) (val interface{}, err error)
	PutSimple(key, value interface{}) (val interface{}, err error)
	PutNormal(key, value interface{}) (val interface{}, err error)
	PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error)
	PutAll(child map[interface{}]interface{}, d time.Duration) (err error)
	Remove(key interface{}) (val interface{}, err error)
	RemoveEntry(key, value interface{}) (b bool, err error)
	Update(key, value interface{}) (b bool, err error)
	IsEmpty() (b bool)
	Clear() (err error)
	ClearUp() (err error)
	Size() (size int)
	Sync(key interface{}) (err error)
	Getfreq(key interface{}) (freq int, err error)
	Chgfreq(key interface{}) (freq int, err error)
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
	/*if ok, ent := chTimeEntity(val); ok {
		val, err = ent.Value()
	}*/
	s.rwlock.RUnlock()
	return
}

func (s *syncMap) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	val = s.m[key]
	/*if val == nil {
		ent := NewTimeEntity(value, d)
		s.m[key] = ent
	} else {
		if ok, ent := chTimeEntity(val); ok {
			val, err = ent.Value()
		} else {
			s.m[key] = value
		}
	}*/
	s.m[key] = value
	s.rwlock.Unlock()
	return
}

func (s *syncMap) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}

func (s *syncMap) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}

func (s *syncMap) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
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

func (s *syncMap) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
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

func (s *syncMap) Update(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
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

func (s *syncMap) ClearUp() (err error) {
	return
}

func (s *syncMap) Size() (size int) {
	s.rwlock.RLock()
	size = len(s.m)
	s.rwlock.RUnlock()
	return
}

func (s *syncMap) Sync(key interface{}) (err error) {
	return nil
}

func (s *syncMap) Getfreq() (freq int) {
	return 0
}

func (s *syncMap) Getfreq() (freq int) {
	return 0
}

func chTimeEntity(val interface{}) (ok bool, ent TimeEntity) {
	if val == nil {
		return false, nil
	}
	switch value := val.(type) {
	case TimeEntity:
		return true, value
	default:
		return false, nil
	}
}

type syncMapEnt struct {
	m      map[interface{}]TimeEntity
	rwlock sync.RWMutex
}

func NewSyncMapEnt() SyncMap {
	return &syncMapEnt{
		m: make(map[interface{}]TimeEntity),
	}
}

func (s *syncMapEnt) Get(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	/*if s.m[key].IsDie() {
		s.rwlock.Lock()
		s.m[key] = nil
		delete(s.m, key)
		val = nil
		err = TimeOutError
		s.rwlock.Unlock()
	} else {
		s.rwlock.RLock()
		val, err = s.m[key].Value()
		s.rwlock.RUnlock()
	}*/
	s.rwlock.RLock()
	if s.m[key] != nil {
		if s.m[key].IsDie() {
			s.rwlock.RUnlock()
			s.rwlock.Lock()
			s.m[key] = nil
			delete(s.m, key)
			val = nil
			err = TimeOutError
			s.rwlock.Unlock()
			return
		} else {
			val, err = s.m[key].Value()
			s.rwlock.RUnlock()
			s.rwlock.Lock()
			s.m[key].Addgetfreq()
			s.rwlock.Unlock()
		}
	}
	s.rwlock.RUnlock()
	return
}

func (s *syncMapEnt) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	oldEnt := s.m[key]
	if oldEnt == nil {
		ent := NewTimeEntity(value, d)
		s.m[key] = ent
	} else {
		/*if val, _ := oldEnt.Value(); val != value {
			s.m[key].Update(value)
		}
		if oldEnt.Dtime() != d {
			s.m[key].ChangeDur(d)
		}*/
		val, _ = s.m[key].Value()
		s.m[key] = nil
		ent := NewTimeEntity(value, d)
		s.m[key] = ent
		err = HasEntError
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}

func (s *syncMapEnt) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}

func (s *syncMapEnt) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	if s.m[key] == nil {
		b = true
		ent := NewTimeEntity(value, d)
		s.m[key] = ent
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	if child != nil {
		s.rwlock.Lock()
		for k, v := range child {
			if s.m[k] != nil {
				s.m[k] = nil
			}
			ent := NewTimeEntity(v, d)
			s.m[k] = ent
		}
		s.rwlock.Unlock()
	}
	return
}

func (s *syncMapEnt) Remove(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	ent := s.m[key]
	if ent != nil {
		val, err = ent.Value()
		s.m[key] = nil
		delete(s.m, key)
	} else {
		err = NotEntError
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) RemoveEntry(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.m[key]
	if val != nil {
		if v, _ := val.Value(); v == value {
			b = true
			s.m[key] = nil
			delete(s.m, key)
		} else {
			err = NotEqualError
		}
	} else {
		err = NotEntError
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) Update(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.m[key]
	if val != nil {
		b = true
		val.Update(value)
		val.Addchgfreq()
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) IsEmpty() (b bool) {
	s.rwlock.RLock()
	if s.m == nil || len(s.m) == 0 {
		b = true
	} else {
		b = false
	}
	s.rwlock.RUnlock()
	return
}

func (s *syncMapEnt) Clear() (err error) {
	s.rwlock.Lock()
	for k, v := range s.m {
		if v != nil {
			v = nil
			delete(s.m, k)
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) ClearUp() (err error) {
	s.rwlock.Lock()
	for k, v := range s.m {
		if v.IsDie() {
			v = nil
			delete(s.m, k)
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) Size() (size int) {
	s.rwlock.RLock()
	size = len(s.m)
	s.rwlock.RUnlock()
	return
}

func (s *syncMapEnt) Sync(key interface{}) (err error) {
	if !ChKey(key) {
		return NilKeyError
	}
	s.rwlock.Lock()
	val := s.m[key]
	if val != nil {
		val.BeUsed()
	} else {
		err = NotEntError
	}
	s.rwlock.Unlock()
	return
}

func (s *syncMapEnt) Getfreq(key interface{}) (freq int, err error) {
	if !ChKey(key) {
		return NilKeyError
	}
	s.rwlock.RLock()
	val := s.m[key]
	if val != nil {
		val.Getfreq()
	} else {
		err = NotEntError
	}
	s.rwlock.RUnlock()
	return
}

func (s *syncMapEnt) Chgfreq(key interface{}) (freq int, err error) {
	s.rwlock.RLock()
	val := s.m[key]
	if val != nil {
		val.Chgfreq()
	} else {
		err = NotEntError
	}
	s.rwlock.RUnlock()
	return
}
