package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {

	bc, err := NewBlockchain(randomBlock(0))

	assert.NotNil(t, bc)
	assert.Nil(t, err)
	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlock := 1_000
	for i := 0; i < lenBlock; i++ {
		block := randomBlock(uint32(i + 1))
		assert.Nil(t, bc.AddBlock(block))
	}
	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.Equal(t, len(bc.headers), lenBlock+1)
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.True(t, bc.HasBlock(uint32(0)))
}
