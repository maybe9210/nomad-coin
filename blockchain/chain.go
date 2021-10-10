package blockchain

import (
	"fmt"
	"sync"

	"github.com/maybe9210/nomad-coin/db"
	"github.com/maybe9210/nomad-coin/utils"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecaculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp - lastRecaculatedBlock.Timestamp) / 60
	expectedTime := difficultyInterval * blockInterval
	if actualTime < expectedTime-allowedRange {
		return b.CurrentDifficulty + 1
	} else if actualTime > expectedTime+allowedRange {
		return b.CurrentDifficulty - 1
	} else {
		return b.CurrentDifficulty
	}
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{Height: 0}
			checkPoint := db.Checkpoint()
			if checkPoint == nil {
				b.AddBlock()
			} else {
				b.restore(checkPoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}
