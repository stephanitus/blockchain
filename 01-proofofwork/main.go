package main

import "fmt"

func main() {
	//Test blockchain code
	bc := NewBlockchain()

	//Add transactions
	bc.AddBlock("Receive 3 BTC from Larry")
	bc.AddBlock("Send 2 BTC to Jennifer")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
