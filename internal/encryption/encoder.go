package encryption

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Encoder interface {
	EncodeHash(hashData HashData) (encoding string, err error)
	DecodeHash(encoding string) (hashData HashData, err error)
	EncodeEncryption(encryptionData EncryptionData) (encoding string, err error)
	DecodeEncryption(encoding string) (encryptionData EncryptionData, err error)
}

type encoder struct {
}

func NewEncoder() Encoder {
	return &encoder{}
}

func (e *encoder) EncodeHash(hashData HashData) (string, error) {
	bodyBase64 := base64.StdEncoding.EncodeToString(hashData.Body)
	saltBase64 := base64.StdEncoding.EncodeToString(hashData.Salt)

	encoding := strings.Builder{}
	encoding.WriteByte('$')
	count := 0
	for _, key := range []string{memoryStr, timeStr, threadsStr} {
		encoding.WriteString(fmt.Sprintf("%v=%v", key, hashData.Params[key]))
		if count < len(hashData.Params)-1 {
			encoding.WriteByte(',')
		}
		count++
	}

	encoding.WriteString(fmt.Sprintf("$%v$%v", saltBase64, bodyBase64))
	return encoding.String(), nil
}

func (e *encoder) DecodeHash(encoding string) (hashData HashData, err error) {
	splits := strings.Split(encoding, "$")
	if len(splits) != 4 {
		return HashData{}, fmt.Errorf("cannot decode hash: invalid encoding")
	}

	var memory int32
	var time int32
	var threads int32
	_, err = fmt.Sscanf(splits[1], "memory=%v,time=%v,threads=%v", &memory, &time, &threads)
	if err != nil {
		return HashData{}, fmt.Errorf("cannot decode hash: %v", err)
	}

	hashData.Salt, err = base64.StdEncoding.DecodeString(splits[2])
	if err != nil {
		return HashData{}, fmt.Errorf("cannot decode hash: %v", err)
	}
	hashData.Body, err = base64.StdEncoding.DecodeString(splits[3])
	if err != nil {
		return HashData{}, fmt.Errorf("cannot decode hash: %v", err)
	}

	hashData.Params = map[string]int32{
		memoryStr:  memory,
		timeStr:    time,
		threadsStr: threads,
	}

	return hashData, nil
}

func (e *encoder) EncodeEncryption(encryptionData EncryptionData) (encoding string, err error) {
	bodyBase64 := base64.StdEncoding.EncodeToString(encryptionData.Body)
	saltBase64 := base64.StdEncoding.EncodeToString(encryptionData.Salt)
	nonceBase64 := base64.StdEncoding.EncodeToString(encryptionData.Nonce)
	return fmt.Sprintf("$%v$%v$%v", nonceBase64, saltBase64, bodyBase64), nil
}

func (e *encoder) DecodeEncryption(encoding string) (encryptionData EncryptionData, err error) {
	splits := strings.Split(encoding, "$")
	if len(splits) != 4 {
		return EncryptionData{}, fmt.Errorf("cannot decode encryption: invalid encoding")
	}

	encryptionData.Nonce, err = base64.StdEncoding.DecodeString(splits[1])
	if err != nil {
		return EncryptionData{}, fmt.Errorf("cannot decode encryption: %v", err)
	}
	encryptionData.Salt, err = base64.StdEncoding.DecodeString(splits[2])
	if err != nil {
		return EncryptionData{}, fmt.Errorf("cannot decode encryption: %v", err)
	}
	encryptionData.Body, err = base64.StdEncoding.DecodeString(splits[3])
	if err != nil {
		return EncryptionData{}, fmt.Errorf("cannot decode encryption: %v", err)
	}

	return encryptionData, nil
}
