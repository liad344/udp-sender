package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	listen()
}

func listen() {
	listeners := make([]*net.UDPConn, 2)
	for i := 0; i < 2; i++ {
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:400%d", i))
		if err != nil {
			logrus.Error("ResolveUDPAddr", err)
		}
		logrus.Info("listening ", addr)
		Conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			logrus.Error(err)
		}
		listeners[i] = Conn
	}
	decode(listeners)
}
