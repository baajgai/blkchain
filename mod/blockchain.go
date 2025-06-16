package mod

import (
	"crypto/sha256"

	Consts "github.com/baajgai/blkchain/consts"
)

// Blockchain struct contains blocks as a chain
type BlockChain struct {
	Chain          []Block
	UnspentOutputs HashSet
}

func (bc *BlockChain) UpdateWithBlock(b *Block) error {

	isize := len(bc.Chain)
	checksum := sha256.Sum256(b.Bytes())
	if b.Index != uint32(isize) {
		return Consts.MismatchedIndex

	} else if !CheckDifficulty(checksum[:], b.Difficulty) {

		return Consts.InvalidHash

	} else if isize != 0 {

		prev_block := bc.Chain[isize-1]
		prev_block_hash := sha256.Sum256(prev_block.Bytes())
		if b.TimeStamp <= prev_block.TimeStamp {
			return Consts.AchronologicalTimeStamp
		} else if b.PrevBlock != prev_block_hash {
			return Consts.MismatchedPreviousHash
		}

	} else {
		if !isNotAllZero(b.PrevBlock) {
			return Consts.InvalidGenesisBlockFormat
		}
	}

	if len(b.Transactions) > 0 {
		first_tx := b.Transactions[0]
		if first_tx.IsCoinbase() {
			return Consts.InvalidCoinBaseTransaction
		}

		// todo: set block spent hash here
		// block_spent := NewHashSet()
	}

	return nil

}

func isNotAllZero(b [32]byte) bool {
	for _, v := range b {
		if v != 0 {
			return true
		}
	}
	return false
}
