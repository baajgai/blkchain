package mod

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Block structure
type Block struct {
	Index         uint32        // Index of block
	TimeStamp     uint64        // Block create time
	Hash          [32]byte      // Block's hash
	PrevBlockHash [32]byte      // Previous block's hash
	Nonce         uint64        // Nonce
	Transactions  []Transaction // Block's transaction list
	MerkleRoot    [32]byte      // Merkle root
	Difficulty    uint64        // Hash difficulty
}

// Creates a new block constructor
func NewBlock(index uint32, timestamp uint64, prev_block [32]byte, transactions []Transaction, difficulty uint64) *Block {

	return &Block{
		Index:         index,
		TimeStamp:     timestamp,
		Hash:          [32]byte{0}, // Fill 32-length zeros
		PrevBlockHash: prev_block,
		Transactions:  transactions,
		Difficulty:    difficulty,
	}
}

// Hashes interface functions
func (b Block) Bytes() []byte {
	var bytes []byte

	// Add Index as bytes
	index_bytes := U32Bytes(b.Index)
	bytes = append(bytes, index_bytes[:]...)

	// Add timestamp as bytes
	tstamp_bytes := U64Bytes(b.TimeStamp)
	bytes = append(bytes, tstamp_bytes[:]...)

	// Add Previous block as bytes
	bytes = append(bytes, b.PrevBlockHash[:]...)

	// Add Merkle root as bytes
	bytes = append(bytes, b.MerkleRoot[:]...)

	// Add Nonce as bytes
	nonce_bytes := U64Bytes(b.Nonce)
	bytes = append(bytes, nonce_bytes[:]...)

	// Add Difficulty as bytes
	diff_bytes := U64Bytes(b.Difficulty)
	bytes = append(bytes, diff_bytes[:]...)

	for _, tx := range b.Transactions {
		hash := tx.Hash()
		bytes = append(bytes, hash[:]...)
	}

	return bytes
}

// Mines bytes
func (b *Block) Mine() {

	targetPrefix := strings.Repeat("6", int(b.Difficulty))

	fmt.Println("Target prefix and difficulty : ", targetPrefix, b.Difficulty)
	nonce := 0
	for {

		b.Nonce = uint64(nonce)

		hash := sha256.Sum256([]byte(b.Bytes()))
		hashStr := hex.EncodeToString(hash[:])

		fmt.Println("Hash generating ... ", hashStr)
		if strings.HasPrefix(hashStr, targetPrefix) {
			b.Hash = hash
			fmt.Println("Hash Found: ", hex.EncodeToString(b.Hash[:]))
			return
		}

		nonce++

	}

}
