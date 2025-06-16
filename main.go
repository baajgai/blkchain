package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/baajgai/blkchain/mod"
)

// Main function
func main() {

	log.Println("Starting main run...")

	var empty_prev_block [32]byte

	block := mod.NewBlock(0, uint64(time.Now().Unix()), empty_prev_block, []mod.Transaction{}, 100)
	log.Println("Start mining first genesis block...")
	block.Mine()

	fmt.Println("Mine genesis block : \n", block)

	last_hash := sha256.Sum256(block.Bytes())

	blockchain := mod.BlockChain{
		Chain: []mod.Block{*block},
	}

	for i := 0; i < 300; i++ {
		block := mod.NewBlock(uint32(i), uint64(time.Now().Unix()), last_hash, []mod.Transaction{}, 200)
		block.Mine()

		fmt.Printf("Mined new block: %+v\n", block)

		blockchain.Chain = append(blockchain.Chain, *block)

	}

	for _, block := range blockchain.Chain {
		fmt.Printf("block in block chain %+v\n", block.Hash)
	}

}
