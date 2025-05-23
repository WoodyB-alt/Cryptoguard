package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

func DecryptFolder(args []string) {
	fs := flag.NewFlagSet("decrypt-folder", flag.ExitOnError)
	password := fs.String("p", "", "Password for decryption")
	recursive := fs.Bool("recursive", false, "Recursively decrypt all files in subdirectories")
	fs.Parse(args)

	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard decrypt-folder -p \"password\" [--recursive] <input_dir> <output_dir>")
		return
	}

	inputDir := fs.Arg(0)
	outputDir := fs.Arg(1)

	err := crypto.DecryptFolder(inputDir, outputDir, *password, *recursive)
	if err != nil {
		fmt.Println("Folder decryption error:", err)
	}
}
