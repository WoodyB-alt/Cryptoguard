package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

// EncryptFile encrypts the contents of a file using AES-GCM and writes to a new file.
// Output format: salt + nonce + base64(ciphertext)
func EncryptFile(inputPath, outputPath, password string) error {
	// Open the input file for reading
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Generate random salt
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return err
	}

	// Derive key from password and salt using PBKDF2
	key := pbkdf2.Key([]byte(password), salt, iterations, keyLen, sha256.New)

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create AES-GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Generate a random nonce
	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return err
	}

	// Read input file into memory
	plainData, err := io.ReadAll(inFile)
	if err != nil {
		return err
	}

	// Encrypt the plaintext data
	cipherData := gcm.Seal(nil, nonce, plainData, nil)

	// Combine salt + nonce + ciphertext
	final := append(salt, nonce...)
	final = append(final, cipherData...)

	// Base64 encode the result before writing to file
	encoder := base64.NewEncoder(base64.StdEncoding, outFile)
	defer encoder.Close()

	_, err = encoder.Write(final)
	return err
}

// DecryptFile decrypts a file encrypted with AES-GCM and PBKDF2.
// Input format must be: base64(salt + nonce + ciphertext)
func DecryptFile(inputPath, outputPath, password string) error {
	// Open the encrypted file
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Base64 decode the entire input
	decoder := base64.NewDecoder(base64.StdEncoding, inFile)
	rawData, err := io.ReadAll(decoder)
	if err != nil {
		return err
	}

	if len(rawData) < saltSize+nonceSize {
		return errors.New("invalid encrypted file format")
	}

	// Extract salt, nonce, and ciphertext
	salt := rawData[:saltSize]
	nonce := rawData[saltSize : saltSize+nonceSize]
	cipherData := rawData[saltSize+nonceSize:]

	// Derive key from password and salt
	key := pbkdf2.Key([]byte(password), salt, iterations, keyLen, sha256.New)

	// Create AES cipher and GCM mode
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Decrypt and verify the ciphertext
	plainData, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return errors.New("decryption failed: possibly incorrect password or tampered file")
	}

	// Write decrypted data to output file
	_, err = outFile.Write(plainData)
	return err
}
