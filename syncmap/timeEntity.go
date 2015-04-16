package syncmap

import (
	"errors"
	"time"
)

var (
	TimeError = errors.New("text")
)

//timeEntity has value, create time, duration, update time.
type timeEntity struct {
	entity  interface{}
	dtime   time.Duration
	ctime   time.Time
	utime   time.Time
	getfreq int
	chgfreq int
}

//TimeEntity is a interface for timeEntity
type TimeEntity interface {
	IsResident() (b bool)
	IsDie() (b bool)
	BeUsed() (err error)
	Update(value interface{}) (err error)
	ChangeDur(d time.Duration) (err error)
	Value() (val interface{}, err error)
	Ctime() (ctime time.Time)
	Utime() (utime time.Time)
	Dtime() (dtime time.Duration)
	Getfreq() (freq int)
	Chgfreq() (freq int)
	Addgetfreq()
	Addchgfreq()
}

//NewTimeEntity init a timeEntity, used value and duration.
func NewTimeEntity(value interface{}, d time.Duration) TimeEntity {
	return &timeEntity{
		entity: value,
		dtime:  d,
		ctime:  time.Now(),
		//utime:  time.Time{},
		getfreq: 0,
		chgfreq: 0,
	}
}

//IsResident if duration time = 0 ,it is resident
func (t *timeEntity) IsResident() (b bool) {
	if t.dtime == 0 {
		b = true
	} else {
		b = false
	}
	return
}

//IsDie not resident and the update time - now time > duration time
func (t *timeEntity) IsDie() (b bool) {
	if t.IsResident() {
		b = false
	} else {
		curTime := time.Now()
		var mTime time.Time
		if t.utime.IsZero() {
			mTime = t.ctime
		} else {
			mTime = t.utime
		}
		if curTime.Sub(mTime) >= t.dtime {
			b = true
		} else {
			b = false
		}
	}
	return
}

//BeUsed change update time
func (t *timeEntity) BeUsed() (err error) {
	/*if !t.IsResident() {
		t.utime = time.Now()
	}*/
	t.utime = time.Now()
	return
}

//Update update the value and change update time
func (t *timeEntity) Update(value interface{}) (err error) {
	t.entity = value
	t.Addchgfreq()
	t.BeUsed()
	return
}

//ChangeDur change duration time
func (t *timeEntity) ChangeDur(d time.Duration) (err error) {
	t.dtime = d
	t.BeUsed()
	return
}

//Value get entity
func (t *timeEntity) Value() (val interface{}, err error) {
	val = t.entity
	t.Addgetfreq()
	t.BeUsed()
	return
}

//Ctime get Ctime
func (t *timeEntity) Ctime() (dtime time.Time) {
	return t.ctime
}

//Utime get Utime
func (t *timeEntity) Utime() (dtime time.Time) {
	return t.utime
}

//Dtime get Dtime
func (t *timeEntity) Dtime() (dtime time.Duration) {
	return t.dtime
}

//Getfreq get getfreq
func (t *timeEntity) Getfreq() (freq int) {
	return t.getfreq
}

//Chgfreq get chgfreq
func (t *timeEntity) Chgfreq() (freq int) {
	return t.chgfreq
}

//Addgetfreq getfreq + 1
func (t *timeEntity) Addgetfreq() {
	t.getfreq = t.getfreq + 1
}

//Addchgfreq chgfreq + 1
func (t *timeEntity) Addchgfreq() {
	t.chgfreq = t.chgfreq + 1
}
