package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	// 版本
	Version int64
	// 前区块的哈希值
	PrevBlockHash []byte
	// 当前区块的哈希值
	Hash []byte
	// 梅克尔根
	MerkelRoot []byte
	// 时间戳
	TimeStamp int64
	// 难度值
	Bits int64
	// 随机数
	Nonce int64

	// 交易信息
	Data []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       3,
		PrevBlockHash: prevBlockHash,
		//Hash:	TODO
		MerkelRoot: []byte{},
		TimeStamp:  time.Now().Unix(),
		Bits:       1,
		Nonce:      1,
		Data:       []byte(data)}

	//block.SetHash()
	pow := NewProofWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	CheckErr("Serialize", err)
	return buffer.Bytes()
}

func Deserialize(data []byte) *Block {
	if len(data) == 0 {
		return nil
	}
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	CheckErr("Deserialize", err)
	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("Gedesis Block!", []byte{})
}
