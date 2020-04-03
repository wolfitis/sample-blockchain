package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block - a single block in blockchain
type Block struct {
	Timestamp     int64  // current timestap
	Data          []byte // actual information
	PrevBlockHash []byte // hash of previous block
	Hash          []byte // hash of block
}

// SetHash - temporary fucntion to calculate hashes
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock - simplify the creation of a block
func NewBlock(data string, PrevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock - a necessary first block aka genesis-block in blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genensis Block", []byte{})
}
