package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	//Time of block creation
	Timestamp int64
	//Transaction data
	Data []byte
	//Hashes
	PrevBlockHash []byte
	Hash          []byte
}

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

func NewBlock(data string, prevBlockHash []byte) *Block {
	//Initialize block struct and store address in block
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	//Create and store the hash
	block.SetHash()
	return block
}

//Genesis block is the starting block of a blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
