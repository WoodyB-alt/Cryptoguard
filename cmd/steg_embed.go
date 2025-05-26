package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/WoodyB-alt/cryptoguard/internal/steg"
)

// StegEmbed handles hiding encrypted text or file contents inside a PNG image.
func StegEmbed(args []string) {
	fs := flag.NewFlagSet("steg-embed", flag.ExitOnError)
	password := fs.String("p", "", "Password for encryption")
	inPath := fs.String("in", "", "Input file (text or any file to embed)")
	imgPath := fs.String("img", "", "Carrier PNG image")
	outPath := fs.String("out", "", "Output image with embedded data")
	fs.Parse(args)

	if *password == "" || *inPath == "" || *imgPath == "" || *outPath == "" {
		fmt.Println("Usage: cryptoguard steg-embed -p \"password\" -in secret.txt -img carrier.png -out hidden.png")
		os.Exit(1)
	}

	// Step 1: Read the input file and encrypt it
	data, err := os.ReadFile(*inPath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	encrypted, err := crypto.EncryptAES(string(data), *password)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	encrypted += "<<<END>>>"

	// Step 2: Embed encrypted string into the PNG
	if err := steg.EmbedStringInPNG(*imgPath, *outPath, encrypted); err != nil {
		fmt.Println("Steganography error:", err)
		return
	}

	fmt.Println("Encrypted data successfully embedded in:", *outPath)
}
