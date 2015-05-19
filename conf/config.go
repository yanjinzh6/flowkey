package conf

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/yanjinzh6/flowkey/tools"
	"os"
	"path"
	"time"
)

type Config struct {
	MC MapConf
	SC StorageConf
	NC NetConf
}

type MapConf struct {
	DuractionTime time.Duration
	ClearUpTime   time.Duration
	StorageTime   time.Duration
	SMC           SubMapConf
}

type SubMapConf struct {
	DefaultSize int
	UsageAmount float64
}

type StorageConf struct {
	SerializationFile string
}

type NetConf struct {
	Addr       string
	Port       int
	BufferSize int
}

var config *Config

func InitConfig(filePath string) (config *Config) {
	if filePath == "" {
		filePath = tools.CONFIG_FILE_PATH
	}
	if config == nil {
		file, err := InitFile(filePath)
		tools.ChErr(err)
		reader := bufio.NewReaderSize(file, tools.DEFAULT_BUFFER_SIZE)
		buf := make([]byte, 0, tools.DEFAULT_BUFFER_SIZE)
		bufs := bytes.NewBuffer(buf)
		for {
			buf, err := reader.ReadBytes('\n')
			tools.Println(buf)
			bufs.Write(buf)
			if err != nil {
				break
			}
		}
		tools.Println(bufs.Len())
		if bufs.Len() == 0 {
			//nothing
			config = &Config{
				MC: MapConf{DuractionTime: time.Minute * 30, ClearUpTime: time.Minute * 5, StorageTime: time.Minute * 10, SMC: SubMapConf{DefaultSize: 10000, UsageAmount: 0.5}},
				SC: StorageConf{SerializationFile: "../data/map"},
				NC: NetConf{Addr: "0.0.0.0", Port: 11223, BufferSize: 4096},
			}
			tools.Println(config)
			writer := bufio.NewWriterSize(file, tools.DEFAULT_BUFFER_SIZE)
			buf, err = json.Marshal(&config)
			tools.ChErr(err)
			n, err := writer.Write(buf)
			tools.Println(n, err)
			writer.Flush()
		} else {
			err = json.Unmarshal(bufs.Bytes(), &config)
			tools.ChErr(err)
		}
	}
	return config
}

func InitFile(filePath string) (file *os.File, err error) {
	file, err = os.OpenFile(filePath, os.O_RDWR, 0x0666)
	if err != nil {
		if os.IsExist(err) {
			tools.Println(err)
		} else {
			err = os.MkdirAll(path.Dir(filePath), 0x0666)
			if err != nil {
				tools.Println(err)
			} else {
				file, err = os.Create(filePath)
				if err != nil {
					tools.Println(err)
				}
			}
		}
	}
	return
}
