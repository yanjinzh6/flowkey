//package syncmap is a map with sync.RWMutex.
package syncmap

import (
	. "github.com/yanjinzh6/flowkey/tools"
	"sync"
	"time"
)

type SyncMapS struct {
	M      map[interface{}]interface{}
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
	CreateTime(key interface{}) (t time.Time)
	UpdateTime(key interface{}) (t time.Time)
	KeyList() (keylist []interface{})
}

func NewSyncMap() SyncMap {
	return &SyncMapS{
		M: make(map[interface{}]interface{}),
	}
}

func (s *SyncMapS) Get(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.RLock()
	val = s.M[key]
	/*if ok, ent := chTimeEntity(val); ok {
		val, err = ent.Value()
	}*/
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapS) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	val = s.M[key]
	/*if val == nil {
		ent := NewTimeEntity(value, d)
		s.M[key] = ent
	} else {
		if ok, ent := chTimeEntity(val); ok {
			val, err = ent.Value()
		} else {
			s.M[key] = value
		}
	}*/
	s.M[key] = value
	s.rwlock.Unlock()
	return
}

func (s *SyncMapS) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}

func (s *SyncMapS) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}

func (s *SyncMapS) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	if s.M[key] == nil {
		b = true
		s.M[key] = value
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapS) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	if child != nil {
		s.rwlock.Lock()
		for k, v := range child {
			s.M[k] = v
		}
		s.rwlock.Unlock()
	}
	return
}

func (s *SyncMapS) Remove(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	val = s.M[key]
	if val != nil {
		delete(s.M, key)
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapS) RemoveEntry(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.M[key]
	if val != nil && val == value {
		b = true
		delete(s.M, key)
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapS) Update(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	return
}

func (s *SyncMapS) IsEmpty() (b bool) {
	s.rwlock.RLock()
	if s.M == nil || len(s.M) == 0 {
		b = true
	} else {
		b = false
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapS) Clear() (err error) {
	s.rwlock.Lock()
	for k := range s.M {
		delete(s.M, k)
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapS) ClearUp() (err error) {
	return
}

func (s *SyncMapS) Size() (size int) {
	s.rwlock.RLock()
	size = len(s.M)
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapS) Sync(key interface{}) (err error) {
	return
}

func (s *SyncMapS) Getfreq(key interface{}) (freq int, err error) {
	return
}

func (s *SyncMapS) Chgfreq(key interface{}) (freq int, err error) {
	return
}

func (s *SyncMapS) CreateTime(key interface{}) (t time.Time) {
	return
}

func (s *SyncMapS) UpdateTime(key interface{}) (t time.Time) {
	return
}

func (s *SyncMapS) KeyList() (keylist []interface{}) {
	s.rwlock.RLock()
	keylist = make([]interface{}, s.Size())
	flag := 0
	for k, _ := range s.M {
		keylist[flag] = k
		flag = flag + 1
	}
	s.rwlock.RUnlock()
	return
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

type SyncMapEnt struct {
	M      map[interface{}]TimeEntity
	rwlock sync.RWMutex
}

func NewSyncMapEnt() SyncMap {
	return &SyncMapEnt{
		M: make(map[interface{}]TimeEntity),
	}
}

func (s *SyncMapEnt) Get(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	/*if s.M[key].IsDie() {
		s.rwlock.Lock()
		s.M[key] = nil
		delete(s.M, key)
		val = nil
		err = TimeOutError
		s.rwlock.Unlock()
	} else {
		s.rwlock.RLock()
		val, err = s.M[key].Value()
		s.rwlock.RUnlock()
	}*/
	s.rwlock.RLock()
	if s.M[key] != nil {
		if s.M[key].IsDie() {
			s.rwlock.RUnlock()
			s.rwlock.Lock()
			s.M[key] = nil
			delete(s.M, key)
			val = nil
			err = TimeOutError
			s.rwlock.Unlock()
			return
		} else {
			val, err = s.M[key].Value()
			s.rwlock.RUnlock()
			/*s.rwlock.Lock()
			s.M[key].Addgetfreq()
			s.rwlock.Unlock()*/
			return
		}
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEnt) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	oldEnt := s.M[key]
	if oldEnt == nil {
		ent := NewTimeEntity(value, d)
		s.M[key] = ent
	} else {
		/*if val, _ := oldEnt.Value(); val != value {
			s.M[key].Update(value)
		}
		if oldEnt.DtimeM() != d {
			s.M[key].ChangeDur(d)
		}*/
		val, _ = s.M[key].Value()
		s.M[key] = nil
		ent := NewTimeEntity(value, d)
		s.M[key] = ent
		err = HasEntError
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEnt) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}

func (s *SyncMapEnt) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}

func (s *SyncMapEnt) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	if s.M[key] == nil {
		b = true
		ent := NewTimeEntity(value, d)
		s.M[key] = ent
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEnt) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	if child != nil {
		s.rwlock.Lock()
		for k, v := range child {
			if s.M[k] != nil {
				s.M[k] = nil
			}
			ent := NewTimeEntity(v, d)
			s.M[k] = ent
		}
		s.rwlock.Unlock()
	}
	return
}

func (s *SyncMapEnt) Remove(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	ent := s.M[key]
	if ent != nil {
		val, err = ent.Value()
		s.M[key] = nil
		delete(s.M, key)
	} else {
		err = NotEntError
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEnt) RemoveEntry(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.M[key]
	if val != nil {
		if v, _ := val.Value(); v == value {
			b = true
			s.M[key] = nil
			delete(s.M, key)
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

func (s *SyncMapEnt) Update(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.M[key]
	if val != nil {
		b = true
		val.Update(value)
		//val.Addchgfreq()
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEnt) IsEmpty() (b bool) {
	s.rwlock.RLock()
	if s.M == nil || len(s.M) == 0 {
		b = true
	} else {
		b = false
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEnt) Clear() (err error) {
	s.rwlock.Lock()
	for k, v := range s.M {
		if v != nil {
			v = nil
			delete(s.M, k)
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEnt) ClearUp() (err error) {
	s.rwlock.Lock()
	for k, v := range s.M {
		if v.IsDie() {
			v = nil
			delete(s.M, k)
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEnt) Size() (size int) {
	s.rwlock.RLock()
	size = len(s.M)
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEnt) Sync(key interface{}) (err error) {
	if !ChKey(key) {
		return NilKeyError
	}
	s.rwlock.Lock()
	val := s.M[key]
	if val.IsDie() {
		val = nil
		err = NotEntError
	} else {
		if val != nil {
			val.BeUsed()
		} else {
			err = NotEntError
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEnt) Getfreq(key interface{}) (freq int, err error) {
	if !ChKey(key) {
		return 0, NilKeyError
	}
	s.rwlock.RLock()
	val := s.M[key]
	if val != nil {
		freq = val.GetfreqM()
	} else {
		err = NotEntError
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEnt) Chgfreq(key interface{}) (freq int, err error) {
	if !ChKey(key) {
		return 0, NilKeyError
	}
	s.rwlock.RLock()
	val := s.M[key]
	if val != nil {
		freq = val.ChgfreqM()
	} else {
		err = NotEntError
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEnt) CreateTime(key interface{}) (t time.Time) {
	return s.M[key].CtimeM()
}

func (s *SyncMapEnt) UpdateTime(key interface{}) (t time.Time) {
	return s.M[key].UtimeM()
}

func (s *SyncMapEnt) KeyList() (keylist []interface{}) {
	s.rwlock.RLock()
	keylist = make([]interface{}, s.Size())
	flag := 0
	for k, _ := range s.M {
		keylist[flag] = k
		flag = flag + 1
	}
	s.rwlock.RUnlock()
	return
}

type SyncMapEntS struct {
	M      map[interface{}]*TimeEntityS
	rwlock sync.RWMutex
}

func NewSyncMapEntS() *SyncMapEntS {
	return &SyncMapEntS{
		M: make(map[interface{}]*TimeEntityS),
	}
}

func (s *SyncMapEntS) Get(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.RLock()
	if s.M[key] != nil {
		if s.M[key].IsDie() {
			s.rwlock.RUnlock()
			s.rwlock.Lock()
			s.M[key] = nil
			delete(s.M, key)
			val = nil
			err = TimeOutError
			s.rwlock.Unlock()
			return
		} else {
			val, err = s.M[key].Value()
			s.rwlock.RUnlock()
			return
		}
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEntS) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	oldEnt := s.M[key]
	if oldEnt == nil {
		ent := NewTimeEntityS(value, d)
		s.M[key] = ent
	} else {
		val, _ = s.M[key].Value()
		s.M[key] = nil
		ent := NewTimeEntityS(value, d)
		s.M[key] = ent
		err = HasEntError
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEntS) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}

func (s *SyncMapEntS) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}

func (s *SyncMapEntS) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	if s.M[key] == nil {
		b = true
		ent := NewTimeEntityS(value, d)
		s.M[key] = ent
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEntS) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	if child != nil {
		s.rwlock.Lock()
		for k, v := range child {
			if s.M[k] != nil {
				s.M[k] = nil
			}
			ent := NewTimeEntityS(v, d)
			s.M[k] = ent
		}
		s.rwlock.Unlock()
	}
	return
}

func (s *SyncMapEntS) Remove(key interface{}) (val interface{}, err error) {
	if !ChKey(key) {
		return nil, NilKeyError
	}
	s.rwlock.Lock()
	ent := s.M[key]
	if ent != nil {
		val, err = ent.Value()
		s.M[key] = nil
		delete(s.M, key)
	} else {
		err = NotEntError
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEntS) RemoveEntry(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.M[key]
	if val != nil {
		if v, _ := val.Value(); v == value {
			b = true
			s.M[key] = nil
			delete(s.M, key)
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

func (s *SyncMapEntS) Update(key, value interface{}) (b bool, err error) {
	if !ChKey(key) {
		return false, NilKeyError
	}
	s.rwlock.Lock()
	val := s.M[key]
	if val != nil {
		b = true
		val.Update(value)
		//val.Addchgfreq()
	} else {
		b = false
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEntS) IsEmpty() (b bool) {
	s.rwlock.RLock()
	if s.M == nil || len(s.M) == 0 {
		b = true
	} else {
		b = false
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEntS) Clear() (err error) {
	s.rwlock.Lock()
	for k, v := range s.M {
		if v != nil {
			v = nil
			delete(s.M, k)
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEntS) ClearUp() (err error) {
	s.rwlock.Lock()
	for k, v := range s.M {
		if v.IsDie() {
			v = nil
			delete(s.M, k)
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEntS) Size() (size int) {
	s.rwlock.RLock()
	size = len(s.M)
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEntS) Sync(key interface{}) (err error) {
	if !ChKey(key) {
		return NilKeyError
	}
	s.rwlock.Lock()
	val := s.M[key]
	if val.IsDie() {
		val = nil
		err = NotEntError
	} else {
		if val != nil {
			val.BeUsed()
		} else {
			err = NotEntError
		}
	}
	s.rwlock.Unlock()
	return
}

func (s *SyncMapEntS) Getfreq(key interface{}) (freq int, err error) {
	if !ChKey(key) {
		return 0, NilKeyError
	}
	s.rwlock.RLock()
	val := s.M[key]
	if val != nil {
		freq = val.GetfreqM()
	} else {
		err = NotEntError
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEntS) Chgfreq(key interface{}) (freq int, err error) {
	if !ChKey(key) {
		return 0, NilKeyError
	}
	s.rwlock.RLock()
	val := s.M[key]
	if val != nil {
		freq = val.ChgfreqM()
	} else {
		err = NotEntError
	}
	s.rwlock.RUnlock()
	return
}

func (s *SyncMapEntS) CreateTime(key interface{}) (t time.Time) {
	return s.M[key].CtimeM()
}

func (s *SyncMapEntS) UpdateTime(key interface{}) (t time.Time) {
	return s.M[key].UtimeM()
}

func (s *SyncMapEntS) KeyList() (keylist []interface{}) {
	s.rwlock.RLock()
	keylist = make([]interface{}, s.Size())
	flag := 0
	for k, _ := range s.M {
		keylist[flag] = k
		flag = flag + 1
	}
	s.rwlock.RUnlock()
	return
}
