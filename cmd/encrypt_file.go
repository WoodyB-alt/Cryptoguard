package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

// EncryptFile is the CLI handler for encrypting a file using AES.
// Usage: cryptoguard encrypt-file -p "password" input.txt output.enc
func EncryptFile(args []string) {
	// Define and parse the flag set for this command
	fs := flag.NewFlagSet("encrypt-file", flag.ExitOnError)
	password := fs.String("p", "", "Password for encryption")
	fs.Parse(args)

	// Ensure input and output file paths are provided
	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard encrypt-file -p \"password\" input.txt output.enc")
		return
	}

	inputPath := fs.Arg(0)
	outputPath := fs.Arg(1)

	// Call the encryption function from the crypto package
	err := crypto.EncryptFile(inputPath, outputPath, *password)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	fmt.Println("File encrypted successfully:", outputPath)
}
