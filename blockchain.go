package main

import (
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain - a simple chain
type Blockchain struct {
	tip []byte   // tip of the chain
	db  *bolt.DB // DB connection
}

// BlockchainIterator - to iterate over blocks
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// AddBlock - a fucntion to add blocks to the chain
func (bc *Blockchain) AddBlock(data string) {
	// blocks will be stored in the DB
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}

// NewBlockchain - function to create a new blockchain with genesis-block
func NewBlockchain() *Blockchain {
	var tip []byte
	// standard way of opening BoltDB file
	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			// if bucket doesn't exist
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash) // updating last block hash
			tip = genesis.Hash
		} else {
			// if bucket exist we read hash of last block as 'l'
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

// Iterator - function to iterate over blocks
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Next - will return next block
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	i.currentHash = block.PrevBlockHash

	return block
}
