package core

import (
	"testing"
	"time"

	"github.com/chalfel/complete-blockchain/crypto"
	"github.com/chalfel/complete-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}
func TestSignBlock(t *testing.T) {

	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
	assert.Equal(t, b.Validator, privKey.PublicKey())

	invalidPrivKey := crypto.GeneratePrivateKey()

	b.Validator = invalidPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())
}

func TestVerifyBlock(t *testing.T) {

	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())
	assert.NotNil(t, b.Signature)
	assert.Equal(t, b.Validator, privKey.PublicKey())

	invalidPrivKey := crypto.GeneratePrivateKey()

	b.Validator = invalidPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
