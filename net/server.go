package net

import (
	"github.com/yanjinzh6/flowkey/tools"
	"net"
)

func init() {

}

func InitServer() {
	addr := ":11223"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	tools.ChErr(err)
	lister, err := net.ListenTCP("tcp", tcpAddr)
	tools.ChErr(err)
	for {
		conn, err := lister.Accept()
		if err != nil {
			tools.Println(err)
			continue
		}
		go handleWorker(conn)
	}
}

func handleWorker(conn *net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
