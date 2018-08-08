package blockchain

import (
	"encoding/hex"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Generation uint64
	Difficulty uint8
	Data       string
	PrevHash   []byte
	Hash       []byte
	Proof      uint64
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	var initialBlock Block
	initialBlock.Generation = 0
	initialBlock.Difficulty = difficulty
	initialBlock.Data = ""
	initialBlock.PrevHash = make([]byte, 32)			//an array of length 32 where each byte= 0
	return initialBlock
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO

	var nextBlock Block
	nextBlock.Generation= prev_block.Generation + 1
	nextBlock.Difficulty= prev_block.Difficulty
	nextBlock.Data= data
	nextBlock.PrevHash= prev_block.Hash
	return nextBlock
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	previousHash := hex.EncodeToString(blk.PrevHash)
	myGeneration := fmt.Sprintf("%v",blk.Generation)
	myDifficulty := fmt.Sprintf("%v",blk.Difficulty)
	myData := blk.Data
	myProof := fmt.Sprintf("%v",blk.Proof)
	myStringHash := (previousHash+":"+myGeneration+":"+myDifficulty+":"+myData+":"+myProof)

	actualHash := sha256.Sum256([]byte(myStringHash[:]))
	return (actualHash[:])
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	difficultyLevel := blk.Difficulty

	for difficultyLevel > 0 {
		if blk.Hash[len(blk.Hash)-int(difficultyLevel)] == 0 {
			difficultyLevel-=1
		} else {
			return false			//one of the last byte != 0
		}
	}
	return true		//all last difficulty bytes == 0
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
