package serialization

import (
	"bytes"
	"github.com/yanjinzh6/flowkey/syncmap"
	"github.com/yanjinzh6/flowkey/tools"
	"testing"
	// "time"
)

var buf *bytes.Buffer
var err error

func TestEncode(t *testing.T) {
	s := syncmap.NewSyncMapEntS()
	s.Put(1, 1, 0)
	t.Log(&s)
	buf, err = Encode(s.Map())
	t.Log(buf, buf.Len(), err)
}

func TestDecode(t *testing.T) {
	s := syncmap.NewSyncMapEntS()
	tools.WhatType(s)
	t.Log(&s, err)
	err := Decode(buf, s)
	t.Log(&s, err)
}
