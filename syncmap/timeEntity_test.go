package syncmap

import (
	"testing"
	"time"
)

func TestNewTimeEntity(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if myEnt != nil && myEnt.IsResident() {
		t.Log("init a resident entity, value is: ")
		t.Log(myEnt.Value())
	}
}

func TestIsResident(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if myEnt.IsResident() {
		t.Log("0 dtime")
	} else {
		t.Error("0 dtime")
	}
	myEnt = NewTimeEntity("value", time.Second)
	if !myEnt.IsResident() {
		t.Log("dtime")
	}
}

func TestIsDie(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if !myEnt.IsDie() {
		t.Log("not die")
	} else {
		t.Error("is die")
	}
	myEnt = NewTimeEntity("value", time.Second)
	if !myEnt.IsDie() {
		t.Log("1s now not die")
	} else {
		t.Error("is die")
	}
	t.Log("sleep 1s ..")
	time.Sleep(time.Second)
	if myEnt.IsDie() {
		t.Log("after 1s is die")
	} else {
		t.Error("not die")
	}
}

func TestBeUser(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if myEnt.UtimeM().IsZero() {
		t.Log(time.Now(), "iszero")
	} else {
		t.Error("utime not zero")
	}
	time.Sleep(time.Second)
	myEnt.BeUsed()
	if myEnt.UtimeM().IsZero() {
		t.Error("utime is zero")
	} else {
		t.Log(time.Now(), "utime not zero", myEnt.UtimeM())
	}
}

func TestUpdate(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	oldfreq := myEnt.ChgfreqM()
	newStr := "new value"
	myEnt.Update(newStr)
	if v, _ := myEnt.Value(); v == newStr {
		t.Log("new value is: ", v)
	} else {
		t.Error(myEnt.Value())
	}
	if oldfreq == myEnt.ChgfreqM() {
		t.Error("chgfreq did not + 1")
	} else if oldfreq == (myEnt.ChgfreqM() - 1) {
		t.Log("chgfreq + 1")
	} else {
		t.Error("chgfreq error value")
	}
	if myEnt.UtimeM().IsZero() {
		t.Error("utime is zero")
	} else {
		t.Log(time.Now(), "utime not zero", myEnt.UtimeM())
	}
}

func TestChangeDur(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	myEnt.ChangeDur(time.Minute)
	if d := myEnt.DtimeM(); d == time.Minute {
		t.Log("change dur ", d)
	} else {
		t.Error(d, "is not time.Minute")
	}
}

func TestValue(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if v, _ := myEnt.Value(); v == "value" {
		t.Log("value true")
	} else {
		t.Error("the value is ", v)
	}
}

func TestAddgetfreq(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if myEnt.GetfreqM() == 0 {
		t.Log("getfreq", myEnt.GetfreqM())
	} else {
		t.Error(myEnt.GetfreqM())
	}
	myEnt.Addgetfreq()
	if myEnt.GetfreqM() == 1 {
		t.Log("getfreq", myEnt.GetfreqM())
	} else {
		t.Error(myEnt.GetfreqM())
	}
}

func TestAddchgfreq(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if myEnt.ChgfreqM() == 0 {
		t.Log("chgfreq", myEnt.ChgfreqM())
	} else {
		t.Error(myEnt.ChgfreqM())
	}
	myEnt.Addchgfreq()
	if myEnt.ChgfreqM() == 1 {
		t.Log("chgfreq", myEnt.ChgfreqM())
	} else {
		t.Error(myEnt.ChgfreqM())
	}
}
