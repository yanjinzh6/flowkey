package syncmap

import (
	"errors"
	. "github.com/yanjinzh6/flowkey/tools"
	"sync"
)

var (
	RepeatNameError     = errors.New("add repeat name")
	StorageNotFindError = errors.New("can not find the name of storage list")
)

type storage struct {
	M       SyncMap
	Name    string
	MapType int
	Size    int
	Key     string
	Rule    string
}

type storageInt interface {
	ReSize()
	Add(key, value interface{})
	Del(key interface{})
}

type storageManage struct {
	mainMap storageInt
	readMap []storageInt
	stat    statistics
}

type StorageManage interface {
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
	SyncM(key interface{}) (err error)
	AddStorage(ns storage) (err error)
	DelStorage(name string) (err error)
}

type statistics struct {
	Getfreq int
	Chgfreq int
}

var sManage StorageManage
var lock *sync.Mutex

func NewStorage(m SyncMap, name string, mapType, size int, key, rule string) storageInt {
	return &storage{
		M:       m,
		Name:    name,
		MapType: mapType,
		Size:    size,
		Key:     key,
		Rule:    rule,
	}
}

func (s *storage) ReSize() {
	if s.M.Size() > s.Size {
		s.M.Clear()
	}
}

func (s *storage) Add(key, value interface{}) {
	s.M.Put(key, value, 0)
}

func (s *storage) Del(key interface{}) {
	s.M.Remove(key)
}

func NewStorageManage() StorageManage {
	lock.Lock()
	if sManage == nil {
		sManage = storageManage{
			mainMap: NewStorage(NewSyncMapEnt(), "mainMap", STORAGE_MAIN_MAP, 0, nil, nil),
			readMap: make([]Storage, 1),
			stat:    NewStatistics(),
		}
		ns := NewStorage(NewSyncMap(), "recentMap", STORAGE_RECENT_USER, STORAGE_DEFAULT_SIZE, nil, nil)
		sManage.AddStorage(ns)
	}
	lock.Unlock()
	return &sManage
}

func NewStatistics() *statistics {
	return &statistics{
		Getfreq: 0,
		Chgfreq: 0,
	}
}

func (s *storageManage) Get(key interface{}) (val interface{}, err error) {
	if s.readMap != nil && len(s.readMap) > 0 {
		for i, v := range s.readMap {
			val, err = v.M.Get(key)
			if val != nil {
				err = s.mainMap.M.Sync(key)
				break
			}
		}
	}
	if val == nil {
		val, err = s.mainMap.M.Get(key)
	}
	lock.Lock()
	s.stat.Getfreq = s.stat.Getfreq + 1
	lock.Unlock()
}

func (s *storageManage) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	return
}

func (s *storageManage) PutSimple(key, value interface{}) (val interface{}, err error) {
	return
}

func (s *storageManage) PutNormal(key, value interface{}) (val interface{}, err error) {
	return
}

func (s *storageManage) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	return
}

func (s *storageManage) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	return
}

func (s *storageManage) Remove(key interface{}) (val interface{}, err error) {
	return
}

func (s *storageManage) RemoveEntry(key, value interface{}) (b bool, err error) {
	return
}

func (s *storageManage) Update(key, value interface{}) (b bool, err error) {
	return
}

func (s *storageManage) IsEmpty() (b bool) {
	return
}

func (s *storageManage) Clear() (err error) {
	return
}

func (s *storageManage) ClearUp() (err error) {
	return
}

func (s *storageManage) Size() (size int) {
	return
}

func (s *storageManage) SyncM(key interface{}) (err error) {
	return
}

func (s *storageManage) AddStorage(ns storage) (err error) {
	lock.Lock()
	s.readMap = append(s.readMap, ns)
	lock.Unlock()
	return
}

func (s *storageManage) DelStorage(name string) (err error) {
	return
}
