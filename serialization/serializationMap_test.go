package serialization

import (
	"bytes"
	// "github.com/yanjinzh6/flowkey/tools"
	"testing"
	// "time"
)

var buf *bytes.Buffer
var err error
var sFile SerializationFile

func TestEncode(t *testing.T) {
}

func TestDecode(t *testing.T) {
}

func TestWriteData(t *testing.T) {
}

func TestWriteDataAt(t *testing.T) {
}

func TestReadData(t *testing.T) {
	ReadData("../data/operate.data")
}

func TestSaveManage(t *testing.T) {
	sFile.SetMapData("../data/test.data")
	var i interface{}
	i = 5
	sFile.SaveManage(i)
}
