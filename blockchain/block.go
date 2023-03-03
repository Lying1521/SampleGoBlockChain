package blockchain

import (
	"MyBlockChain/utils"
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	Nonce     int64
	Target    []byte
}

func (b *Block) SetHash() {
	info := bytes.Join([][]byte{utils.ToHexInt(b.Timestamp), b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(prevHash, data []byte) *Block {
	block := Block{
		Timestamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
	}
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	genesisWords := "Hello, blockChain"
	return CreateBlock([]byte{}, []byte(genesisWords))
}
