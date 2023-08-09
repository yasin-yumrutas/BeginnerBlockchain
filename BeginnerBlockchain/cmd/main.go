package main

import (
	"fmt"

	"block/entities"
	"block/helper"
)

type Blockchain struct {
	Chain []entities.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := helper.CreateBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Chain = append(bc.Chain, newBlock)
}

func main() {
	genesisBlock := helper.CreateBlock(0, "Genesis Block", "")
	blockchain := Blockchain{Chain: []entities.Block{genesisBlock}}

	fmt.Println("Genesis Block Oluşturuldu:")
	fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n",
		genesisBlock.Index, genesisBlock.Timestamp, genesisBlock.Data, genesisBlock.PrevHash, genesisBlock.Hash)

	fmt.Println("\nBir sonraki bloğu eklemek için veri girin:")
	var data string
	fmt.Scanln(&data)
	blockchain.AddBlock(data)

	newBlock := blockchain.Chain[len(blockchain.Chain)-1]
	fmt.Printf("\nYeni Blok Oluşturuldu:\nIndex: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n",
		newBlock.Index, newBlock.Timestamp, newBlock.Data, newBlock.PrevHash, newBlock.Hash)
}
