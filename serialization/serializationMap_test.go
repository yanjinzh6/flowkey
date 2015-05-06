package serialization

import (
	"bytes"
	// "encoding/json"
	"github.com/yanjinzh6/flowkey/syncmap"
	"github.com/yanjinzh6/flowkey/tools"
	"testing"
	// "time"
)

var buf *bytes.Buffer

// var bufs []byte
var err error

func TestEncode(t *testing.T) {
	s := syncmap.NewStorageManageS()
	s.Put(1, 1, 0)
	tools.WhatType(s)
	t.Log(&s, s)
	buf, err = Encode(s)
	// bufs, err = json.Marshal(s)
	t.Log(buf, err)
}

func TestDecode(t *testing.T) {
	s := syncmap.NewStorageManageS()
	s.Put(1, 2, 0)
	tools.WhatType(s)
	t.Log(&s, s, err)
	err := Decode(buf, s)
	// json.Unmarshal(bufs, &s)
	t.Log(&s, s, err)
}
