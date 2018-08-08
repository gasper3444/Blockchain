package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/hex"
)
//testing Add and IsValid
func TestBlockChainValidity(t *testing.T) {
	b0 := Initial(2)
	b0.Mine(1)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)

	var bChain Blockchain
	bChain.Add(b0)
	bChain.Add(b1)
	bChain.Add(b2)

//should be valid
	if !bChain.IsValid() {
		t.Error("Wrong valid blockchain function. (Valid chain is shown to be invalid)")
	}

	b2.Generation= 0
	var bChain2 Blockchain
	bChain2.Add(b0)
	bChain2.Add(b1)
	bChain2.Add(b2)

	//should be invalid
	if bChain2.IsValid() {
		t.Error("Wrong valid blockchain function. (Invalid chain is shown to be valid)")
	}
}
func TestNextFunc(t * testing.T) {
	b0 := Initial(2)
	b0.Mine(1)

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)

	cond1 := b0.Generation == b1.Generation -1
	cond2 := b0.Difficulty == b1.Difficulty
	cond3 := b1.Data == "this is an interesting message"

	mainCond := cond1 && cond2 && cond3
	assert.Equal(t, mainCond, true)
}
func TestCalcHash(t * testing.T) {
	b0 := Initial(2)
	b0.Mine(1)
	b1 := b0.Next("message")
	b1.Mine(1)

	c1 := hex.EncodeToString(b0.CalcHash())
	c2 := hex.EncodeToString(b1.CalcHash())

	assert.Equal(t, c1, "29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000")
	assert.Equal(t, c2, "02b09bde9ff60582ef21baa4bef87a95dfcd67efaf258e6df60463da0a940000" )
}
func TestValidHash(t * testing.T) {
	 b0 := Initial(2)
	 b0.Proof = 1111
	 b0.Hash= b0.CalcHash()

	 b1 := Initial(2)
	 b1.Proof = 242278
	 b1.Hash= b1.CalcHash()

	 assert.Equal(t, b0.ValidHash(), false)
	 assert.Equal(t, b1.ValidHash(), true)
}
