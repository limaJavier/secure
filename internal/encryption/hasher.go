package encryption

import (
	"crypto/subtle"

	"golang.org/x/crypto/argon2"
)

type HashData struct {
	Body, Salt []byte
	Params     map[string]int32
}

type Hasher interface {
	Hash(data []byte) (hashData HashData, err error)
	Verify(data []byte, hashData HashData) (bool, error)
}

type hasherArgon2Id struct {
	provider KeyProvider
}

func NewHasher(provider KeyProvider) Hasher {
	return &hasherArgon2Id{
		provider: provider,
	}
}

func (h *hasherArgon2Id) Hash(data []byte) (hashData HashData, err error) {
	hashData.Salt, err = h.provider.GenerateRandomKey(randomSaltBytes)
	if err != nil {
		return HashData{}, err
	}

	hashData.Body = argon2.IDKey(data, hashData.Salt, argon2idTime, argon2idMemory, argon2idThreads, argon2idKeyLen)

	hashData.Params = map[string]int32{
		memoryStr:  int32(argon2idMemory),
		timeStr:    int32(argon2idTime),
		threadsStr: int32(argon2idThreads),
	}

	return hashData, nil
}

func (h *hasherArgon2Id) Verify(data []byte, hashData HashData) (bool, error) {
	// TODO: Validate params before usage

	// Extract params
	var memory uint32 = uint32(hashData.Params[memoryStr])
	var time uint32 = uint32(hashData.Params[timeStr])
	var threads uint8 = uint8(hashData.Params[threadsStr])
	var keyLen = uint32(len(hashData.Body))

	// Hash data
	hash := argon2.IDKey(data, hashData.Salt, time, memory, threads, keyLen)

	return subtle.ConstantTimeCompare(hashData.Body, hash) == 1, nil // Compare hashes in constant time
}
