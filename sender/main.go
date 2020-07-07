package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:4000")
	if err != nil {
		logrus.Error("ResolveUDPAddr ", err)
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		logrus.Error("DialUDP", err)
	}
	written, err := conn.Write([]byte("yooo"))

	logrus.Info(written)

}
