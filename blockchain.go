package main

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain {
	block := NewGenesisBlock()
	return &Blockchain{blocks: []*Block{block}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlockHash := bc.blocks[len(bc.blocks)-1].Hash
	block:=NewBlock(data, prevBlockHash)
	bc.blocks=append(bc.blocks,block)
}
