package crypto

import (
	"crypto/aes"      // Provides AES cipher implementation
	"crypto/cipher"   // Provides GCM mode for authenticated encryption
	"crypto/rand"     // Secure random number generation
	"crypto/sha256"   // SHA-256 hash function for PBKDF2
	"encoding/base64" // Base64 encoding for encrypted output
	"errors"          // Error handling

	"golang.org/x/crypto/pbkdf2" // PBKDF2 key derivation implementation
)

const (
	keyLen     = 32      // AES-256 requires a 32-byte (256-bit) key
	saltSize   = 16      // 128-bit salt for PBKDF2 key derivation
	nonceSize  = 12      // 96-bit nonce (recommended size for AES-GCM)
	iterations = 100_000 // Number of PBKDF2 iterations
)

// deriveKeyPBKDF2 generates a strong encryption key from a password and salt using PBKDF2.
// This function slows down brute-force attempts by using many hash iterations.
func deriveKeyPBKDF2(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLen, sha256.New)
}

// EncryptAES encrypts the input plaintext using AES-256 in GCM mode.
// It returns a base64-encoded string containing the salt, nonce, and ciphertext.
func EncryptAES(plainText, password string) (string, error) {
	// Step 1: Generate a random salt for key derivation
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Step 2: Derive a secure key from the password and salt using PBKDF2
	key := deriveKeyPBKDF2([]byte(password), salt)

	// Step 3: Initialize the AES cipher block with the derived key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Step 4: Wrap the cipher block with AES-GCM for authenticated encryption
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Step 5: Generate a random nonce (used once per encryption)
	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// Step 6: Encrypt the plaintext using AES-GCM
	// - Includes authentication tag
	// - No additional data used (set to nil)
	cipherText := gcm.Seal(nil, nonce, []byte(plainText), nil)

	// Step 7: Combine salt + nonce + ciphertext for final output
	final := append(salt, nonce...)
	final = append(final, cipherText...)

	// Step 8: Base64-encode the final encrypted blob for safe text handling
	return base64.StdEncoding.EncodeToString(final), nil
}

// DecryptAES decrypts a base64-encoded string that was encrypted using EncryptAES.
// It validates the authentication tag to ensure data integrity and correct password.
func DecryptAES(cipherTextB64, password string) (string, error) {
	// Step 1: Base64-decode the encrypted input string
	raw, err := base64.StdEncoding.DecodeString(cipherTextB64)
	if err != nil {
		return "", err
	}

	// Step 2: Validate that the input is long enough to contain salt + nonce
	if len(raw) < saltSize+nonceSize {
		return "", errors.New("invalid encrypted data")
	}

	// Step 3: Extract the components from the decoded input
	salt := raw[:saltSize]                      // First 16 bytes
	nonce := raw[saltSize : saltSize+nonceSize] // Next 12 bytes
	cipherData := raw[saltSize+nonceSize:]      // Remaining bytes = ciphertext + auth tag

	// Step 4: Derive the same key using the extracted salt and provided password
	key := deriveKeyPBKDF2([]byte(password), salt)

	// Step 5: Initialize AES block cipher with the derived key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Step 6: Wrap the cipher block with AES-GCM for decryption
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Step 7: Attempt to decrypt and authenticate the ciphertext
	plainText, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		// If authentication fails, it's likely the password is wrong or the data was tampered with
		return "", errors.New("decryption failed: possibly incorrect password or tampered data")
	}

	// Step 8: Return the original plaintext as a string
	return string(plainText), nil
}
