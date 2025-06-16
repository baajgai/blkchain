package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/baajgai/blkchain/mod"
)

// Main function
func main() {

	log.Println("Starting main run...")

	var empty_prev_block [32]byte

	difficulty := 2

	block := mod.NewBlock(0, uint64(time.Now().Unix()), empty_prev_block, []mod.Transaction{}, uint64(difficulty))
	log.Println("Start mining first genesis block...")
	block.Mine()

	last_hash := sha256.Sum256(block.Bytes())

	blockchain := mod.BlockChain{
		Chain: []mod.Block{*block},
	}

	for i := 1; i < 100; i++ {
		block := mod.NewBlock(uint32(i), uint64(time.Now().Unix()), last_hash, []mod.Transaction{}, uint64(difficulty))
		block.Mine()

		fmt.Printf("Prev block hash === %s\n", hex.EncodeToString(last_hash[:]))
		fmt.Printf("Own block hash === %s\n", hex.EncodeToString(block.Hash[:]))

		copy(last_hash[:], block.Hash[:])
		blockchain.Chain = append(blockchain.Chain, *block)

		fmt.Println("INDEXXX : ", block.Index)
	}

	fmt.Println("Total blocks : ", len(blockchain.Chain))
	for _, block := range blockchain.Chain {
		fmt.Printf("prev hash in block chain %s\n", hex.EncodeToString(block.PrevBlock[:]))
		fmt.Printf("block in block chain %s\n", hex.EncodeToString(block.Hash[:]))
	}

}
