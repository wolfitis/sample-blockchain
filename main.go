package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to user1")
	bc.AddBlock("Send 3 more BTC to user1")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
