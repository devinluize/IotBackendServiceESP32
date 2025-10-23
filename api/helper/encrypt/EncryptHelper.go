package encrypthelper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AES key and IV must both be 16 bytes for AES-128
var key = []byte("ThisIsAESkey1234")
var iv = []byte("ESP32InitVector1")

// DecryptAESCTR decrypts Base64 AES-CTR string back to plain text
func DecryptAESCTR(cipherB64 string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherData, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(cipherData))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plainText, cipherData)

	return string(plainText), nil
}

// EncryptAESCTR encrypts plain text using AES-CTR and returns Base64 string
func EncryptAESCTR(plainText string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintextBytes := []byte(plainText)
	ciphertext := make([]byte, len(plaintextBytes))

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintextBytes)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}
