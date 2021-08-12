package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *block) getHash() string {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	return fmt.Sprintf("%x", hash)
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	newBlock.hash = newBlock.getHash()
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Prev: %s\n", block.prevHash)
		fmt.Printf("Hash: %s\n", block.hash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("genesis block1")
	chain.addBlock("genesis block2")
	chain.addBlock("genesis block3")
	chain.listBlocks()
}
