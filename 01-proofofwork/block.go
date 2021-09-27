package main

import (
	"time"
)

type Block struct {
	Timestamp     int64  //Time of block creation
	Data          []byte //Transaction data
	PrevBlockHash []byte //Hash
	Hash          []byte //Hash
	Nonce         int    //Hash component
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	//Initialize block struct and store address in block
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	//Generate proof of work
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	//Create and store the hash
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//Genesis block is the starting block of a blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

/*
Deprecated: Hash function is now provided
			during proof of work generation

//Allows SetHash() to be called on Block typed objects
func (b *Block) SetHash() {
	//Convert timestamp to integer, store in byte array
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	//Concatenate block headers into slice
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	//Create checksum hash of headers
	hash := sha256.Sum256(headers)
	//Store hash in block
	b.Hash = hash[:]
}
*/
