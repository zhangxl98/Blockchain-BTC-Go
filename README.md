# 说明

> 创建时间：2018/12/05

> 开发环境：Ubuntu 18.04 LTS | go1.10.4

# 技术栈

Go | Blockchain

# 项目运行

## 环境

**go1.10及以上**

## 运行

```
git clone https://github.com/zhangxl98/Blockchain-BTC-Go.git
cd Blockchain-BTC-Go
go build *.go
./blockchain
```

## 结果

```
Begin Mining...
target hash :    010000000000000000000000000000000000000000000000000000000000
found hash : 000000369106e5a77440be8c3f0ed14a3f30ffd68140321e0396def2a198d8cf,nonce : 18617162
Begin Mining...
target hash :    010000000000000000000000000000000000000000000000000000000000
found hash : 0000006b43e1c088de9d0f5a9f0098167e9abc65552a44ab96617858c8b0c2bd,nonce : 8826594
Begin Mining...
target hash :    010000000000000000000000000000000000000000000000000000000000
found hash : 000000673ca701551735ac840889a4cdff66060525f0d9d93beb7d999cc77a69,nonce : 9604036
Version: 1
PrevBlochHash: 
CurrBlockHash: 000000369106e5a77440be8c3f0ed14a3f30ffd68140321e0396def2a198d8cf
TimeStamp: 1544016225
Bits: 1
Nonce: 18617162
Data: Gedesis Block!
Isvalid: true

Version: 1
PrevBlochHash: 000000369106e5a77440be8c3f0ed14a3f30ffd68140321e0396def2a198d8cf
CurrBlockHash: 0000006b43e1c088de9d0f5a9f0098167e9abc65552a44ab96617858c8b0c2bd
TimeStamp: 1544016245
Bits: 1
Nonce: 8826594
Data: A Send B 1 BTC
Isvalid: true

Version: 1
PrevBlochHash: 0000006b43e1c088de9d0f5a9f0098167e9abc65552a44ab96617858c8b0c2bd
CurrBlockHash: 000000673ca701551735ac840889a4cdff66060525f0d9d93beb7d999cc77a69
TimeStamp: 1544016256
Bits: 1
Nonce: 9604036
Data: B Send C 1 BTC
Isvalid: true
```



# 文件结构

```
Blockchain-BTC-Go/
├── blockchain.go
├── block.go
├── main.go
├── proofOfWork.go
├── README.md
└── utils.go

0 directories, 6 files
```



# 功能实现

- [x] 区块结构
- [x] 区块链结构
- [x] 生成新区块
- [x] 工作量证明
- [ ] 数据库存储
- [ ] 命令行操作

