package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64 //Time of block creation
	Transactions  []*Transaction
	PrevBlockHash []byte //Hash
	Hash          []byte //Hash
	Nonce         int    //Hash component
}

//Serialize a block for storage
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	//Initialize block struct and store address in block
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	//Generate proof of work
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	//Create and store the hash
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//Genesis block is the starting block of a blockchain
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
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
