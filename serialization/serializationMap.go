package serialization

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"github.com/yanjinzh6/flowkey/tools"
	"io"
	"os"
	"path"
)

type SerializationFile struct {
	MapData *os.File
	Operate *os.File
	RWM     *bufio.ReadWriter
	RWO     *bufio.ReadWriter
}

func Encode(data interface{}) (buf *bytes.Buffer, err error) {
	b := make([]byte, 0, 4096)
	buf = bytes.NewBuffer(b)
	enc := gob.NewEncoder(buf)
	err = enc.Encode(data)
	return
}

func EncodeByte(data interface{}) (res []byte, err error) {
	b := make([]byte, 0, 4096)
	buf := bytes.NewBuffer(b)
	enc := gob.NewEncoder(buf)
	err = enc.Encode(data)
	return buf.Bytes(), err
}

func Decode(data *bytes.Buffer, target interface{}) (err error) {
	buf := bytes.NewBuffer(data.Bytes())
	dec := gob.NewDecoder(buf)
	err = dec.Decode(target)
	return
}

func DecodeByte(data []byte, target interface{}) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(target)
	return
}

func WriteData(filePath string, data *bytes.Buffer) (err error) {
	tools.Println("WriteData")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriterSize(file, 4096)
	i, err := write.Write(data.Bytes())
	tools.Println(i, i == data.Len())
	if err != nil {
		tools.Println(err)
	}
	tools.Println(write.Available())
	write.Flush()
	return
}

func WriteDataAt(filePath string, data *bytes.Buffer) {
	tools.Println("WriteDataAt")
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0x0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		tools.Println(err)
	}
	tools.Println("file size: ", info.Size())
	file.Seek(info.Size(), 0)
	// file.Write(data.Bytes())
	write := bufio.NewWriterSize(file, 4096)
	i, err := write.Write(data.Bytes())
	tools.Println(i, i == data.Len(), write.Buffered())
	if err != nil {
		tools.Println(err)
	}
	tools.Println(write.Available())
	write.Flush()
}

func AppendBuffer(bs ...[]byte) (buf *bytes.Buffer) {
	b := make([]byte, 0, 4096)
	buf = bytes.NewBuffer(b)
	for _, bf := range bs {
		buf.Write(bf)
		buf.WriteByte('#')
	}
	buf.WriteRune('\n')
	tools.Println(buf.String())
	return
}

func ReadData(filePath string) {
	tools.Println("ReadData")
	file, err := os.Open(filePath)
	if err != nil {
		tools.Println(err)
	}
	reader := bufio.NewReaderSize(file, 4096)
	var key, value int
	for b, err := reader.ReadBytes('\n'); err == nil; b, err = reader.ReadBytes('\n') {
		tools.Println("line:", b, err)
		bs := bytes.NewBuffer(b)
		sb, err := bs.ReadBytes('#')
		if len(sb) > 0 {
			DecodeByte(sb[:len(sb)-1], &key)
			tools.Println("bytes 1:", string(sb[:len(sb)-1]), err, key)
		} else {
			tools.Println("bytes 1:", string(sb), err, key)
		}
		sb2, err := bs.ReadBytes('#')
		if len(sb2) > 0 {
			DecodeByte(sb2[:len(sb2)-1], &value)
			tools.Println("bytes 2:", string(sb2[:len(sb2)-1]), err, value)
		} else {
			tools.Println("bytes 2:", string(sb2), err, key)
		}
		sb3, err := bs.ReadBytes('#')
		tools.Println("bytes 3:", sb3, err)
	}
}

func (s *SerializationFile) SetMapData(filePath string) (err error) {
	s.MapData, err = os.OpenFile(filePath, os.O_RDWR, 0x0666)
	if err != nil {
		if os.IsExist(err) {
			tools.Println(err)
		} else {
			err = os.MkdirAll(path.Dir(filePath), 0x0666)
			if err != nil {
				tools.Println(err)
			} else {
				s.MapData, err = os.Create(filePath)
				if err != nil {
					tools.Println(err)
				}
			}
		}
	}
	if s.MapData != nil {
		reader := bufio.NewReaderSize(s.MapData, 4096)
		writer := bufio.NewWriterSize(s.MapData, 4096)
		s.RWM = bufio.NewReadWriter(reader, writer)
	}
	tools.Println("SetMapData", s)
	return
}

func (s *SerializationFile) SetOperate(filePath string) (err error) {
	s.Operate, err = os.OpenFile(filePath, os.O_RDWR, 0x0666)
	if err != nil {
		if os.IsExist(err) {
			tools.Println(err)
		} else {
			err = os.MkdirAll(path.Dir(filePath), 0x0666)
			if err != nil {
				tools.Println(err)
			} else {
				s.Operate, err = os.Create(filePath)
				if err != nil {
					tools.Println(err)
				}
			}
		}
	}
	if s.Operate != nil {
		reader := bufio.NewReaderSize(s.Operate, 4096)
		writer := bufio.NewWriterSize(s.Operate, 4096)
		s.RWO = bufio.NewReadWriter(reader, writer)
	}
	tools.Println("SetOperate", s)
	return
}

func (s *SerializationFile) InitManage(target interface{}) (err error) {
	if s.MapData != nil {
		info, err := s.MapData.Stat()
		if err != nil {
			tools.Println(err)
		}
		if info.Size() > 0 {
			p := make([]byte, 4096)
			buf := bytes.NewBuffer(make([]byte, 0, 4096))
			for n, err := s.RWM.Read(p); n != 0 && err == nil; n, err = s.RWM.Read(p) {
				buf.Write(p[:n])
				tools.Println("buf size", buf.Len(), n)
			}
			// p, err := s.RWM.Peek(4096)
			// tools.Println(err)
			// buf.Write(p)
			if err == io.EOF {
				tools.Println(err)
			} else {
				tools.Println(err)
			}
			if buf.Len() > 0 {
				err = Decode(buf, target)
				if err != nil {
					tools.Println(err)
				}
			}
		} else {
			return tools.FileEmplyError
		}
	} else {
		return tools.FileNotFouldError
	}
	return
}

func (s *SerializationFile) SaveManage(sManage interface{}) (err error) {
	if s.MapData != nil {
		bs, err := EncodeByte(sManage)
		if err != nil {
			tools.Println(err)
		}
		s.RWM.Write(bs)
		s.RWM.Flush()
	} else {
		return tools.FileNotFouldError
	}
	return
}

func (s *SerializationFile) Close() {
	s.MapData.Close()
	s.Operate.Close()
}
