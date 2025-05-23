package main

import (
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/cmd"
)

// printHelp prints usage instructions for the CLI
func printHelp() {
	fmt.Println(`
	Cryptoguard 🔐 - AES-256 CLI encryption tool

	Usage:
	cryptoguard <command> [options]

	Available Commands:
	encrypt        Encrypt a string with a password
	decrypt        Decrypt a base64 string with a password
	encrypt-file   Encrypt a file with a password
	decrypt-file   Decrypt a file with a password
	help           Show this help message

	Examples:
	cryptoguard encrypt -p "mypassword" "Secret text"
	cryptoguard decrypt -p "mypassword" "base64cipher"
	cryptoguard encrypt-file -p "mypassword" input.txt output.enc
	cryptoguard decrypt-file -p "mypassword" input.enc output.txt
	`)
}

func main() {
	// Show help if no arguments or explicitly requested
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "help" {
		printHelp()
		return
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
		printHelp()
	}
}
