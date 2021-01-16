package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func Encrypt(plaintext string) (string, error) {
	key, ok := os.LookupEnv("AES_KEY")
	if !ok {
		return "", errors.New("private key not set")
	}
	keyB, _ := hex.DecodeString(key)
	plaintextB := []byte(plaintext)

	block, err := aes.NewCipher(keyB)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipertext := aesGCM.Seal(nonce, nonce, plaintextB, nil)

	return fmt.Sprintf("%x", cipertext), nil
}

func Decrypt(encryptedStr string) (string, error) {
	key, ok := os.LookupEnv("AES_KEY")
	if !ok {
		return "", errors.New("private key not set")
	}
	keyB, _ := hex.DecodeString(key)
	cipertextB, _ := hex.DecodeString(encryptedStr)

	block, err := aes.NewCipher(keyB)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := cipertextB[:nonceSize], cipertextB[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)

	return fmt.Sprintf("%s", plaintext), nil
}
