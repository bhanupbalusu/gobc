package main

import (
	"fmt"
	"strconv"

	bc "github.com/bhanupbalusu/gobc/blockchain"
)

func main() {
	chain := bc.InitBlockchain()
	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash : %x\n", block.PrevHash)
		fmt.Printf("Data in the Block : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)

		pow := bc.NewProof(block)
		fmt.Printf("POW : %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}
