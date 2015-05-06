package syncmap

import (
	. "github.com/yanjinzh6/flowkey/tools"
	"runtime"
	"sync"
	"time"
)

type Storage struct {
	M       SyncMap
	Name    string
	MapType int
	Size    int
	Index   int
	Key     string
	Rule    string
	Value   string
}

type StorageS struct {
	M       *SyncMapEntS
	Name    string
	MapType int
	Size    int
	Index   int
	Key     string
	Rule    string
	Value   string
}

type storageInt interface {
	Map() (m SyncMap)
	ReSize()
	Add(key, value interface{})
	Del(key interface{})
	Update(key, value interface{})
	SetIndex(index int)
}

type StorageManageS struct {
	MainMap     storageInt
	readMap     []storageInt
	Stat        *statistics
	mytick      *time.Ticker
	AutoClearUp bool
	ClearUpDur  time.Duration
}

type StorageManageSS struct {
	MainMap     *StorageS
	readMap     []storageInt
	Stat        *statistics
	mytick      *time.Ticker
	AutoClearUp bool
	ClearUpDur  time.Duration
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
	StartAutoClearUp() (err error)
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
	// SaveProfile("F:/go_workspace", "cpupprof", "heap", 1)
	// sManage = NewStorageManageUD(time.Second * 2)
	// sManage = NewStorageManageS()
}

func NewStorage(m SyncMap, name string, mapType, size, index int, key, rule, value string) storageInt {
	return &Storage{
		M:       m,
		Name:    name,
		MapType: mapType,
		Size:    size,
		Index:   index,
		Key:     key,
		Rule:    rule,
		Value:   value,
	}
}

func NewStorageS(m *SyncMapEntS, name string, mapType, size, index int, key, rule, value string) *StorageS {
	return &StorageS{
		M:       m,
		Name:    name,
		MapType: mapType,
		Size:    size,
		Index:   index,
		Key:     key,
		Rule:    rule,
		Value:   value,
	}
}

func (s *Storage) Map() (m SyncMap) {
	return s.M
}

func (s *Storage) SetIndex(index int) {
	s.Index = index
}

func (s *Storage) ReSize() {
	if s.M.Size() > s.Size {
		delItem := int(float64(s.Size) * STORAGE_USAGE_AMOUNT)
		for k := range s.M.KeyList() {
			if delItem > 0 {
				s.M.Remove(k)
				delItem = delItem - 1
			}
		}
	}
}

func (s *Storage) Add(key, value interface{}) {
	s.M.Put(key, value, 0)
}

func (s *Storage) Del(key interface{}) {
	s.M.Remove(key)
}

func (s *Storage) Update(key, value interface{}) {
	s.M.Update(key, value)
}

func (s *StorageS) Map() (m SyncMap) {
	return
}

func (s *StorageS) SetIndex(index int) {
	s.Index = index
}

func (s *StorageS) ReSize() {
	if s.M.Size() > s.Size {
		delItem := int(float64(s.Size) * STORAGE_USAGE_AMOUNT)
		for k := range s.M.KeyList() {
			if delItem > 0 {
				s.M.Remove(k)
				delItem = delItem - 1
			}
		}
	}
}

func (s *StorageS) Add(key, value interface{}) {
	s.M.Put(key, value, 0)
}

func (s *StorageS) Del(key interface{}) {
	s.M.Remove(key)
}

func (s *StorageS) Update(key, value interface{}) {
	s.M.Update(key, value)
}

func NewStorageManage() StorageManage {
	lock.Lock()
	if sManage == nil {
		sManage = &StorageManageS{
			MainMap:    NewStorage(NewSyncMapEnt(), "mainMap", STORAGE_MAIN_MAP, 0, 0, "nil", "nil", "nil"),
			readMap:    make([]storageInt, 1),
			Stat:       NewStatistics(),
			mytick:     time.NewTicker(DEFAULT_CLEARUP_TIME),
			ClearUpDur: DEFAULT_CLEARUP_TIME,
		}
		ns := NewStorage(NewSyncMap(), "recentMap", STORAGE_RECENT_USER, STORAGE_DEFAULT_SIZE, 0, "nil", "nil", "nil")
		sManage.AddStorage(ns)
		sManage.StartAutoClearUp()
		Println("init StorageManage, auto ", sManage.IsAutoClearUp())
	} else {
		Println("StorageManage exits, auto ", sManage.IsAutoClearUp())
	}
	lock.Unlock()
	return sManage
}

func NewStorageManageS() StorageManage {
	lock.Lock()
	if sManage == nil {
		sManage = &StorageManageSS{
			MainMap:    NewStorageS(NewSyncMapEntS(), "mainMap", STORAGE_MAIN_MAP, 0, 0, "nil", "nil", "nil"),
			Stat:       NewStatistics(),
			mytick:     time.NewTicker(DEFAULT_CLEARUP_TIME),
			ClearUpDur: DEFAULT_CLEARUP_TIME,
		}
		ns := NewStorage(NewSyncMap(), "recentMap", STORAGE_RECENT_USER, STORAGE_DEFAULT_SIZE, 0, "nil", "nil", "nil")
		sManage.AddStorage(ns)
		sManage.StartAutoClearUp()
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
		sManage = &StorageManageS{
			MainMap:    NewStorage(NewSyncMapEnt(), "mainMap", STORAGE_MAIN_MAP, 0, 0, "nil", "nil", "nil"),
			readMap:    make([]storageInt, 1),
			Stat:       NewStatistics(),
			mytick:     time.NewTicker(d),
			ClearUpDur: d,
		}
		ns := NewStorage(NewSyncMap(), "recentMap", STORAGE_RECENT_USER, STORAGE_DEFAULT_SIZE, 0, "nil", "nil", "nil")
		sManage.AddStorage(ns)
		sManage.StartAutoClearUp()
		Println("init StorageManage, auto ", sManage.IsAutoClearUp())
	} else {
		Println("StorageManage exits, auto ", sManage.IsAutoClearUp())
	}
	lock.Unlock()
	return sManage
}

func NewStorageManageSUD(d time.Duration) StorageManage {
	lock.Lock()
	Println(sManage, &sManage)
	if sManage == nil {
		if d == 0 {
			d = DEFAULT_DURATION_TIME
		}
		sManage = &StorageManageSS{
			MainMap:    NewStorageS(NewSyncMapEntS(), "mainMap", STORAGE_MAIN_MAP, 0, 0, "nil", "nil", "nil"),
			Stat:       NewStatistics(),
			mytick:     time.NewTicker(d),
			ClearUpDur: d,
		}
		ns := NewStorage(NewSyncMap(), "recentMap", STORAGE_RECENT_USER, STORAGE_DEFAULT_SIZE, 0, "nil", "nil", "nil")
		sManage.AddStorage(ns)
		sManage.StartAutoClearUp()
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

func (s *StorageManageS) Get(key interface{}) (val interface{}, err error) {
	Println("Get key:", key)
	if s.readMap != nil && len(s.readMap) > 0 {
		for _, v := range s.readMap {
			val, err = v.Map().Get(key)
			if val != nil {
				err = s.MainMap.Map().Sync(key)
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
		val, err = s.MainMap.Map().Get(key)
		if val != nil {
			s.StorageRule(key, val, STORAGE_MAP_ADD)
		}
	}
	lock.Lock()
	s.Stat.Chgfreq = s.Stat.Getfreq + 1
	lock.Unlock()
	return
}

func (s *StorageManageS) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	Println("Put key:", key, "value:", value, "d:", d)
	val, err = s.MainMap.Map().Put(key, value, d)
	if val != nil {
		s.StorageRule(key, value, STORAGE_MAP_UPD)
	} else {
		s.StorageRule(key, value, STORAGE_MAP_ADD)
	}
	lock.Lock()
	s.Stat.Chgfreq = s.Stat.Chgfreq + 1
	lock.Unlock()
	return
}

func (s *StorageManageS) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}

func (s *StorageManageS) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}

func (s *StorageManageS) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	Println("PutIfAbsent key:", key, "value:", value, "d:", d)
	b, err = s.MainMap.Map().PutIfAbsent(key, value, d)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_ADD)
	}
	return
}

func (s *StorageManageS) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	err = s.MainMap.Map().PutAll(child, d)
	//just add
	for k, v := range child {
		if k != nil && v != nil {
			s.StorageRule(k, v, STORAGE_MAP_ADD)
		}
	}
	return
}

func (s *StorageManageS) Remove(key interface{}) (val interface{}, err error) {
	val, err = s.MainMap.Map().Remove(key)
	if val != nil {
		s.StorageRule(key, val, STORAGE_MAP_DEL)
	}
	return
}

func (s *StorageManageS) RemoveEntry(key, value interface{}) (b bool, err error) {
	b, err = s.MainMap.Map().RemoveEntry(key, value)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_DEL)
	}
	return
}

func (s *StorageManageS) Update(key, value interface{}) (b bool, err error) {
	b, err = s.MainMap.Map().Update(key, value)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_UPD)
	}
	return
}

func (s *StorageManageS) IsEmpty() (b bool) {
	b = s.MainMap.Map().IsEmpty()
	return
}

func (s *StorageManageS) Clear() (err error) {
	err = s.MainMap.Map().Clear()
	for _, v := range s.readMap {
		if v != nil {
			v.Map().Clear()
		}
	}
	return
}

func (s *StorageManageS) ClearUp() (err error) {
	Println("ClearUp")
	err = s.MainMap.Map().ClearUp()
	for _, v := range s.readMap {
		if v != nil {
			v.Map().Clear()
		}
	}
	return
}

func (s *StorageManageS) Size() (size int) {
	Println("Size")
	size = s.MainMap.Map().Size()
	return
}

func (s *StorageManageS) SyncM(key interface{}) (err error) {
	err = s.MainMap.Map().Sync(key)
	return
}

func (s *StorageManageS) AddStorage(ns storageInt) (index int, err error) {
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

func (s *StorageManageS) DelStorage(index int) (b bool, err error) {
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

func (s *StorageManageS) StorageRule(key, value interface{}, t int) (b bool, err error) {
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

func (s *StorageManageS) ClearStorage(index, multiple int) (err error) {
	nowTime := time.Now()
	durTime := int64((nowTime.Sub(s.Stat.Start)) / time.Second)
	allUsage := int64(s.Stat.Getfreq) / durTime
	allSize := s.MainMap.Map().Size()
	size := s.readMap[index].Map().Size()
	ratio := float64(allSize / size * multiple)
	for k := range s.readMap[index].Map().KeyList() {
		//check /s
		ugfreq, _ := s.MainMap.Map().Getfreq(k)
		usage := int64(ugfreq) / durTime
		if float64(usage/allUsage)*STORAGE_USAGE_AMOUNT < ratio {
			s.readMap[index].Map().Remove(k)
		}
	}
	return
}

func (s *StorageManageS) MyTick() (t *time.Ticker) {
	return s.mytick
}

func (s *StorageManageS) ChTick(d time.Duration) {
	s.mytick.Stop()
	s.mytick = time.NewTicker(d)
}

func (s *StorageManageS) IsAutoClearUp() (b bool) {
	return s.AutoClearUp
}

func (s *StorageManageS) StartAutoClearUp() (err error) {
	if !s.AutoClearUp {
		s.AutoClearUp = true
		go func(sManage *StorageManageS) {
			runtime.Gosched()
			sManage.AutoClearUp = true
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

func (s *StorageManageS) StopAutoClearUp() (err error) {
	s.mytick.Stop()
	s.AutoClearUp = false
	return
}

func (s *StorageManageSS) Get(key interface{}) (val interface{}, err error) {
	Println("Get key:", key)
	if s.readMap != nil && len(s.readMap) > 0 {
		for _, v := range s.readMap {
			val, err = v.Map().Get(key)
			if val != nil {
				err = s.MainMap.M.Sync(key)
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
		val, err = s.MainMap.M.Get(key)
		if val != nil {
			s.StorageRule(key, val, STORAGE_MAP_ADD)
		}
	}
	lock.Lock()
	s.Stat.Chgfreq = s.Stat.Getfreq + 1
	lock.Unlock()
	return
}
func (s *StorageManageSS) Put(key, value interface{}, d time.Duration) (val interface{}, err error) {
	Println("Put key:", key, "value:", value, "d:", d)
	val, err = s.MainMap.M.Put(key, value, d)
	if val != nil {
		s.StorageRule(key, value, STORAGE_MAP_UPD)
	} else {
		s.StorageRule(key, value, STORAGE_MAP_ADD)
	}
	lock.Lock()
	s.Stat.Chgfreq = s.Stat.Chgfreq + 1
	lock.Unlock()
	return
}
func (s *StorageManageSS) PutSimple(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, DEFAULT_DURATION_TIME)
	return
}
func (s *StorageManageSS) PutNormal(key, value interface{}) (val interface{}, err error) {
	val, err = s.Put(key, value, 0)
	return
}
func (s *StorageManageSS) PutIfAbsent(key, value interface{}, d time.Duration) (b bool, err error) {
	Println("PutIfAbsent key:", key, "value:", value, "d:", d)
	b, err = s.MainMap.M.PutIfAbsent(key, value, d)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_ADD)
	}
	return
}
func (s *StorageManageSS) PutAll(child map[interface{}]interface{}, d time.Duration) (err error) {
	err = s.MainMap.M.PutAll(child, d)
	//just add
	for k, v := range child {
		if k != nil && v != nil {
			s.StorageRule(k, v, STORAGE_MAP_ADD)
		}
	}
	return
}
func (s *StorageManageSS) Remove(key interface{}) (val interface{}, err error) {
	val, err = s.MainMap.M.Remove(key)
	if val != nil {
		s.StorageRule(key, val, STORAGE_MAP_DEL)
	}
	return
}
func (s *StorageManageSS) RemoveEntry(key, value interface{}) (b bool, err error) {
	b, err = s.MainMap.M.RemoveEntry(key, value)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_DEL)
	}
	return
}
func (s *StorageManageSS) Update(key, value interface{}) (b bool, err error) {
	b, err = s.MainMap.M.Update(key, value)
	if b {
		s.StorageRule(key, value, STORAGE_MAP_UPD)
	}
	return
}
func (s *StorageManageSS) IsEmpty() (b bool) {
	b = s.MainMap.M.IsEmpty()
	return
}
func (s *StorageManageSS) Clear() (err error) {
	err = s.MainMap.M.Clear()
	for _, v := range s.readMap {
		if v != nil {
			v.Map().Clear()
		}
	}
	return
}

func (s *StorageManageSS) ClearUp() (err error) {
	Println("ClearUp")
	err = s.MainMap.M.ClearUp()
	for _, v := range s.readMap {
		if v != nil {
			v.Map().Clear()
		}
	}
	return
}

func (s *StorageManageSS) Size() (size int) {
	Println("Size")
	size = s.MainMap.M.Size()
	return
}

func (s *StorageManageSS) SyncM(key interface{}) (err error) {
	err = s.MainMap.M.Sync(key)
	return
}

func (s *StorageManageSS) AddStorage(ns storageInt) (index int, err error) {
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

func (s *StorageManageSS) DelStorage(index int) (b bool, err error) {
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

func (s *StorageManageSS) StorageRule(key, value interface{}, t int) (b bool, err error) {
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

func (s *StorageManageSS) ClearStorage(index, multiple int) (err error) {
	nowTime := time.Now()
	durTime := int64((nowTime.Sub(s.Stat.Start)) / time.Second)
	allUsage := int64(s.Stat.Getfreq) / durTime
	allSize := s.MainMap.M.Size()
	size := s.readMap[index].Map().Size()
	ratio := float64(allSize / size * multiple)
	for k := range s.readMap[index].Map().KeyList() {
		//check /s
		ugfreq, _ := s.MainMap.M.Getfreq(k)
		usage := int64(ugfreq) / durTime
		if float64(usage/allUsage)*STORAGE_USAGE_AMOUNT < ratio {
			s.readMap[index].Map().Remove(k)
		}
	}
	return
}

func (s *StorageManageSS) MyTick() (t *time.Ticker) {
	return s.mytick
}

func (s *StorageManageSS) ChTick(d time.Duration) {
	s.mytick.Stop()
	s.mytick = time.NewTicker(d)
}

func (s *StorageManageSS) IsAutoClearUp() (b bool) {
	return s.AutoClearUp
}

func (s *StorageManageSS) StartAutoClearUp() (err error) {
	if !s.AutoClearUp {
		s.AutoClearUp = true
		go func(sManage *StorageManageSS) {
			runtime.Gosched()
			sManage.AutoClearUp = true
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

func (s *StorageManageSS) StopAutoClearUp() (err error) {
	s.mytick.Stop()
	s.AutoClearUp = false
	return
}
