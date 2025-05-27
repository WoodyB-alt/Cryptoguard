package cmd

import (
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/WoodyB-alt/cryptoguard/internal/steg"
	"github.com/spf13/cobra"
)

var (
	stegExtractPassword string
	stegExtractImage    string
	stegExtractOutput   string
)

// stegExtractCmd defines the Cobra command for extracting and decrypting hidden content from a PNG image.
var stegExtractCmd = &cobra.Command{
	Use:   "steg-extract",
	Short: "Extract and decrypt hidden data from a PNG image",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate required flags
		if stegExtractPassword == "" || stegExtractImage == "" {
			fmt.Println("Usage: cryptoguard steg-extract -p \"password\" -img hidden.png [--out output.txt]")
			os.Exit(1)
		}

		// Step 1: Attempt to extract the hidden encrypted message
		encoded, err := steg.ExtractStringFromPNG(stegExtractImage)
		if err != nil {
			fmt.Println("Extraction error:", err)
			return
		}

		// Step 2: Decrypt the extracted message using AES-GCM and PBKDF2
		plain, err := crypto.DecryptAES(encoded, stegExtractPassword)
		if err != nil {
			fmt.Println("Decryption error:", err)
			return
		}

		// Step 3: If an output file was specified, save the result; otherwise print to console
		if stegExtractOutput != "" {
			err := os.WriteFile(stegExtractOutput, []byte(plain), 0644)
			if err != nil {
				fmt.Println("Failed to write output file:", err)
				return
			}
			fmt.Println("✅ Message decrypted and saved to:", stegExtractOutput)
		} else {
			fmt.Println("✅ Decrypted Message:")
			fmt.Println("--------------------------")
			fmt.Println(plain)
		}
	},
}

func init() {
	// Define CLI flags for steg-extract
	stegExtractCmd.Flags().StringVarP(&stegExtractPassword, "password", "p", "", "Password used during embedding (required)")
	stegExtractCmd.Flags().StringVar(&stegExtractImage, "img", "", "PNG image containing hidden message (required)")
	stegExtractCmd.Flags().StringVar(&stegExtractOutput, "out", "", "Optional file to write the decrypted result")

	// Mark required flags
	stegExtractCmd.MarkFlagRequired("password")
	stegExtractCmd.MarkFlagRequired("img")

	// Register with root
	rootCmd.AddCommand(stegExtractCmd)
}
