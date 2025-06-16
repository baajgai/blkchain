package mod

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand/v2"
)

// Block structure
type Block struct {
	Index        uint32
	TimeStamp    uint64
	Hash         [32]byte
	PrevBlock    [32]byte
	Nonce        uint64
	Transactions []Transaction
	Difficulty   uint64
}

// Creates a new block constructor
func NewBlock(index uint32, timestamp uint64, prev_block [32]byte, transactions []Transaction, difficulty uint64) *Block {

	return &Block{
		Index:        index,
		TimeStamp:    timestamp,
		Hash:         [32]byte{0}, // Fill 32-length zeros
		PrevBlock:    prev_block,
		Transactions: transactions,
		Difficulty:   difficulty,
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
	bytes = append(bytes, b.PrevBlock[:]...)

	// Add Nonce as bytes
	nonce_bytes := U64Bytes(b.Nonce)
	bytes = append(bytes, nonce_bytes[:]...)

	// Add Difficulty as bytes
	diff_bytes := U64Bytes(b.Difficulty)
	bytes = append(bytes, diff_bytes[:]...)

	for _, tx := range b.Transactions {
		bytes = append(bytes, tx.Bytes()[:]...)
	}

	return bytes
}

// Mines bytes
func (b *Block) Mine() {
	rand_num := rand.Int32N(500)

	for i := 0; i < int(rand_num); i++ {

		b.Nonce = uint64(i)

		hash := sha256.Sum256(b.Bytes())

		fmt.Printf("HASH FROM MINE %v\n", hex.EncodeToString(hash[:]))
		if CheckDifficulty(hash[:], b.Difficulty) {
			b.Hash = hash
			return
		}

	}
}

// Checks difficulty level of hash
func CheckDifficulty(hash []byte, difficulty uint64) bool {
	return difficulty > DifficultyBytesAsU128(hash)
}
