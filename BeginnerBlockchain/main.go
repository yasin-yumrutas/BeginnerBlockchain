package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

type Blockchain struct {
	Chain []Block
}

func createBlock(index int, data string, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = calculateHash(block)
	return block
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (bc *Blockchain) addBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := createBlock(prevBlock.Index+1, data, prevBlock.Hash)
	bc.Chain = append(bc.Chain, newBlock)
}

func main() {
	genesisBlock := createBlock(0, "Genesis Block", "")
	blockchain := Blockchain{Chain: []Block{genesisBlock}}

	fmt.Println("Genesis Block Oluşturuldu:")
	fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n",
		genesisBlock.Index, genesisBlock.Timestamp, genesisBlock.Data, genesisBlock.PrevHash, genesisBlock.Hash)

	fmt.Println("\nBir sonraki bloğu eklemek için veri girin:")
	var data string
	fmt.Scanln(&data)
	blockchain.addBlock(data)

	newBlock := blockchain.Chain[len(blockchain.Chain)-1]
	fmt.Printf("\nYeni Blok Oluşturuldu:\nIndex: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\n",
		newBlock.Index, newBlock.Timestamp, newBlock.Data, newBlock.PrevHash, newBlock.Hash)
}
