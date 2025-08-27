package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("No param KEY in environment params")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encryp(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGSM.NonceSize())
	return aesGSM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encrypter) Decrup(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSIze := aesGSM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSIze], encryptedStr[nonceSIze:]
	plainStr, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainStr
}
