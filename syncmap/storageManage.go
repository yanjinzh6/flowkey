package syncmap

import (
	"errors"
	. "github.com/yanjinzh6/flowkey/tools"
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

func NewStorage(name string, mapType int) Storage {
	return &storage{
		m:       NewSyncMapEnt(),
		name:    name,
		mapType: mapType,
	}
}

func NewStorageManage(name string, mapType int) StorageManage {
	if sManage == nil {
		sManage = storageManage{
			mainMap: NewStorage(name, mapType),
			readMap: make([]Storage),
		}
	}
	return &sManage
}
