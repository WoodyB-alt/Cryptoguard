package main

import (
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/cmd"
)

func main() {
	// Ensure at least one CLI argument is provided (the command)
	if len(os.Args) < 2 {
		fmt.Println("Usage: cryptoguard [encrypt|decrypt|encrypt-file|decrypt-file] ...")
		os.Exit(1)
	}

	// Dispatch based on the first CLI argument (the command)
	switch os.Args[1] {
	case "encrypt":
		// Encrypt a text string
		cmd.Encrypt(os.Args[2:])
	case "decrypt":
		// Decrypt a text string
		cmd.Decrypt(os.Args[2:])
	case "encrypt-file":
		// Encrypt a file
		cmd.EncryptFile(os.Args[2:])
	case "decrypt-file":
		// Decrypt a file
		cmd.DecryptFile(os.Args[2:])
	default:
		// Handle unknown commands
		fmt.Println("Unknown command:", os.Args[1])
	}
}
