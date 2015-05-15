package net

import (
	// "bytes"
	"github.com/yanjinzh6/flowkey/syncmap"
	"github.com/yanjinzh6/flowkey/tools"
	"net"
)

func init() {
	go InitServer()
	sManage = syncmap.NewStorageManageS()
}

var (
	sManage syncmap.StorageManage
)

type MyServer struct {
	Addr string
}

func func_name() {

}

func InitServer() {
	addr := ":11223"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	tools.ChErr(err)
	lister, err := net.ListenTCP("tcp", tcpAddr)
	tools.ChErr(err)
	tools.Println("server start success")
	for {
		conn, err := lister.Accept()
		if err != nil {
			tools.Println(err)
			continue
		}
		go handleWorker(conn)
	}
}

func handleWorker(conn net.Conn) {
	defer conn.Close()
	readerChan := make(chan []byte)
	go Reader(readerChan, conn)
	tools.Println(conn.RemoteAddr(), " connected!")
	tmp := make([]byte, 0)
	buf := make([]byte, 8)
	for {
		n, err := conn.Read(buf)
		tools.Println(conn)
		if err != nil {
			tools.Println("conn read error:", err, n, buf)
			return
		}
		tmp = Unpack(append(tmp, buf[:n]...), readerChan)
		tools.Println("more message:", tmp)
		tools.Println("string:", string(buf[:n]))
	}
}

func Reader(readerChan chan []byte, conn net.Conn) {
	for {
		select {
		case data := <-readerChan:
			// decode data and query map and return val
			tools.Println("full message:", string(data))
			ot := NewOperateType()
			ot.DecodeByJson(data)
			switch ot.T {
			case tools.NET_TYPE_GET:
				val, err := sManage.Get(ot.Key)
				ot.Val, ot.Err = val, err
				/*buf, err := ot.EncodeByJson()
				tools.ChErr(err)
				Send(conn, buf)*/
			case tools.NET_TYPE_PUT:
				val, err := sManage.Put(ot.Key, ot.Val, ot.D)
				ot.Val, ot.Err = val, err
			}
			buf, err := ot.EncodeByJson()
			tools.ChErr(err)
			Send(conn, buf)
		}
	}
}

func Send(conn *net.Conn, data []byte) (n int, err error) {
	n, err = conn.Write(Packet(data))
	tools.Println("send size:", n, string(data), err)
	return
}

func Close(conn *net.Conn) (err error) {
	err = conn.Close()
	return
}
