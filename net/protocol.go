package net

import (
	// "bytes"
	"encoding/binary"
	"github.com/yanjinzh6/flowkey/tools"
)

const (
	MY_HEADER     = "###"
	HEADER_LENGTH = 3
	DATA_LENGTH   = 4
)

func Packet(message []byte) (pack []byte) {
	pack = append(append([]byte(MY_HEADER), IntToBytes(len(message))...), message...)
	tools.Println("pack:", string(pack), "pack size:", len(pack))
	return
}

func Unpack(buf []byte, readerChan chan []byte) (message []byte) {
	length := len(buf)
	tools.Println("now bytes:", string(buf))
	var i int
	for i = 0; i < length; i++ {
		tools.Println("i:", i)
		if length < i+HEADER_LENGTH+DATA_LENGTH {
			tools.Println("length:", length, "i+HEADER_LENGTH+DATA_LENGTH :", i+HEADER_LENGTH+DATA_LENGTH)
			break
		}
		if string(buf[i:i+HEADER_LENGTH]) == MY_HEADER {
			messageLength := BytesToInt(buf[i+HEADER_LENGTH : i+HEADER_LENGTH+DATA_LENGTH])
			if length < i+HEADER_LENGTH+DATA_LENGTH+messageLength {
				break
			}
			tmp := buf[i+HEADER_LENGTH+DATA_LENGTH : i+HEADER_LENGTH+DATA_LENGTH+messageLength]
			readerChan <- tmp
			i += HEADER_LENGTH + DATA_LENGTH + messageLength - 1
			tools.Println("message length:", messageLength, "tmp:", string(tmp), "new i:", i)
		}
	}
	if i == length {
		message = make([]byte, 0)
	}
	message = buf[i:]
	tools.Println("return tmp:", string(message), "size:", len(message))
	return
}

func IntToBytes(n int) (data []byte) {
	x := uint32(n)
	/*buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, x)*/
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, x)
	tools.Println("message length:", x, n, data, len(data))
	return
}

func BytesToInt(data []byte) (n int) {
	/*buffer := bytes.NewBuffer(data)
	var x int32
	err := binary.Read(buffer, binary.BigEndian, &x)*/
	n = int(binary.BigEndian.Uint32(data))
	tools.Println("byte to int:", n)
	return
}
