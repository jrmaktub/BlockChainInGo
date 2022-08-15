package blockchain

//derive the Hash from the block's data & previous Hash, and other meta data
type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

//flow 4
// func (b *Block) DeriveHash() {
// 	// This will join our previous block's relevant info with the new blocks
// 	//2Dimensional slice of bytes
// 	//ask
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	//This performs the actual hashing algorithm
// 	hash := sha256.Sum256(info)
// 	//placeholder to show how hash changes as the data changes
// 	b.Hash = hash[:]
// }

//flow 3
//takes in astring of Data,and the prevHash, and returns a pointer to a block
func CreateBlock(data string, prevHash []byte) *Block {
	//using the block constructor
	//hash fields empty slice of bytes
	//Data field taking the data string and  convert into a slice of bytes
	//prevHash is prevHash
	//0  for nonce
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	//run proof of work, and store nonce and hash
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	//flow 4

	block.Hash = hash[:]
	block.Nonce = nonce

	//flow 5
	return block
}

//takes in the pointer from our BlockChain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)

}

//return createBlock call with Data for Genesis block, and an empty previous Hash
//flow 2
func Genesis() *Block {
	//flow 3: create block call
	return CreateBlock("Genesis", []byte{})
}

//return a reference to the blockchain
//inside of it is  an array of Blocks
//with a call to the Genesis function
//flow:1
func InitBlockChain() *BlockChain {
	//flow 2: Genesis Function
	return &BlockChain{[]*Block{Genesis()}}
}
