package mod

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/baajgai/blkchain/mod"
)

func TestBlock(t *testing.T) {
	var empty_prev_block [32]byte

	// set hash difficulty here...
	difficulty := 1

	block := NewBlock(0, uint64(time.Now().Unix()), empty_prev_block, []mod.Transaction{}, uint64(difficulty))

	block.Mine()

	last_hash := sha256.Sum256(block.Bytes())

	blockchain := mod.BlockChain{
		Chain: []mod.Block{*block},
	}

	for i := 1; i < 100; i++ {

		tx := mod.NewTransaction("addresstoBaajgai", "addressFromBaajgai", rand.Uint64()+uint64(i))
		tx.SignTransaction()
		block := mod.NewBlock(uint32(i), uint64(time.Now().Unix()), last_hash, []mod.Transaction{}, uint64(difficulty))

		block.Transactions = append(block.Transactions, *tx)
		block.Mine()

		copy(last_hash[:], block.Hash[:])
		blockchain.Chain = append(blockchain.Chain, *block)
	}

	fmt.Println("Total blocks : ", len(blockchain.Chain))
	for _, block := range blockchain.Chain {
		fmt.Printf("prev hash in block chain %s\n", hex.EncodeToString(block.PrevBlockHash[:]))
		fmt.Printf("block in block chain %s\n", hex.EncodeToString(block.Hash[:]))
	}
}
