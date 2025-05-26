package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/WoodyB-alt/cryptoguard/internal/steg"
)

// StegExtract handles extracting and decrypting a hidden message from a PNG image.
// It supports optionally writing the output to a file.
func StegExtract(args []string) {
	fs := flag.NewFlagSet("steg-extract", flag.ExitOnError)
	password := fs.String("p", "", "Password used during embedding")
	imgPath := fs.String("img", "", "PNG image containing the hidden message")
	outPath := fs.String("out", "", "Optional output file to save decrypted message")
	fs.Parse(args)

	if *password == "" || *imgPath == "" {
		fmt.Println("Usage: cryptoguard steg-extract -p \"password\" -img hidden.png [--out output.txt]")
		os.Exit(1)
	}

	// Step 1: Extract the encrypted message
	encoded, err := steg.ExtractStringFromPNG(*imgPath)
	if err != nil {
		fmt.Println("Extraction error:", err)
		return
	}

	// Step 2: Decrypt the embedded content
	plain, err := crypto.DecryptAES(encoded, *password)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	// Step 3: Output to file or print to console
	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(plain), 0644); err != nil {
			fmt.Println("Failed to write output file:", err)
			return
		}
		fmt.Println("✅ Message decrypted and saved to:", *outPath)
	} else {
		fmt.Println("✅ Decrypted Message:")
		fmt.Println("--------------------------")
		fmt.Println(plain)
	}
}
