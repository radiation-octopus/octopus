package blockchain

import (
	"octopus/consensus"
)

//处理器结构体
type BlockProcessor struct {
	//config *params.ChainConfig // 链配置
	bc     *BlockChain      // 标准链
	engine consensus.Engine // 共识引擎
}

//构建处理器
func NewBlockProcessor(bc *BlockChain, engine consensus.Engine) *BlockProcessor {
	bp := &BlockProcessor{
		bc:     bc,
		engine: engine,
	}
	return bp
}

//处理器接口
type Processor interface {
	//处理改变区块状态，将区块加入主链

}
