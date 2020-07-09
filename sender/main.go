package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func main() {
	udpConns := make([]*net.UDPConn, 2)
	for i := 0; i < 2; i++ {
		raddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:400%d", i))
		if err != nil {
			logrus.Error("ResolveUDPAddr ", err)
		}
		logrus.Info("dailed ", raddr)
		conn, err := net.DialUDP("udp", nil, raddr)
		if err != nil {
			logrus.Error("DialUDP", err)
		}
		udpConns[i] = conn
	}
	err := sendFile(udpConns)
	if err != nil {
		logrus.Error(err)
	}

}

func sendFile(conn []*net.UDPConn) error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}
	return encode(file, conn)
}
