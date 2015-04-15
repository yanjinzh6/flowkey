package syncmap

import (
	"errors"
	. "github.com/yanjinzh6/flowkey/tools"
	"sync"
)

type storage struct {
	m       SyncMap
	name    string
	mapType int
}

type Storage interface {
	MapType() (t int)
}

type storageManage struct {
	mainMap Storage
	readMap []Storage
}

type StorageManage interface {
}

var sManage StorageManage
var lock *sync.Mutex

func NewStorage(name string, mapType int) Storage {
	return &storage{
		m:       NewSyncMapEnt(),
		name:    name,
		mapType: mapType,
	}
}

func NewStorageManage(name string, mapType int) StorageManage {
	lock.Lock()
	if sManage == nil {
		sManage = storageManage{
			mainMap: NewStorage(name, mapType),
			readMap: make([]Storage),
		}
	}
	lock.Unlock()
	return &sManage
}
