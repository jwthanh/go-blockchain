package main

import (
	"time"
)

// Block store valueable information and technical data
type Block struct {
	Timestamp		int64
	Data			[]byte
	PrevBlockHash	[]byte
	Hash			[]byte
	Nonce 			int
}

// NewBlock generate a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}


// NewGenesisBlock return a Genesis Block, it's the first block in blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}