package chain

import (
	"github.com/skycoin/encoder"
	//"time"
	"errors"
)

type BlockHeader struct {
	Time     uint64
	BkSeq    uint64 //increment every block
	PrevHash SHA256 //hash of header of previous block
	BodyHash SHA256 //hash of block body
}

func (self *BlockHeader) Hash() SHA256 {
	b1 := encoder.Serialize(*self)
	return SumDoubleSHA256(b1)
}

func (self *BlockHeader) Bytes() []byte {
	return encoder.Serialize(*self)
}

type Block struct {
	Head BlockHeader
	Body []byte //data here
}

/*
	Blockchain
*/

type BlockChain struct {
	Blocks []Block
	//Head   *Block
}

func NewBlockChain(phash SHA256) *BlockChain {
	var bc BlockChain

	var b Block

	b.Head.Time = 0
	b.Head.BkSeq = 0
	b.Head.PrevHash = phash

	bc.Blocks = append(bc.Blocks, b)

	return &bc
}

//returns head block
func (bc *BlockChain) Head() *Block {
	return &bc.Blocks[len(bc.Blocks)-1]
}

//creates new block
func (bc *BlockChain) NewBlock(blockTime uint64, data []byte) Block {
	var b Block
	b.Head.Time = blockTime
	b.BkSeq = bc.Head().Head.BkSeq + 1
	b.PrevHash = bc.Head().Head.Hash()
	b.BodyHash = SumSHA256(data)
	b.Body = data
	return b
}

func (bc *BlockChain) ApplyBlock(block Block) error {
	//do time check
	//do prevhash check
	//check body hash
	//check BkSeq

	bc.Blocks = append(bc.Blocks, block)
	return nil
}