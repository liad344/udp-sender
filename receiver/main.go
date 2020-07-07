package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:4000")
	if err != nil {
		logrus.Error("ResolveUDPAddr", err)
	}
	listener, err := net.ListenUDP("udp", raddr)

	b := bytes.NewBuffer(make([]byte, 5))
	_, err = listener.Read(b.Bytes())
	if err != nil {
		logrus.Error("ListenUDP", err)
	}
	logrus.Info(b.String())
}
