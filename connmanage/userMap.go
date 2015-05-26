package connmanage

import (
	"github.com/yanjinzh6/flowkey/syncmap"
	"sync"
	"time"
)

type User struct {
	Id         string
	Name       string
	Password   string
	CreateTime time.Time
	UseTime    time.Time
}

type MapManage struct {
	UserMap    syncmap.SyncMapS
	UserMapMap syncmap.SyncMapS
}

var manage MapManage
var lock sync.Locker

func NewMapManage() {
	lock.Lock()
	lock.Unlock()
}
