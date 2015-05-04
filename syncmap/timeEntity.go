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
	Entity  interface{}
	Dtime   time.Duration
	Ctime   time.Time
	Utime   time.Time
	Getfreq int
	Chgfreq int
}

//TimeEntity is a interface for timeEntity
type TimeEntity interface {
	IsResident() (b bool)
	IsDie() (b bool)
	BeUsed() (err error)
	Update(value interface{}) (err error)
	ChangeDur(d time.Duration) (err error)
	Value() (val interface{}, err error)
	Addgetfreq()
	Addchgfreq()
	GetfreqM() (freq int)
	ChgfreqM() (freq int)
	DtimeM() (d time.Duration)
	CtimeM() (c time.Time)
	UtimeM() (u time.Time)
}

//NewTimeEntity init a timeEntity, used value and duration.
func NewTimeEntity(value interface{}, d time.Duration) TimeEntity {
	return &timeEntity{
		Entity: value,
		Dtime:  d,
		Ctime:  time.Now(),
		//utime:  time.Time{},
		Getfreq: 0,
		Chgfreq: 0,
	}
}

//IsResident if duration time = 0 ,it is resident
func (t *timeEntity) IsResident() (b bool) {
	if t.Dtime == 0 {
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
		if t.Utime.IsZero() {
			mTime = t.Ctime
		} else {
			mTime = t.Utime
		}
		if curTime.Sub(mTime) >= t.Dtime {
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
	t.Utime = time.Now()
	return
}

//Update update the value and change update time
func (t *timeEntity) Update(value interface{}) (err error) {
	t.Entity = value
	t.Addchgfreq()
	t.BeUsed()
	return
}

//ChangeDur change duration time
func (t *timeEntity) ChangeDur(d time.Duration) (err error) {
	t.Dtime = d
	t.BeUsed()
	return
}

//Value get entity
func (t *timeEntity) Value() (val interface{}, err error) {
	val = t.Entity
	t.Addgetfreq()
	t.BeUsed()
	return
}

//Addgetfreq getfreq + 1
func (t *timeEntity) Addgetfreq() {
	t.Getfreq = t.Getfreq + 1
}

//Addchgfreq chgfreq + 1
func (t *timeEntity) Addchgfreq() {
	t.Chgfreq = t.Chgfreq + 1
}

//getfreq
func (t *timeEntity) GetfreqM() (freq int) {
	return t.Getfreq
}

//chgfreq
func (t *timeEntity) ChgfreqM() (freq int) {
	return t.Chgfreq
}

//DtimeM
func (t *timeEntity) DtimeM() (d time.Duration) {
	return t.Dtime
}

//CtimeM
func (t *timeEntity) CtimeM() (c time.Time) {
	return t.Ctime
}

//UtimeM
func (t *timeEntity) UtimeM() (u time.Time) {
	return t.Utime
}
