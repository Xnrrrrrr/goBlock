package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// CreditCardTransaction represents a credit card transaction.
type CreditCardTransaction struct {
	CardNumber      string
	TransactionDate time.Time
	Amount          float64
	Merchant        string
}

// StudentGrade represents a student's grade.
type StudentGrade struct {
	StudentID  string
	CourseCode string
	Grade      int
}

// BankingTransaction represents a banking transaction.
type BankingTransaction struct {
	AccountNumber   string
	TransactionDate time.Time
	Amount          float64
	TransactionType string
}

// Block represents a block in the blockchain.
type Block struct {
	Index        int
	Timestamp    int64
	PreviousHash string
	Hash         string
	Data         string
	Difficulty   int
}

// Blockchain represents a blockchain.
type Blockchain struct {
	Chain []Block
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	fmt.Println("Creating a new blockchain...")
	genesisBlock := createGenesisBlock()
	return &Blockchain{Chain: []Block{genesisBlock}}
}

// createGenesisBlock creates the genesis block of the blockchain.
func createGenesisBlock() Block {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		PreviousHash: "0",
		Data:         "Genesis Block",
		Difficulty:   3,
	}

	// Calculate the hash for the genesis block
	genesisBlock.Hash = calculateHash(genesisBlock)

	return genesisBlock
}

// addBlock adds a new block to the blockchain.
func (bc *Blockchain) addBlock(data string) {
	previousBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := mineBlock(Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousBlock.Hash,
		Data:         data,
		Difficulty:   3,
	})
	bc.Chain = append(bc.Chain, newBlock)
}

// mineBlock mines a block with the specified difficulty.
func mineBlock(b Block) Block {
	target := "000"        // Difficulty set to 3, so the target starts with three zeros
	maxAttempts := 1000000 // Set a reasonable maximum number of attempts

	fmt.Println("Mining block...")

	for i := 0; i < maxAttempts; i++ {
		b.Timestamp = time.Now().Unix()
		b.Hash = calculateHash(b)

		if b.Hash[:3] == target {
			fmt.Println("Block mined successfully!")
			return b
		}

		// Print intermediate hash values for tracking progress
		if i%10000 == 0 {
			fmt.Printf("Hash attempt %d: %s\n", i, b.Hash)
		}
	}

	// Indicate that mining failed if the loop completes without finding a suitable hash
	fmt.Println("Mining failed. Maximum attempts reached.")
	return b
}

// calculateHash calculates the hash of a block.
func calculateHash(b Block) string {
	record := fmt.Sprintf("%d%d%s%s", b.Index, b.Timestamp, b.PreviousHash, b.Data)
	hashInBytes := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hashInBytes[:])
}

// printBlockchain prints the contents of the blockchain to the console.
func (bc *Blockchain) printBlockchain() {
	fmt.Println("Printing the blockchain...")
	for _, block := range bc.Chain {
		fmt.Printf("Block #%d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n\n", block.Data)
	}
	fmt.Println("Blockchain printed.")
}

func main() {
	// Create a new blockchain
	fmt.Println("Creating a new blockchain...")
	blockchain := NewBlockchain()

	// Add some blocks to the blockchain with POGO data
	fmt.Println("Adding blocks to the blockchain...")

	creditCardTransaction := CreditCardTransaction{
		CardNumber:      "1234-5678-9012-3456",
		TransactionDate: time.Now(),
		Amount:          50.0,
		Merchant:        "Online Store",
	}

	studentGrade := StudentGrade{
		StudentID:  "S123456",
		CourseCode: "Math",
		Grade:      90,
	}

	bankingTransaction := BankingTransaction{
		AccountNumber:   "123456789",
		TransactionDate: time.Now(),
		Amount:          1000.0,
		TransactionType: "Deposit",
	}

	// Adjust the difficulty level for testing
	newBlock := mineBlock(Block{
		Index:        1, // Assuming this is the second block
		Timestamp:    time.Now().Unix(),
		PreviousHash: blockchain.Chain[0].Hash, // Replace with the actual hash
		Data:         fmt.Sprintf("%+v", creditCardTransaction),
		Difficulty:   1, // Lower difficulty for testing
	})
	blockchain.Chain = append(blockchain.Chain, newBlock)

	newBlock = mineBlock(Block{
		Index:        2, // Assuming this is the third block
		Timestamp:    time.Now().Unix(),
		PreviousHash: blockchain.Chain[1].Hash, // Replace with the actual hash
		Data:         fmt.Sprintf("%+v", studentGrade),
		Difficulty:   1, // Lower difficulty for testing
	})
	blockchain.Chain = append(blockchain.Chain, newBlock)

	newBlock = mineBlock(Block{
		Index:        3, // Assuming this is the fourth block
		Timestamp:    time.Now().Unix(),
		PreviousHash: blockchain.Chain[2].Hash, // Replace with the actual hash
		Data:         fmt.Sprintf("%+v", bankingTransaction),
		Difficulty:   1, // Lower difficulty for testing
	})
	blockchain.Chain = append(blockchain.Chain, newBlock)

	// Print the blockchain
	blockchain.printBlockchain()
}
