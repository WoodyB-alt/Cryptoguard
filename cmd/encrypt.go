package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

// Encrypt handles the 'encrypt' CLI command.
// It accepts a plaintext message and a password flag, then prints the AES-encrypted output.
func Encrypt(args []string) {
	// Define a flag set for the 'encrypt' command and parse the provided args
	fs := flag.NewFlagSet("encrypt", flag.ExitOnError)
	password := fs.String("p", "", "Password for encryption")
	fs.Parse(args)

	// Get the plaintext string from remaining args
	plainText := fs.Arg(0)

	// Validate that both password and plaintext were provided
	if *password == "" || plainText == "" {
		fmt.Println("Usage: cryptoguard encrypt \"text\" -p \"password\"")
		return
	}

	// Encrypt the plaintext using AES-256 encryption
	cipherText, err := crypto.EncryptAES(plainText, *password)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	// Print the base64-encoded ciphertext
	fmt.Println(cipherText)
}
