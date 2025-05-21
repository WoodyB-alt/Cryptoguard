package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// deriveKey creates a 32-byte AES key from the password using SHA-256 hashing
func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:] // Return the first 32 bytes (256 bits)
}

// EncryptAES encrypts the given plaintext using AES-256 with CFB mode
func EncryptAES(text string, password string) (string, error) {
	key := deriveKey(password)

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainText := []byte(text)

	// Allocate space for IV + ciphertext
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	// Dummy IV generation — for production use a random IV (e.g., crypto/rand)
	for i := range iv {
		iv[i] = byte(i)
	}

	// Create a stream and encrypt the plaintext
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	// Return base64-encoded ciphertext
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptAES decrypts the given base64-encoded ciphertext using AES-256
func DecryptAES(cryptoText string, password string) (string, error) {
	key := deriveKey(password)

	// Decode the base64-encoded ciphertext
	cipherText, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Ensure the ciphertext is valid length
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// Separate IV and actual ciphertext
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Create a stream and decrypt
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
