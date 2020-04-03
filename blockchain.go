package main

// Blockchain - a simple chain
type Blockchain struct {
	// array of blocks
	blocks []*Block
}

// AddBlock - a fucntion to add blocks to the chain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain - function to create a new blockchain with genesis-block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
