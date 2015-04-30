package syncmap

import (
	. "github.com/yanjinzh6/flowkey/tools"
	"runtime"
	"sync"
	"time"
)

type storage struct {
	m       SyncMap
	name    string
	mapType int
	size    int
	index   int
	key     string
	rule    string
	value   string
}

type storageInt interface {
	Map() (m SyncMap)
	Name() (name string)
	MapType() (mapType int)
	Size() (size int)
	Index() (index int)
	Key() (key string)
	Rule() (rule string)
	Value() (value string)
	ReSize()
	Add(key, value interface{})
	Del(key interface{})
	Update(key, value interface{})
	SetIndex(index int)
}

type storageManage struct {
	mainMap     storageInt
	readMap     []storageInt
	stat        *statistics
	mytick      *time.Ticker
	autoClearUp bool
	clearUpDur  time.Duration
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
	AddStorage(ns storageInt) (index int, err error)
	DelStorage(index int) (b bool, err error)
	StorageRule(key, value interface{}, t int) (b bool, err error)
	ClearStorage(index, multiple int) (err error)
	MyTick() (t *time.Ticker)
	ChTick(d time.Duration)
	IsAutoClearUp() (b bool)
	AutoClearUp() (err error)
	StopAutoClearUp() (err error)
}

type statistics struct {
	Start   time.Time
	Getfreq int
	Chgfreq int
}

var sManage StorageManage
var lock sync.Mutex

func init() {
	//SaveProfile("F:/go_workspace", "cpupprof", "heap", 1)
	sManage = NewStorageManageUD(time.Second * 2)
}

func NewStorage(m SyncMap, name string, mapType, size, index int, key, rule, value string) storageInt {
	return &storage{
		m:       m,
		name:    name,
		mapType: mapType,
		size:    size,
		index:   index,
		key:     key,
		rule:    rule,
		value:   value,
	}
}

func (s *storage) Map() (m SyncMap) {
	return s.m
}

func (s *storage) Name() (name string) {
	return s.name
}

func (s *storage) MapType() (mapType int) {
	return s.mapType
}

func (s *storage) Size() (size int) {
	return s.size
}

func (s *storage) Index() (index int) {
	return s.index
}

func (s *storage) Key() (key string) {
	return s.key
}

func (s *storage) Rule() (rule string) {
	return s.rule
}

func (s *storage) Value() (value string) {
	return s.value
}

func (s *storage) SetIndex(index int) {
	s.index = index
}

func (s *storage) ReSize() {
	if s.Map().Size() > s.Size() {
		delItem := int(float64(s.Size()) * STORAGE_USAGE_AMOUNT)
		for k := range s.Map().KeyList() {
			if delItem > 0 {
				s.Map().Remove(k)
				delItem = delItem - 1
			}
		}
	}
}

func (s *storage) Add(key, value interface{}) {
	s.Map().Put(key, value, 0)
}

func (s *storage) Del(key interface{}) {
	s.Map().Remove(key)
}

func (s *storage) Update(key, value interface{}) {
	s.Map().Update(key, value)
}

func NewStorageManage() StorageManage {
	lock.Lock()
	if sManage == nil {
		sManage = &storageManage{
			mainMap:    NewStorage(NewSyncMapEnt(), "mainMap", STORAGE_MAIN_MAP, 0, 0, "nil", "nil", "nil"),
			readMap:    make([]storageInt, 1),
			stat:       NewStatistics(),
			mytick:     time.NewTicker(DEFAULT_CLEARUP_TIME),
			clearUpDur: DEFAULT_CLEARUP_TIME,
		}
		ns := NewStorage(NewSyncMap(), "recentMap", STORAGE_RECENT_USER, STORAGE_DEFAULT_SIZE, 0, "nil", "nil", "nil")
		sManage.AddStorage(ns)
		sManage.AutoClearUp()
		Println("init StorageManage, auto ", sManage.IsAutoClearUp())
	} else {
		Println("StorageManage exits, auto ", sManage.IsAutoClearUp())
	}
	lock.Unlock()
	return sManage
}

func NewStorageManageUD(d time.Duration) StorageManage {
	lock.Lock()
	if sManage == nil {
		if d == 0 {
			d = DEFAULT_DURATION_TIME
		}
		sManage = &storageManage{
			mainMap:    NewStorage(NewSyncMapEnt(), "mainMap", STORAGE_MAIN_MAP, 0, 0, "nil", "nil", "nil"),
			readMap:    make([]storageInt, 1),
			stat:       NewStatistics(),
			mytick:     time.NewTicker(d),
			clearUpDur: d,
		}
		ns := NewStorage(NewSyncMap(), "recentMap", STORAGE_RECENT_USER, STORAGE_DEFAULT_SIZE, 0, "nil", "nil", "nil")
		sManage.AddStorage(ns)
		sManage.AutoClearUp()
		Println("init StorageManage, auto ", sManage.IsAutoClearUp())
	} else {
		Println("StorageManage exits, auto ", sManage.IsAutoClearUp())
	}
	lock.Unlock()
	return sManage
}

func NewStatistics() *statistics {
	return &statistics{
		Start:   time.Now(),
		Getfreq: 0,
		Chgfreq: 0,
	}
}

func (s *storageManage) Get(key interface{}) (val interface{}, err error) {
	Println("Get key:", key)
	if s.readMap != nil && len(s.readMap) > 0 {
		for _, v := range s.readMap {
			val, err = v.Map().Get(key)
			if val != nil {
				err = s.mainMap.Map().Sync(key)
				if err != nil && err == NotEntError {
					val = nil
					v.Map().Remove(key)
				} else {
					s.SyncM(key)
				}
				break
			}
		}
	}
	if val == nil && err != NotEntError {
		val, err = s.mainMap.Map().Get(key)
		if val != nil {
			s.StorageRule(key, val, STORAGE_MAP_ADD)
		}
	}
	lock.Lock()
	s.stat.Chgfreq = s.stat.Getfreq + 1
	lock.Unlock()
	return
}

func (s *storageManage) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	Println("Put key:", key, "value:", value, "d:", d)
	val, err = s.mainMap.Map().Put(key, value, d)
	if val != nil {
		s.StorageRule(key, value, STORAGE_MAP_UPD)
	} else {
		s.StorageRule(key, value, STORAGE_MAP_ADD)
	}
	lock.Lock()
	s.stat.Chgfreq = s.stat.Chgfreq + 1
	lock.Unlock()
	return
}

func (s *storageManage) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}

func (s *storageManage) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}

func (s *storageManage) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	Println("PutIfAbsent key:", key, "value:", value, "d:", d)
	b, err = s.mainMap.Map().PutIfAbsent(key, value, d)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_ADD)
	}
	return
}

func (s *storageManage) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	err = s.mainMap.Map().PutAll(child, d)
	//just add
	for k, v := range child {
		if k != nil && v != nil {
			s.StorageRule(k, v, STORAGE_MAP_ADD)
		}
	}
	return
}

func (s *storageManage) Remove(key interface{}) (val interface{}, err error) {
	val, err = s.mainMap.Map().Remove(key)
	if val != nil {
		s.StorageRule(key, val, STORAGE_MAP_DEL)
	}
	return
}

func (s *storageManage) RemoveEntry(key, value interface{}) (b bool, err error) {
	b, err = s.mainMap.Map().RemoveEntry(key, value)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_DEL)
	}
	return
}

func (s *storageManage) Update(key, value interface{}) (b bool, err error) {
	b, err = s.mainMap.Map().Update(key, value)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_UPD)
	}
	return
}

func (s *storageManage) IsEmpty() (b bool) {
	b = s.mainMap.Map().IsEmpty()
	return
}

func (s *storageManage) Clear() (err error) {
	err = s.mainMap.Map().Clear()
	for _, v := range s.readMap {
		if v != nil {
			v.Map().Clear()
		}
	}
	return
}

func (s *storageManage) ClearUp() (err error) {
	Println("ClearUp")
	err = s.mainMap.Map().ClearUp()
	for _, v := range s.readMap {
		if v != nil {
			v.Map().Clear()
		}
	}
	return
}

func (s *storageManage) Size() (size int) {
	Println("Size")
	size = s.mainMap.Map().Size()
	return
}

func (s *storageManage) SyncM(key interface{}) (err error) {
	err = s.mainMap.Map().Sync(key)
	return
}

func (s *storageManage) AddStorage(ns storageInt) (index int, err error) {
	s.readMap = append(s.readMap, ns)
	ok := false
	for i, v := range s.readMap {
		if v == nil {
			index = i
			ns.SetIndex(index)
			s.readMap[i] = ns
			ok = true
			break
		}
	}
	if !ok {
		index = len(s.readMap)
		ns.SetIndex(index)
		s.readMap = append(s.readMap, ns)
	}
	return
}

func (s *storageManage) DelStorage(index int) (b bool, err error) {
	lock.Lock()
	if s.readMap[index] != nil {
		s.readMap[index] = nil
		b = true
	} else {
		b = false
		err = NotEntError
	}
	lock.Unlock()
	return
}

func (s *storageManage) StorageRule(key, value interface{}, t int) (b bool, err error) {
	switch t {
	case STORAGE_MAP_ADD:
	case STORAGE_MAP_DEL:
		for _, v := range s.readMap {
			if v != nil {
				v.Del(key)
			}
		}
	case STORAGE_MAP_UPD:
		for _, v := range s.readMap {
			if v != nil {
				v.Update(key, value)
			}
		}
	default:
		b = false
		err = ParameterTypeError
	}
	return
}

func (s *storageManage) ClearStorage(index, multiple int) (err error) {
	nowTime := time.Now()
	durTime := int64((nowTime.Sub(s.stat.Start)) / time.Second)
	allUsage := int64(s.stat.Getfreq) / durTime
	allSize := s.mainMap.Map().Size()
	size := s.readMap[index].Map().Size()
	ratio := float64(allSize / size * multiple)
	for k := range s.readMap[index].Map().KeyList() {
		//check /s
		ugfreq, _ := s.mainMap.Map().Getfreq(k)
		usage := int64(ugfreq) / durTime
		if float64(usage/allUsage)*STORAGE_USAGE_AMOUNT < ratio {
			s.readMap[index].Map().Remove(k)
		}
	}
	return
}

func (s *storageManage) MyTick() (t *time.Ticker) {
	return s.mytick
}

func (s *storageManage) ChTick(d time.Duration) {
	s.mytick.Stop()
	s.mytick = time.NewTicker(d)
}

func (s *storageManage) IsAutoClearUp() (b bool) {
	return s.autoClearUp
}

func (s *storageManage) AutoClearUp() (err error) {
	if !s.autoClearUp {
		s.autoClearUp = true
		go func(sManage *storageManage) {
			runtime.Gosched()
			sManage.autoClearUp = true
			// for range sManage.mytick.C {
			// 	// Println(t)
			// 	sManage.ClearUp()
			// }
			for {
				select {
				case <-sManage.mytick.C:
					sManage.ClearUp()
				default:
				}
			}
			// mtick := time.Tick(sManage.clearUpDur)
			// for t := range mtick {
			// 	sManage.ClearUp()
			// }
		}(s)
	}
	return
}

func (s *storageManage) StopAutoClearUp() (err error) {
	s.mytick.Stop()
	s.autoClearUp = false
	return
}
