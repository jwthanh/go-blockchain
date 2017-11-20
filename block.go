package main

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
)

// Block store valueable information and technical data
type Block struct {
	Timestamp		int64
	Data			[]byte
	PrevBlockHash	[]byte
	Hash			[]byte
}

// SetHash generate new hash. Simply, concatenate block fields and calculate a SHA-256 hash on it.
func (b *Block) SetHash()  {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock generate a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}


// NewGenesisBlock return a Genesis Block, it's the first block in blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}