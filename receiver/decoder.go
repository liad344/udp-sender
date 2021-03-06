package main

import (
	"bytes"
	"github.com/klauspost/reedsolomon"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
)

func decode(udp []*net.UDPConn) {
	shards := make([]io.Reader, 2)
	for i := range udp {
		b := make([]byte, 512)
		n, addr, err := udp[i].ReadFromUDP(b)
		logrus.Info("i ", i, " read from ", addr, " n ", n, " ", string(b))
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info(string(b[511]))
		shards[i] = bytes.NewReader(bytes.Trim(b, string(b[511])))
	}

	enc, err := reedsolomon.NewStream(1, 1)
	if err != nil {
		logrus.Error("Could not create encoder ", err)
	}

	ok, err := enc.Verify(shards)
	if err != nil {
		logrus.Info("Could not verify shards ", err)
	}
	if !ok {
		logrus.Info("Shards lost ", ok)
		out := make([]io.Writer, len(shards))
		err := enc.Reconstruct(shards, out)
		if err != nil {
			logrus.Error("could not reconstruct ", err)
		}
		logrus.Info(enc.Verify(shards))

	} else {
		logrus.Info("All shards received")
	}
	join(enc, shards)
}

func join(enc reedsolomon.StreamEncoder, shards []io.Reader) {
	f, err := os.Create("out.txt")
	if err != nil {
		logrus.Error("could not create file ", err)
	}
	logrus.Info(len(shards))
	err = enc.Join(f, shards, 318)
	if err != nil {
		logrus.Error("Could not join shards ", err)
	}
	logrus.Info("joined")
}
