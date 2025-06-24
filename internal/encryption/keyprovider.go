package encryption

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

type KeyProvider interface {
	GenerateRandomKey(bytes int) ([]byte, error)
	GenerateSaltedKey(key, salt []byte) ([]byte, error)
}

type keyProviderArgon2Id struct{}

func NewKeyProvider() KeyProvider {
	return &keyProviderArgon2Id{}
}

func (k *keyProviderArgon2Id) GenerateRandomKey(bytes int) ([]byte, error) {
	key := make([]byte, bytes)
	rand.Read(key)
	return key, nil
}

func (k *keyProviderArgon2Id) GenerateSaltedKey(key []byte, salt []byte) ([]byte, error) {
	return argon2.IDKey(key, salt, argon2idTime, argon2idMemory, argon2idThreads, argon2idKeyLen), nil
}
