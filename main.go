package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//derive the Hash from the block's data & previous Hash, and other meta data
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

//flow 4
func (b *Block) DeriveHash() {
	// This will join our previous block's relevant info with the new blocks
	//2Dimensional slice of bytes
	//ask
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	//This performs the actual hashing algorithm
	hash := sha256.Sum256(info)
	//placeholder to show how hash changes as the data changes
	b.Hash = hash[:]
}

//flow 3
//takes in astring of Data,and the prevHash, and returns a pointer to a block
func createBlock(data string, prevHash []byte) *Block {
	//using the block constructor
	//hash fields empty slice of bytes
	//Data field taking the data string and  convert into a slice of bytes
	//prevHash is prevHash
	block := &Block{[]byte{}, []byte(data), prevHash}

	//flow 4
	block.DeriveHash()

	//flow 5
	return block
}

//takes in the pointer from our BlockChain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}

//return createBlock call with Data for Genesis block, and an empty previous Hash
//flow 2
func Genesis() *Block {
	//flow 3: create block call
	return createBlock("Genesis", []byte{})
}

//return a reference to the blockchain
//inside of it is  an array of Blocks
//with a call to the Genesis function
//flow:1
func InitBlockChain() *BlockChain {
	//flow 2: Genesis Function
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	//flow: 1
	chain := InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}

}
