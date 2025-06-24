package encryption

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	encryptorKeys = [][]byte{
		{},
		{1},
		{1, 2, 3},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{255, 254, 253, 200, 199, 100, 101, 102},
	}
	encryptorInstances = [][]byte{
		{},
		{1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{24, 25, 26},
		{255, 254, 253, 200, 199, 100, 101, 102},
	}
)

func TestEncryptAndDecrypt(t *testing.T) {
	//** Arrange
	keyProvider := NewKeyProvider()
	encryptor := NewEncryptor(keyProvider)

	//** Act and Assert
	for _, key := range encryptorKeys {
		for _, data := range encryptorInstances {
			encryptionData, err := encryptor.Encrypt(key, data)
			assert.Nil(t, err)
			decryptedData, err := encryptor.Decrypt(key, encryptionData)
			assert.Nil(t, err)
			assert.True(t, slices.Equal(data, decryptedData))
		}
	}
}
