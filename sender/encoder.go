package main

import (
	"bytes"
	"github.com/klauspost/reedsolomon"
	"github.com/pkg/errors"
	"io"
	"net"
	"os"
)

func encode(file *os.File, udpConns []*net.UDPConn) error {
	enc, err := reedsolomon.NewStream(1, 1)
	if err != nil {
		return err
	}
	dataShardsMW := make([]io.Writer, 1)
	dataShardsBuff := make([]*bytes.Buffer, 1)
	dataShards := make([]io.Writer, 1)
	parityShards := make([]io.Writer, 1) //Parity shards will be written to -> udpConn

	for i := range dataShardsMW {
		dataShardsBuff[i] = bytes.NewBuffer(make([]byte, 64))
		dataShards[i] = udpConns[i]
		parityShards[i] = udpConns[1+i]
		dataShardsMW[i] = io.MultiWriter(dataShards[i], dataShardsBuff[i])
	}
	err = enc.Split(file, dataShardsMW, 318)
	if err != nil {
		return errors.WithMessage(err, "Could not split ")
	}

	parityReader := make([]io.Reader, 1)
	for i := range dataShardsBuff {
		parityReader[i] = dataShardsBuff[i]
	}
	err = enc.Encode(parityReader, parityShards)
	if err != nil {
		return errors.WithMessage(err, "Could not encode parity")
	}
	return closeUDPConnection(udpConns)
}

func closeUDPConnection(out []*net.UDPConn) error {
	for i := range out {
		err := out[i].Close()
		if err != nil {
			return errors.WithMessage(err, "Could not close connectio")
		}
	}
	return nil
}
