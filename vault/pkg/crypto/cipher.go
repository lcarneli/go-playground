package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	nonceSize = 12
)

var (
	ErrInvalidCipherHex = fmt.Errorf("invalid cipher hex")
)

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := sha256.New()
	_, err := fmt.Fprintf(hasher, key)
	if err != nil {
		return nil, err
	}

	cipherKey := hasher.Sum(nil)

	return aes.NewCipher(cipherKey)
}

func Encrypt(key, plaintext string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceSize)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	combined := append(nonce, ciphertext...)

	return hex.EncodeToString(combined), nil
}

func Decrypt(key, cipherHex string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	if len(cipherHex) < nonceSize*2 {
		return "", ErrInvalidCipherHex
	}

	nonce, err := hex.DecodeString(cipherHex[:nonceSize*2])
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cipherHex[nonceSize*2:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
