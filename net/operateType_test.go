package net

import (
	"github.com/yanjinzh6/flowkey/tools"
	"testing"
)

func TestGet(t *testing.T) {
	ot := NewOperateType()
	ot.Get(1)
	data, err := ot.EncodeByJson()
	ot2 := NewOperateType()
	ot2.DecodeByJson(data)
	tools.Println(ot, ot2, ot == ot2, err)
}
