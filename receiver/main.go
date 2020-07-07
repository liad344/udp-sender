package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:4000")
	if err != nil {
		logrus.Error("ResolveUDPAddr", err)
	}
	logrus.Info("listening")
	listen(addr)
}

func listen(addr *net.UDPAddr) {

	listener, err := net.ListenUDP("udp", addr)
	b := bytes.NewBuffer(make([]byte, 64))
	for {
		_, err = listener.Read(b.Bytes())
		if err != nil {
			logrus.Error("ListenUDP", err)
		}
		err := ioutil.WriteFile("output/test.txt", b.Bytes(), 0644)
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info("got file")
	}
}
