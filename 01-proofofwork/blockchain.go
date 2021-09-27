package main

type Blockchain struct {
	blocks []*Block
}

//AddBlock can be called on Blockchain pointers
func (bc *Blockchain) AddBlock(data string) {
	//Grab address of last block in blockchain
	prevBlock := bc.blocks[len(bc.blocks)-1]
	//Make a new block and store pointer
	newBlock := NewBlock(data, prevBlock.Hash)
	//Add block to blockchain
	bc.blocks = append(bc.blocks, newBlock)
}

//Initialize a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
