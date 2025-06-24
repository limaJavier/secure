package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var hasherInstances = [][]byte{
	{},
	{1},
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	{24, 25, 26},
	{255, 254, 253, 200, 199, 100, 101, 102},
}

func TestHashAndVerify(t *testing.T) {
	//** Arrange
	keyProvider := NewKeyProvider()
	hasher := NewHasher(keyProvider)

	//** Act and Assert
	for _, data := range hasherInstances {
		hashData, err := hasher.Hash(data)
		assert.Nil(t, err)
		match, err := hasher.Verify(data, hashData)
		assert.Nil(t, err)
		assert.True(t, match)
	}
}
