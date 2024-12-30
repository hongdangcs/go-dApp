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
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block, Data string) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}

func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func main() {
	genesisBlock := Block{0, time.Now().String(), "Genesis Block", "", ""}
	bobBlock := Block{1, time.Now().String(), "Bob Block", "", ""}
	aliceBlock := Block{2, time.Now().String(), "Alice Block", "", ""}

	genesisBlock.Hash = CalculateHash(genesisBlock)
	genesisBlock.Hash = CalculateHash(genesisBlock)
	aliceBlock.Hash = CalculateHash(aliceBlock)

	Blockchain = append(Blockchain, genesisBlock)
	Blockchain = append(Blockchain, bobBlock)
	Blockchain = append(Blockchain, aliceBlock)

	newBlock := GenerateBlock(genesisBlock, "First Block After Genesis")
	if IsBlockValid(newBlock, genesisBlock) {
		Blockchain = append(Blockchain, newBlock)
	}

	for _, block := range Blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Println()
	}
}
