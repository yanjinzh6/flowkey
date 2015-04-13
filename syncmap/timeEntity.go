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
	entity interface{}
	dtime  time.Duration
	ctime  time.Time
	utime  time.Time
}

//TimeEntity is a interface for timeEntity
type TimeEntity interface {
	IsResident() (b bool)
	IsDie() (b bool)
	BeUsed() (err error)
	Update(value interface{}) (err error)
	ChangeDur(d time.Duration) (err error)
	Value() (val interface{}, err error)
}

//NewTimeEntity init a timeEntity, used value and duration.
func NewTimeEntity(value interface{}, d time.Duration) TimeEntity {
	return &timeEntity{
		entity: value,
		dtime:  d,
		ctime:  time.Now(),
		//utime:  time.Time{},
	}
}

func (t *timeEntity) IsResident() (b bool) {
	if t.dtime == 0 {
		b = true
	} else {
		b = false
	}
	return
}

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

func (t *timeEntity) BeUsed() (err error) {
	if !t.IsResident() {
		t.utime = time.Now()
	}
	return
}

func (t *timeEntity) Update(value interface{}) (err error) {
	t.entity = value
	t.BeUsed()
	return
}

func (t *timeEntity) ChangeDur(d time.Duration) (err error) {
	t.dtime = d
	t.BeUsed()
	return
}

func (t *timeEntity) Value() (val interface{}, err error) {
	val = t.entity
	return
}
