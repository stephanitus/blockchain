package main

import "fmt"

func main() {
	//Read in blockchain
	bc := NewBlockchain()
	defer bc.db.Close()

	//Create iterator
	bci := bc.Iterator()

	currentBlock := bci.Next()

	for bci.currentHash != nil {
		fmt.Printf("Prev. hash: %x\n", currentBlock.PrevBlockHash)
		fmt.Printf("Data: %s\n", currentBlock.Data)
		fmt.Printf("Hash: %x\n", currentBlock.Hash)
		currentBlock = bci.Next()
	}

	fmt.Printf("Prev. hash: %x\n", currentBlock.PrevBlockHash)
	fmt.Printf("Data: %s\n", currentBlock.Data)
	fmt.Printf("Hash: %x\n", currentBlock.Hash)
}
