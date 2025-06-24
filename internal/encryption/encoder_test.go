package encryption

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	encoderKeys = [][]byte{
		{},
		{1},
		{1, 2, 3},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{255, 254, 253, 200, 199, 100, 101, 102},
	}
	encoderInstances = [][]byte{
		{},
		{1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{24, 25, 26},
		{255, 254, 253, 200, 199, 100, 101, 102},
	}
)

func TestEncodeAndDecodeHash(t *testing.T) {
	//** Arrange
	encoder := NewEncoder()
	keyProvider := NewKeyProvider()
	hasher := NewHasher(keyProvider)

	//** Act and Assert
	for _, data := range encoderInstances {
		hashData, err := hasher.Hash(data)
		assert.Nil(t, err)
		encoding, err := encoder.EncodeHash(hashData)
		assert.Nil(t, err)
		decodedHashData, err := encoder.DecodeHash(encoding)
		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(hashData, decodedHashData))
	}
}

func TestEncodeAndDecodeEncryption(t *testing.T) {
	//** Arrange
	encoder := NewEncoder()
	keyProvider := NewKeyProvider()
	encryptor := NewEncryptor(keyProvider)

	//** Act and Assert
	for _, key := range encoderKeys {
		for _, data := range encoderInstances {
			encryptionData, err := encryptor.Encrypt(key, data)
			assert.Nil(t, err)
			encoding, err := encoder.EncodeEncryption(encryptionData)
			assert.Nil(t, err)
			decodedEncryptionData, err := encoder.DecodeEncryption(encoding)
			assert.Nil(t, err)
			assert.True(t, reflect.DeepEqual(encryptionData, decodedEncryptionData))
		}
	}
}
