package serialization

import (
	"bytes"
	"encoding/gob"
)

func Encode(data interface{}) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err = enc.Encode(data)
	return
}

func Decode(data *bytes.Buffer, target interface{}) (err error) {
	buf := bytes.NewBuffer(data.Bytes())
	dec := gob.NewDecoder(buf)
	err = dec.Decode(target)
	return
}
