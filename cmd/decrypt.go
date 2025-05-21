package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

func Decrypt(args []string) {
	// Define flag set for the decrypt command
	fs := flag.NewFlagSet("decrypt", flag.ExitOnError)
	password := fs.String("p", "", "Password for decryption")
	fs.Parse(args)

	// Get the ciphertext string to decrypt
	cipherText := fs.Arg(0)

	// Ensure both ciphertext and password are provided
	if *password == "" || cipherText == "" {
		fmt.Println("Usage: cryptoguard decrypt \"cipherText\" -p \"password\"")
		return
	}

	// Decrypt the ciphertext using AES
	plainText, err := crypto.DecryptAES(cipherText, *password)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	// Output the result
	fmt.Println(plainText)
}
