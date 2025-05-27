package cmd

import (
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/spf13/cobra"
)

var decryptFilePassword string

// decryptFileCmd represents the Cobra CLI command to decrypt a file using AES-GCM
var decryptFileCmd = &cobra.Command{
	Use:   "decrypt-file [input.enc] [output.txt]",
	Short: "Decrypt a file encrypted with AES-GCM",
	Args:  cobra.ExactArgs(2), // Expect exactly 2 arguments: input path and output path
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]  // Encrypted input file (e.g., "secret.enc")
		outputPath := args[1] // Destination file for decrypted content (e.g., "secret.txt")

		if decryptFilePassword == "" {
			fmt.Println("Error: password is required. Use -p flag.")
			return
		}

		// Perform decryption using the crypto package
		err := crypto.DecryptFile(inputPath, outputPath, decryptFilePassword)
		if err != nil {
			fmt.Println("Decryption error:", err)
			return
		}

		fmt.Println("File decrypted successfully:", outputPath)
	},
}

func init() {
	// Register the -p / --password flag for decryption
	decryptFileCmd.Flags().StringVarP(&decryptFilePassword, "password", "p", "", "Password for decryption (required)")
	decryptFileCmd.MarkFlagRequired("password")

	// Add this command to the root command
	rootCmd.AddCommand(decryptFileCmd)
}
