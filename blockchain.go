package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain.
type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

// Blockchain represents the entire blockchain.
type Blockchain struct {
	Blocks []Block
}

// DisplayAllBlocks prints all the blocks in the blockchain.
func DisplayAllBlocks(bc Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n\n", block.PrevHash)
	}
}

// NewBlock creates a new block and adds it to the blockchain.
func NewBlock(prevBlock Block, data string) Block {
	index := prevBlock.Index + 1
	timestamp := time.Now().String()
	hash := calculateHash(index, timestamp, data, prevBlock.Hash)

	return Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
		Hash:      hash,
		PrevHash:  prevBlock.Hash,
	}
}

// ModifyBlock modifies the data of a specific block in the blockchain.
func ModifyBlock(bc *Blockchain, index int, newData string) {
	if index >= 0 && index < len(bc.Blocks) {
		prevBlock := bc.Blocks[index-1]
		bc.Blocks[index] = NewBlock(prevBlock, newData)
	}
}

// calculateHash calculates the SHA-256 hash of the given parameters.
func calculateHash(index int, timestamp, data, prevHash string) string {
	hashInput := fmt.Sprintf("%d%s%s%s", index, timestamp, data, prevHash)
	hashBytes := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hashBytes[:])
}

func main() {
	// Create genesis block
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		Hash:      calculateHash(0, time.Now().String(), "Genesis Block", ""),
		PrevHash:  "",
	}

	// Create blockchain with genesis block
	blockchain := Blockchain{Blocks: []Block{genesisBlock}}

	// Display all blocks in the blockchain
	fmt.Println("Initial Blockchain:")
	DisplayAllBlocks(blockchain)

	// Add a new block
	newData := "New Data"
	newBlock := NewBlock(blockchain.Blocks[len(blockchain.Blocks)-1], newData)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

	// Display all blocks in the updated blockchain
	fmt.Println("\nBlockchain after adding a new block:")
	DisplayAllBlocks(blockchain)

	// Modify the data of the second block
	modifiedData := "Modified Data"
	ModifyBlock(&blockchain, 2, modifiedData)

	// Display all blocks in the blockchain after modification
	fmt.Println("\nBlockchain after modifying the second block:")
	DisplayAllBlocks(blockchain)
}