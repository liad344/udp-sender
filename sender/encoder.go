package main

import (
	"github.com/klauspost/reedsolomon"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
)

func encode(file *os.File, udpConnOUT []*net.UDPConn) {
	enc, err := reedsolomon.NewStream(4, 4)
	if err != nil {
		logrus.Error("Could not create encoder ", err)
	}

	data := make([]io.Writer, 4)
	for i := range data {
		data[i] = udpConnOUT[i]
	}
	err = enc.Split(file, data, 318)
	if err != nil {
		logrus.Error("Could not split ", err)
	}
	parity := make([]io.Writer, 4)
	for i := range parity {
		parity[i] = udpConnOUT[1+i]
	}
	input := make([]io.Reader, 4)
	for i := range data {
		file, err := os.Open("test.txt")
		if err != nil {
			logrus.Error(err)
		}
		input[i] = file
	}
	err = enc.Encode(input, parity)
	if err != nil {
		logrus.Error("Could not encode parity ", err)
	}
	for i := range udpConnOUT {
		err = udpConnOUT[i].Close()
		if err != nil {
			logrus.Error("could not close connection ", err)
		}
	}
	logrus.Info("sent")
}
