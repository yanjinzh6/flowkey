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
	buf, err = Encode(s)
	t.Log(buf, err)
}

func TestDecode(t *testing.T) {
	s := syncmap.NewStorageManage()
	err := Decode(buf, s)
	t.Log(s, err)
}
