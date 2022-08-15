package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//We're shifting the first arg target, however many units left we set our 2nd arg. The closer we get to 256, the easier the computation will be. Increasing our difficulty will increase the runtime of our algorithm.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//256 number of bytes in one Hash
	//LSH means Left-Shifted since substraction
	target.Lsh(target, uint(256-Difficulty))
	//put the target into an instance of ProofOfWork, as well as add  a block to it
	pow := &ProofOfWork{b, target}

	return pow
}

//will be on the ProofOfWork Struct
//as the parameter, it will take the nonce integer
//return a slyce of bytes

func (pow *ProofOfWork) InitNonce(nonce int) []byte {
	data := bytes.Join(
		//like the derivedHash function, will grab the blocks previous Hash and Block Data

		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		//combine it  with another byte struct using the join function
		//to create a cohesive set of bytes
		[]byte{},
	)
	return data
}

//turning int into Hex
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

//run on the ProofOfWork, will return an integer, and a slice  of bytes inside of a tuple
//it will take in no parameters
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	// This is essentially an infinite loop due to how large
	// MaxInt64 is.
	for nonce < math.MaxInt64 {
		//prepare Data, and then compare with our target Big Integer inside our PoW Struct
		data := pow.InitNonce(nonce)
		//hash Data into Sha256 format
		hash = sha256.Sum256(data)
		//hash is changing until it  find apropriate hash
		fmt.Printf("\r%x", hash)
		//convert hash into big Int,
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]

}

//after Run function,  we'll have the nonce
//use the  nonce to derive hash
func (pow *ProofOfWork) Validate() bool {
	//like  with Run  func,   we set up a  big int of our  Hash
	var intHash big.Int

	data := pow.InitNonce(pow.Block.Nonce)
	//turn the data into a Hash
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
