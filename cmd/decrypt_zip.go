package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
)

// DecryptZip handles decryption of an encrypted .zip.enc file and extraction of its contents.
func DecryptZip(args []string) {
	fs := flag.NewFlagSet("decrypt-zip", flag.ExitOnError)
	password := fs.String("p", "", "Password for decryption")
	deleteOriginal := fs.Bool("delete-original", false, "Delete .zip and .enc file after successful extraction")
	fs.Parse(args)

	if fs.NArg() < 2 {
		fmt.Println("Usage: cryptoguard decrypt-zip -p \"password\" --delete-original encrypted.zip.enc output_dir")
		return
	}

	encPath := fs.Arg(0)
	outputDir := fs.Arg(1)

	err := crypto.DecryptZipFile(encPath, outputDir, *password, *deleteOriginal)
	if err != nil {
		fmt.Println("DecryptZip error:", err)
		os.Exit(1)
	}
}
