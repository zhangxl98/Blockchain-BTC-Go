package main

import (
	"github.com/boltdb/bolt"
	"os"
)

const dbFile = "blockchain.db"
const blockBucket = "bucket"
const lastHashKey = "key"

type Blockchain struct {
	// 数据库操作句柄
	db *bolt.DB
	// 尾巴，表示最后一个区块的哈希值
	tail []byte
}

func NewBlockchain() *Blockchain {
	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {}
	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("NewBlockchain_1", err)

	var lastHash []byte

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			// 取出最后区块的哈希值
			lastHash = bucket.Get([]byte(lastHashKey))
		} else {
			// 没有 bucket，要去创建创世块，将数据填写到数据库的 bucket 中
			genesis := NewGenesisBlock()
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			CheckErr("NewBlockchain_2", err)
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			CheckErr("NewBlockchain_3", err)
			err = bucket.Put([]byte(lastHashKey), genesis.Hash)
			CheckErr("NewBlockchain_4", err)
			lastHash = genesis.Hash
		}

		return nil
	})
	return &Blockchain{db, lastHash}
}

func (bc *Blockchain) AddBlock(data string) {
	var prevBlockHash []byte
	bc.db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket ==nil{
			os.Exit(1)
		}

		prevBlockHash=bucket.Get([]byte(lastHashKey))
		return nil
	})

	block:=NewBlock(data,prevBlockHash)

	err:=bc.db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket ==nil{
			os.Exit(1)
		}

		err := bucket.Put(block.Hash, block.Serialize())
		CheckErr("AddBlock_1", err)
		err = bucket.Put([]byte(lastHashKey), block.Hash)
		CheckErr("AddBlock_2", err)
		bc.tail=block.Hash
		return nil
	})
	CheckErr("AddBlock_3",err)
}
