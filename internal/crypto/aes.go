package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keyLen     = 32      // AES-256 requires a 32-byte key (256 bits)
	saltSize   = 16      // 128-bit salt for PBKDF2
	iterations = 100_000 // PBKDF2 iterations (recommended minimum)
)

// deriveKeyPBKDF2 uses PBKDF2 to derive a secure key from a password and salt.
// It performs multiple iterations of SHA-256 hashing to slow down brute-force attacks.
func deriveKeyPBKDF2(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLen, sha256.New)
}

// EncryptAES encrypts plaintext using AES-256 in CFB mode with a key derived from the password using PBKDF2.
// It generates a random salt, derives the key, and prepends the salt to the output.
func EncryptAES(plainText, password string) (string, error) {
	// Step 1: Generate a random salt
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Step 2: Derive a secure 256-bit AES key from the password and salt
	key := deriveKeyPBKDF2([]byte(password), salt)

	// Step 3: Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Step 4: Create a static IV (for now) — replace with random IV for stronger security
	iv := make([]byte, aes.BlockSize)
	for i := range iv {
		iv[i] = byte(i) // NOTE: Use crypto/rand for real-world scenarios
	}

	// Step 5: Prepare the ciphertext buffer (IV + encrypted content)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	copy(cipherText[:aes.BlockSize], iv)

	// Step 6: Create a CFB stream cipher and encrypt the plaintext
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	// Step 7: Prepend salt to the ciphertext and base64-encode the whole output
	final := append(salt, cipherText...)
	return base64.StdEncoding.EncodeToString(final), nil
}

// DecryptAES decrypts a base64-encoded string using AES-256 in CFB mode with a PBKDF2-derived key.
// It expects the input to be in the format: salt + iv + ciphertext.
func DecryptAES(cipherTextB64, password string) (string, error) {
	// Step 1: Base64 decode the input ciphertext
	raw, err := base64.StdEncoding.DecodeString(cipherTextB64)
	if err != nil {
		return "", err
	}

	// Step 2: Validate that the input is long enough to include salt and IV
	if len(raw) < saltSize+aes.BlockSize {
		return "", errors.New("invalid ciphertext")
	}

	// Step 3: Extract the salt from the input
	salt := raw[:saltSize]

	// Step 4: Extract the IV and encrypted data from the remaining bytes
	cipherData := raw[saltSize:]
	iv := cipherData[:aes.BlockSize]
	cipherBytes := cipherData[aes.BlockSize:]

	// Step 5: Derive the AES key using the extracted salt and provided password
	key := deriveKeyPBKDF2([]byte(password), salt)

	// Step 6: Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Step 7: Create a CFB stream cipher and decrypt the ciphertext in-place
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherBytes, cipherBytes)

	// Step 8: Return the decrypted plaintext as a string
	return string(cipherBytes), nil
}
