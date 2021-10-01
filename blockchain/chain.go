package blockchain

import (
	"fmt"
	"sync"

	"github.com/maybe9210/nomad-coin/db"
	"github.com/maybe9210/nomad-coin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			checkPoint := db.Checkpoint()
			if checkPoint == nil {
				b.AddBlock("Genesis")
			} else {
				b.restore(checkPoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}
