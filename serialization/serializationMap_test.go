package serialization

import (
	"bytes"
	"github.com/yanjinzh6/flowkey/syncmap"
	"testing"
	"time"
)

var buf *bytes.Buffer
var err error

func TestEncode(t *testing.T) {
	s := syncmap.NewStorageManage()
	for i := 0; i < 10; i++ {
		s.Put(i, (i + 1), time.Second*10)
	}
	t.Log(s)
	buf, err = Encode(s)
	t.Log(buf, buf.Len(), err)
}

func TestDecode(t *testing.T) {
	s := syncmap.NewStorageManage()
	t.Log(s, err)
	s.Clear()
	err := Decode(buf, s)
	t.Log(s, err)
	for i := 0; i < 10; i++ {
		t.Log(s.Get(i))
	}
}
