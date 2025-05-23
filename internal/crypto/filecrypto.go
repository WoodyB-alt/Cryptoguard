package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

// deriveKey creates a 256-bit key from a password using SHA-256.
// This is a quick method but could be replaced by PBKDF2/scrypt later.

// EncryptFile encrypts a file using AES in CFB mode and saves the result to a new file.
// The IV is randomly generated and prepended to the output.
// The output is base64-encoded.
func EncryptFile(inputPath, outputPath, password string) error {
	// Open the input file
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

	// Derive the encryption key from the password
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return err
	}
	key := deriveKeyPBKDF2([]byte(password), salt)

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Generate a secure random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	// Write Salt and IV to the start of the output file (not encrypted)
	outFile.Write(salt)
	outFile.Write(iv)

	// Wrap the writer with AES stream and base64 encoder
	stream := cipher.NewCFBEncrypter(block, iv)
	writer := &cipher.StreamWriter{
		S: stream,
		W: base64.NewEncoder(base64.StdEncoding, outFile),
	}

	// Copy file contents through the encryption writer
	_, err = io.Copy(writer, inFile)
	if err != nil {
		return err
	}

	return nil
}

// DecryptFile reads an encrypted file, extracts the IV, and decrypts the contents.
// Assumes IV is prepended and contents are base64 encoded.
func DecryptFile(inputPath, outputPath, password string) error {
	// Open the encrypted input file
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Create the decrypted output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Derive the key from the password
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(inFile, salt); err != nil {
		return err
	}
	key := deriveKeyPBKDF2([]byte(password), salt)

	// Read the IV (first 16 bytes of the input file)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(inFile, iv); err != nil {
		return err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Wrap the input with a base64 decoder and decrypt stream
	stream := cipher.NewCFBDecrypter(block, iv)
	reader := &cipher.StreamReader{
		S: stream,
		R: base64.NewDecoder(base64.StdEncoding, inFile),
	}

	// Copy the decrypted content to the output file
	_, err = io.Copy(outFile, reader)
	if err != nil {
		return err
	}

	return nil
}
