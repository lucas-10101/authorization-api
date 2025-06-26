package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

// EncryptAES encrypts the plaintext using AES encryption with the provided key.
// The key must be either 16, 24, or 32 bytes long.
// It returns the base64 encoded ciphertext or an error if encryption fails.
func EncryptAES(plaintext string) (string, error) {

	var key []byte = []byte(os.Getenv("AES32_ENCRYPTION_KEY"))

	block, err := aes.NewCipher(key)
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
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES decrypts the base64 encoded ciphertext using AES decryption with the provided key.
// The key must be either 16, 24, or 32 bytes long.
// It returns the decrypted plaintext or an error if decryption fails.
func DecryptAES(encrypted string) (string, error) {
	var key []byte = []byte(os.Getenv("AES32_ENCRYPTION_KEY"))

	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
