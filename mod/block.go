package mod

import (
	"crypto/sha256"
	"math"
)

// Block structure
type Block struct {
	Index        uint32
	TimeStamp    uint64
	Hash         []byte
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
		Hash:         make([]byte, 32), // Fill 32-length zeros
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
	max64 := uint64(math.MaxUint64)

	for nonce_attempt := range max64 {
		b.Nonce = nonce_attempt

		hash := sha256.Sum256(b.Bytes())

		if CheckDifficulty(hash[:], b.Difficulty) {
			b.Hash = hash[:]
			return
		}
	}
}

// Checks difficulty level of hash
func CheckDifficulty(hash []byte, difficulty uint64) bool {
	return difficulty > DifficultyBytesAsU128(hash)
}
