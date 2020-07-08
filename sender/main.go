package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func main() {
	udpConns := make([]*net.UDPConn, 8)
	for i := 0; i < 8; i++ {
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
	sendFile(udpConns)

}

func sendFile(conn []*net.UDPConn) {
	file, err := os.Open("test.txt")
	if err != nil {
		logrus.Error(err)
	}
	encode(file, conn)
	//_, err = conn.Write(file)
	//if err != nil {
	//	logrus.Error("", err)
	//}
	//logrus.Info("sent ", )
}
