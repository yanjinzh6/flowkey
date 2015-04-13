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
