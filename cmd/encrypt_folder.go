package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

// EncryptFolder handles CLI input for encrypting folders.
// Supports optional recursive walking, zipping, and original file deletion.
func EncryptFolder(args []string) {
	// Define CLI flags
	fs := flag.NewFlagSet("encrypt-folder", flag.ExitOnError)
	password := fs.String("p", "", "Password for encryption")
	recursive := fs.Bool("recursive", false, "Recursively encrypt all files in subdirectories")
	deleteOriginal := fs.Bool("delete-original", false, "Delete original files/folders after encryption")
	zip := fs.Bool("zip", false, "Zip the entire folder before encrypting")
	fs.Parse(args)

	// Require input and output directory paths
	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard encrypt-folder -p \"password\" [--recursive] [--delete-original] [--zip] input_dir output_dir")
		return
	}

	inputDir := fs.Arg(0)
	outputDir := fs.Arg(1)

	// Invoke the core encryption function with all options
	err := crypto.EncryptFolder(inputDir, outputDir, *password, *recursive, *deleteOriginal, *zip)
	if err != nil {
		fmt.Println("Folder encryption error:", err)
	}
}
