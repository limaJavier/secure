package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

const (
	nonceBytes = 12
)

type EncryptionData struct {
	Body, Salt, Nonce []byte
}

type Encryptor interface {
	Encrypt(key, plainData []byte) (encryptionData EncryptionData, err error)
	Decrypt(key []byte, encryptionData EncryptionData) (plainData []byte, err error)
}

type encryptor struct {
	provider KeyProvider
}

func NewEncryptor(provider KeyProvider) Encryptor {
	return &encryptor{
		provider: provider,
	}
}

func (e *encryptor) Encrypt(key, plainData []byte) (encryptionData EncryptionData, err error) {
	encryptionData.Salt, err = e.provider.GenerateRandomKey(randomSaltBytes)
	if err != nil {
		return EncryptionData{}, err
	}

	saltedKey, err := e.provider.GenerateSaltedKey(key, encryptionData.Salt)
	if err != nil {
		return EncryptionData{}, err
	}

	block, err := aes.NewCipher(saltedKey)
	if err != nil {
		return EncryptionData{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return EncryptionData{}, err
	}

	encryptionData.Nonce, err = e.provider.GenerateRandomKey(aesGCM.NonceSize())
	if err != nil {
		return EncryptionData{}, err
	}

	// Assert nonce length
	if nonceBytes != len(encryptionData.Nonce) {
		log.Panicf("nonce length is not as expected: %v bytes", nonceBytes)
	}

	encryptionData.Body = aesGCM.Seal(nil, encryptionData.Nonce, plainData, nil)

	return encryptionData, nil
}

func (e *encryptor) Decrypt(key []byte, encryptionData EncryptionData) (plainData []byte, err error) {
	saltedKey, err := e.provider.GenerateSaltedKey(key, encryptionData.Salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(saltedKey)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainData, err = aesGCM.Open(nil, encryptionData.Nonce, encryptionData.Body, nil)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}
