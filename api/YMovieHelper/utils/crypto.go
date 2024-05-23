package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

// 復号不可
func CreateHMAC(message string) string {
	secretKey := os.Getenv("CryptoSecretKey")
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func Encrypt(plainText string) (string, error) {
	secretKey := os.Getenv("CryptoSecretKey")
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherText[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTextBytes)

	return hex.EncodeToString(cipherText), nil
}

func Decrypt(cipherTextHex string) (string, error) {
	secretKey := os.Getenv("CryptoSecretKey")
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	cipherText, _ := hex.DecodeString(cipherTextHex)
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
