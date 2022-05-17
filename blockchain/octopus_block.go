package blockchain

import (
	"math/big"
	"sync/atomic"
)

type BlockNonce [8]byte

//区块头结构体
type Header struct {
	ParentHash Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.parentHash"` //父hash
	UncleHash  Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.uncleHash"`  //叔hash
	//Coinbase    common.Address
	Root        Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.root"`        //根hash
	TxHash      Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.txhash"`      //交易hash
	ReceiptHash Hash `autoInjectCfg:"octopus.blockchain.binding.genesis.header.receiptHash"` //收据hash
	//Bloom       Bloom
	Difficulty *big.Int `autoInjectCfg:"octopus.blockchain.binding.genesis.header.difficulty"` //难度值
	Number     *big.Int `autoInjectCfg:"octopus.blockchain.binding.genesis.header.number"`     //数量
	GasLimit   uint64   `autoInjectCfg:"octopus.blockchain.binding.genesis.header.gasLimit"`   //gas限制
	GasUsed    uint64   `autoInjectCfg:"octopus.blockchain.binding.genesis.header.gasUsed"`    //gas总和
	Time       uint64   `autoInjectCfg:"octopus.blockchain.binding.genesis.header.time"`       //时间戳
	//Extra       []byte
	//MixDigest   common.Hash
	Nonce BlockNonce `autoInjectCfg:"octopus.blockchain.binding.genesis.header.nonce"` //唯一标识
}

//数据容器
type Body struct {
	//Transactions []*Transaction
	Uncles []*Header
}

type Block struct {
	header *Header   //区块头信息
	uncles []*Header //叔块头信息
	//transactions 	Transactions //交易信息
	// caches
	hash atomic.Value //缓存hash
	size atomic.Value //缓存大小
	td   *big.Int     //交易总难度
}

func (b Block) newGenesis() {

}
