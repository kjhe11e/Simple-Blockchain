package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Welcome ---\n")

	bc := NewBlockchain()

	bc.AddBlock("Send 1 coin to Kyle")
	bc.AddBlock("Send 2 coins to Mariners")

	for _, block := range bc.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	fmt.Println("--- End of blockchain ---\n")
}

