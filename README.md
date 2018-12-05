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
Version: 1
PrevBlochHash: 
CurrHash: 74eb6fb9ce8405ebd438a20c8d582ddc4ffae831777001d7ccc5c95906458b38
TimeStamp: 1543994236
Bits: 1
Nonce: 1
Data: Gedesis Block!

Version: 1
PrevBlochHash: 74eb6fb9ce8405ebd438a20c8d582ddc4ffae831777001d7ccc5c95906458b38
CurrHash: ea267ca23ef615b947e1d2346fcda4cca5967324f25bf4680a38d28614386f8b
TimeStamp: 1543994236
Bits: 1
Nonce: 1
Data: A Send B 1 BTC

Version: 1
PrevBlochHash: ea267ca23ef615b947e1d2346fcda4cca5967324f25bf4680a38d28614386f8b
CurrHash: 431b053ca9500f0dfa40326ec4d563ddf12b0144ffae892ce06d5365023e3d8c
TimeStamp: 1543994236
Bits: 1
Nonce: 1
Data: B Send C 1 BTC
```



# 文件结构

```
Blockchain-BTC-Go/
├── blockchain.go
├── block.go
├── main.go
├── README.md
└── utils.go

0 directories, 5 files
```



# 功能实现

- [x] 区块结构
- [x] 区块链结构
- [x] 生成新区块
- [ ] 工作量证明
- [ ] 数据库存储
- [ ] 命令行操作

