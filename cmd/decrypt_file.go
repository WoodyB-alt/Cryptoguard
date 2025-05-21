package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

// DecryptFile is the CLI handler for decrypting a file using AES.
// Usage: cryptoguard decrypt-file -p "password" input.enc output.txt
func DecryptFile(args []string) {
	// Define and parse the flag set for this command
	fs := flag.NewFlagSet("decrypt-file", flag.ExitOnError)
	password := fs.String("p", "", "Password for decryption")
	fs.Parse(args)

	// Ensure input and output file paths are provided
	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard decrypt-file -p \"password\" input.enc output.txt")
		return
	}

	inputPath := fs.Arg(0)
	outputPath := fs.Arg(1)

	// Call the decryption function from the crypto package
	err := crypto.DecryptFile(inputPath, outputPath, *password)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("File decrypted successfully:", outputPath)
}
