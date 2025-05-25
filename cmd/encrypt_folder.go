package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

// EncryptFolder handles CLI command for encrypting a directory of files.
// Supports recursive walking and deletion of original files after encryption.
func EncryptFolder(args []string) {
	fs := flag.NewFlagSet("encrypt-folder", flag.ExitOnError)

	password := fs.String("p", "", "Password for encryption")
	recursive := fs.Bool("recursive", false, "Recursively encrypt all files in subdirectories")
	deleteOriginal := fs.Bool("delete-original", false, "Delete original files after encryption")
	fs.Parse(args)

	// Expecting exactly 2 arguments: input and output directories
	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard encrypt-folder -p \"password\" [--recursive] [--delete-original] input_dir output_dir")
		return
	}

	inputDir := fs.Arg(0)
	outputDir := fs.Arg(1)

	// Call the core folder encryption logic
	err := crypto.EncryptFolder(inputDir, outputDir, *password, *recursive, *deleteOriginal)
	if err != nil {
		fmt.Println("Folder encryption error:", err)
	}
}
