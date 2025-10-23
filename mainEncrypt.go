package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

// AES key and IV must both be 16 bytes for AES-128
var key = []byte("ThisIsAESkey1234")
var iv = []byte("ESP32InitVector1")

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

func main2() {
	//// Example usage
	//msg := "hello golang service"
	//encrypted, _ := EncryptAESCTR(msg)
	//fmt.Println("Encrypted (Base64):", encrypted)
	//
	//fmt.Println("Decrypted:", decrypted)
	key := []byte("ThisIsAESkey1234")
	iv := []byte("ESP32InitVector1")

	cipherB64 := strings.TrimSpace(`
iDTzNqD1FZwJCroy0LWEXKIfcKNCju6wKSY0iXylLfujPauLMw4NIdrishQqRW63Pff02Z+jSl4RkeFr4tk0qYoeLpEERLnTLHqRNd+5RXK5ZOJ0ZXQ3oC859aH82uJ1Mwm0uQ==
`)
	decrypted, _ := DecryptAESCTR(cipherB64)
	fmt.Println("Cipher decryptedlength:", len(decrypted)) // should equal plaintext length (≈97)
	fmt.Println(string(decrypted))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherData, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cipher length:", len(cipherData)) // should equal plaintext length (≈97)

	stream := cipher.NewCTR(block, iv)
	plain := make([]byte, len(cipherData))
	stream.XORKeyStream(plain, cipherData)

	fmt.Println("Decrypted text:")
	fmt.Println(string(plain))
}
