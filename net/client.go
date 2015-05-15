package net

import (
	"github.com/yanjinzh6/flowkey/tools"
	"net"
)

func InitClient() (conn *net.Conn) {
	addr := "127.0.0.1:11223"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	tools.ChErr(err)

	conn, err = net.Dial("tcp", nil, tcpAddr)
	tools.ChErr(err)
	tools.Println("client init success!")
	return
}
