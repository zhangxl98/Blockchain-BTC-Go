package main

import (
	"bytes"
	"crypto/sha256"
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
		Version:       1,
		PrevBlockHash: prevBlockHash,
		//Hash:	TODO
		MerkelRoot: []byte{},
		TimeStamp:  time.Now().Unix(),
		Bits:       1,
		Nonce:      1,
		Data:       []byte(data)}
	block.SetHash()
	return &block
}

func (block *Block) SetHash() {

	// Sum256 returns the SHA256 checksum of the data.
	//func Sum256(data []byte) [Size]byte {}

	// Join concatenates the elements of s to create a new byte slice. The separator
	// sep is placed between elements in the resulting slice.
	//func Join(s [][]byte, sep []byte) []byte {}

	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerkelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Gedesis Block!", []byte{})
}
