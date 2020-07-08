package udp

import "crypto"

type shard struct {
	data []byte
}

type shards struct {
	parity map[crypto.Hash]shard
	data   map[crypto.Hash]shard
}

type header struct {
	fileSize int64
	fileName string
	shards   shards
}
