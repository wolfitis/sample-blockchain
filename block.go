package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Block - a single block in blockchain
type Block struct {
	Timestamp     int64  // current timestap
	Data          []byte // actual information
	PrevBlockHash []byte // hash of previous block
	Hash          []byte // hash of block
	Nonce         int
}

// SetHash - temporary fucntion to calculate hashes
// func (b *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)

// 	b.Hash = hash[:]
// }

// NewBlock - simplify the creation of a block
func NewBlock(data string, PrevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock - a necessary first block aka genesis-block in blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genensis Block", []byte{})
}

// Serialize - serialize for storing block to BoltDB
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)

	return result.Bytes()
}

// DeserializeBlock -
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&block)

	return &block
}
