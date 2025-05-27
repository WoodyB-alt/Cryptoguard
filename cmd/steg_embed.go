package cmd

import (
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/WoodyB-alt/cryptoguard/internal/steg"
	"github.com/spf13/cobra"
)

var (
	stegEmbedPassword string
	stegEmbedInput    string
	stegEmbedImage    string
	stegEmbedOutput   string
)

// stegEmbedCmd defines the Cobra command for steganographic embedding.
var stegEmbedCmd = &cobra.Command{
	Use:   "steg-embed",
	Short: "Encrypt and hide text/file data inside a PNG image using steganography",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate required flags
		if stegEmbedPassword == "" || stegEmbedInput == "" || stegEmbedImage == "" || stegEmbedOutput == "" {
			fmt.Println("Usage: cryptoguard steg-embed -p \"password\" -in secret.txt -img carrier.png -out hidden.png")
			os.Exit(1)
		}

		// Step 1: Read the input file to hide
		data, err := os.ReadFile(stegEmbedInput)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}

		// Step 2: Encrypt the data using AES-GCM and PBKDF2
		encrypted, err := crypto.EncryptAES(string(data), stegEmbedPassword)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}

		// Step 3: Append a special end marker so we can extract it later
		encrypted += "<<<END>>>"

		// Step 4: Embed the encrypted message into the image using LSB steganography
		err = steg.EmbedStringInPNG(stegEmbedImage, stegEmbedOutput, encrypted)
		if err != nil {
			fmt.Println("Steganography error:", err)
			return
		}

		fmt.Println("Encrypted data successfully embedded in:", stegEmbedOutput)
	},
}

func init() {
	// Required CLI flags
	stegEmbedCmd.Flags().StringVarP(&stegEmbedPassword, "password", "p", "", "Password for encryption (required)")
	stegEmbedCmd.Flags().StringVar(&stegEmbedInput, "in", "", "Input file to hide (required)")
	stegEmbedCmd.Flags().StringVar(&stegEmbedImage, "img", "", "Carrier PNG image file (required)")
	stegEmbedCmd.Flags().StringVar(&stegEmbedOutput, "out", "", "Output PNG file with embedded data (required)")

	// Mark all flags as required for this command
	stegEmbedCmd.MarkFlagRequired("password")
	stegEmbedCmd.MarkFlagRequired("in")
	stegEmbedCmd.MarkFlagRequired("img")
	stegEmbedCmd.MarkFlagRequired("out")

	// Register command to root
	rootCmd.AddCommand(stegEmbedCmd)
}
