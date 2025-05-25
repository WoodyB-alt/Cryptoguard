package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

// EncryptFile handles command-line file encryption.
// It supports optional deletion of the original file after successful encryption.
func EncryptFile(args []string) {
	// Define flags for password and optional deletion
	fs := flag.NewFlagSet("encrypt-file", flag.ExitOnError)
	password := fs.String("p", "", "Password for encryption")
	deleteOriginal := fs.Bool("delete-original", false, "Delete input file after encryption")
	fs.Parse(args)

	// Expecting exactly 2 arguments: input and output file paths
	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard encrypt-file -p \"password\" [--delete-original] input.txt output.enc")
		return
	}

	inputPath := fs.Arg(0)
	outputPath := fs.Arg(1)

	// Perform file encryption using the provided password
	err := crypto.EncryptFile(inputPath, outputPath, *password)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	// If --delete-original is set, attempt to delete the input file
	if *deleteOriginal {
		if err := os.Remove(inputPath); err != nil {
			fmt.Println("Warning: failed to delete original file:", err)
		} else {
			fmt.Println("Original file deleted:", inputPath)
		}
	}

	fmt.Println("File encrypted successfully:", outputPath)
}
