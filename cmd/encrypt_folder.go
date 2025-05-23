package cmd

import (
	"flag"
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

func EncryptFolder(args []string) {
	fs := flag.NewFlagSet("encrypt-folder", flag.ExitOnError)
	password := fs.String("p", "", "Password for encryption")
	recursive := fs.Bool("recursive", false, "Recursively encrypt all files in subdirectories")
	fs.Parse(args)

	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard encrypt-folder -p \"password\" [--recursive] <input_dir> <output_dir>")
		return
	}

	inputDir := fs.Arg(0)
	outputDir := fs.Arg(1)

	err := crypto.EncryptFolder(inputDir, outputDir, *password, *recursive)
	if err != nil {
		fmt.Println("Folder encryption error:", err)
	}
}
