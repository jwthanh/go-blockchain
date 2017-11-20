package main

// Blockchain contains blocks array
type Blockchain struct {
	blocks []*Block
}

// AddBlock appends a new block to blockchain
func (bc *Blockchain) AddBlock(data string)  {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain create a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}