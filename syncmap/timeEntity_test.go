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
	if myEnt.Utime().IsZero() {
		t.Log(time.Now(), "iszero")
	} else {
		t.Error("utime not zero")
	}
	time.Sleep(time.Second)
	myEnt.BeUsed()
	if myEnt.Utime().IsZero() {
		t.Error("utime is zero")
	} else {
		t.Log(time.Now(), "utime not zero", myEnt.Utime())
	}
}

func TestUpdate(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	oldfreq := myEnt.Chgfreq()
	newStr := "new value"
	myEnt.Update(newStr)
	if v, _ := myEnt.Value(); v == newStr {
		t.Log("new value is: ", v)
	} else {
		t.Error(myEnt.Value())
	}
	if oldfreq == myEnt.Chgfreq() {
		t.Error("chgfreq did not + 1")
	} else if oldfreq == (myEnt.Chgfreq() - 1) {
		t.Log("chgfreq + 1")
	} else {
		t.Error("chgfreq error value")
	}
	if myEnt.Utime().IsZero() {
		t.Error("utime is zero")
	} else {
		t.Log(time.Now(), "utime not zero", myEnt.Utime())
	}
}

func TestChangeDur(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	myEnt.ChangeDur(time.Minute)
	if d := myEnt.Dtime(); d == time.Minute {
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

func TestCtime(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	t.Log(myEnt.Ctime())
}

func TestUtime(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
	if myEnt.Utime().IsZero() {
		t.Log(time.Now(), "iszero")
	} else {
		t.Error("utime not zero")
	}
}

func TestDtime(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
}

func TestGetfreq(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
}

func TestChgfreq(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
}

func TestAddgetfreq(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
}

func TestAddchgfreq(t *testing.T) {
	myEnt := NewTimeEntity("value", 0)
}
