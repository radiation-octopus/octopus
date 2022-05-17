package blockchain

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"octopus/consensus"
	"octopus/db"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxFutureBlocks = 256
)

//blockchain结构体
type BlockChain struct {
	db            db.Database //数据库
	txLookupLimit uint64      `autoInjectCfg:"octopus.block.binding.txLookupLimit"` //一个区块容纳最大交易限制
	blockProcFeed Feed        //区块过程注入事件
	genesisBlock  *Block

	//chainmu *syncx.ClosableMutex	//互斥锁，同步链写入操作使用
	currentBlock atomic.Value // 当前区块
	//currentFastBlock atomic.Value	//快速同步链的当前区块

	futureBlocks  *lru.Cache     //新区块缓存区
	wg            sync.WaitGroup //同步等待属性
	quit          chan struct{}  //关闭属性
	running       int32          //运行属性
	procInterrupt int32          //块处理中断信号

	engine    consensus.Engine //共识引擎
	validator Validator        //区块验证器
	processor Processor        //区块处理器
	forker    *ForkChoice      //最高难度值
}

//链启动类，配置参数启动
func (bc *BlockChain) start() {

	newBlockChain(db.Database{}, nil, nil)
}

//链终止
func (bc *BlockChain) close() {

}

//构建区块链结构体
func newBlockChain(db db.Database, engine consensus.Engine, shouldPreserve func(header *Header) bool) (*BlockChain, error) {
	futureBlocks, _ := lru.New(maxFutureBlocks)
	bc := &BlockChain{
		db:   db,
		quit: make(chan struct{}),
		//chainmu:       syncx.NewClosableMutex(),
		futureBlocks: futureBlocks,
		engine:       engine,
	}
	//bc.forker = NewForkChoice(bc, shouldPreserve)
	//构建区块验证器
	bc.validator = NewBlockValidator(bc, engine)
	//构建区块处理器
	bc.processor = NewBlockProcessor(bc, engine)
	//获取创世区块
	bc.genesisBlock = bc.getBlockByNumber(0)
	//if bc.genesisBlock == nil {
	//	return nil, errors.New("创世区块未发现")
	//}
	if bc.empty() {
	}
	if err := bc.loadLastState(); err != nil {
		return nil, err
	}
	//确保区块有效可用一系列校验
	//head := bc.CurrentBlock()

	//开启未来区块处理
	bc.wg.Add(1)
	go bc.updateFutureBlocks()

	return bc, nil
}

//该创世区块在数据库对应是否有数据
func (bc *BlockChain) empty() bool {

	return true
}

//加载数据库最新链的状态
//同步区块数据
func (bc *BlockChain) loadLastState() error {

	return nil
}

//循环更新区块
func (bc *BlockChain) updateFutureBlocks() {
	futureTimer := time.NewTicker(5 * time.Second)
	defer futureTimer.Stop()
	defer bc.wg.Done()
	for {
		select {
		case <-futureTimer.C:
			bc.procFutureBlocks()
		case <-bc.quit:
			return
		}
	}
}

func (bc *BlockChain) procFutureBlocks() {
	fmt.Println("新增区块：")
	blocks := make([]*Block, 0, bc.futureBlocks.Len())
	for _, hash := range bc.futureBlocks.Keys() {
		if block, exist := bc.futureBlocks.Peek(hash); exist {
			blocks = append(blocks, block.(*Block))
		}
	}
	if len(blocks) > 0 {
		for i := range blocks {
			bc.InsertChain(blocks[i : i+1])
		}
	}
}

type Blocks []*Block

func (bc *BlockChain) InsertChain(chain Blocks) (int, error) {
	for i := 1; i < len(chain); i++ {
		block, prev := chain[i], chain[i-1]
		fmt.Println("新增区块：", block.hash, prev.hash)
	}
	return 1, nil
}
