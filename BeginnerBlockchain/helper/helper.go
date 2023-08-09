package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"block/entities"
)

func CreateBlock(index int, data string, prevHash string) entities.Block {
	block := entities.Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = CalculateHash(block)
	return block
}

func CalculateHash(block entities.Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
