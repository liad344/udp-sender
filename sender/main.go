package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
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
	write(conn)

}

func write(conn *net.UDPConn) {
	inputDir, err := ioutil.ReadDir("input")
	if err != nil {
		logrus.Error(err)
	}
	for _, file := range inputDir {
		b, err := ioutil.ReadFile("input/" + file.Name())
		if err != nil {
			logrus.Error(err)
		}
		_, err = conn.Write(b)
		if err != nil {
			logrus.Error("", err)
		}
		logrus.Info("sent ", file.Name())
	}

}
