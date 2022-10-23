package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/chalfel/complete-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	buf := &bytes.Buffer{}

	assert.Nil(t, h.EncodeBinary(buf))

	hDecoded := &Header{}

	assert.Nil(t, hDecoded.DecoreBinary(buf))
	assert.Equal(t, h, hDecoded)
}

func TestBlock_Encode_Decode(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	b := &Block{
		Header:       h,
		Transactions: nil,
	}

	buf := &bytes.Buffer{}

	assert.Nil(t, b.EncodeBinary(buf))

	blockDecoded := &Block{}

	assert.Nil(t, blockDecoded.DecoreBinary(buf))
	assert.Equal(t, b, blockDecoded)
}

func TestBlockHash(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	b := &Block{
		Header:       h,
		Transactions: nil,
	}

	hash := b.Hash()
	fmt.Println(hash)
	assert.False(t, hash.IsZero())
}
