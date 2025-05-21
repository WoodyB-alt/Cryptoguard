package main

import (
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/cmd"
)

func main() {
	// Check if at least one argument (command) is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: cryptoguard [encrypt|decrypt] [text] -p [password]")
		os.Exit(1)
	}

	// Choose between encrypt and decrypt commands
	switch os.Args[1] {
	case "encrypt":
		cmd.Encrypt(os.Args[2:])
	case "decrypt":
		cmd.Decrypt(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
