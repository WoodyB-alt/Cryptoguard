package cmd

import (
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/spf13/cobra"
)

var (
	encryptFilePassword     string
	encryptFileDeleteSource bool
)

// encryptFileCmd defines the "encrypt-file" Cobra command.
var encryptFileCmd = &cobra.Command{
	Use:   "encrypt-file <input> <output>",
	Short: "Encrypt a file using AES-GCM",
	Args:  cobra.ExactArgs(2), // Requires input and output file paths
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		outputPath := args[1]

		if encryptFilePassword == "" {
			fmt.Println("Error: password is required. Use -p flag.")
			return
		}

		// Encrypt the file with AES-GCM
		err := crypto.EncryptFile(inputPath, outputPath, encryptFilePassword)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}

		// If --delete-original is enabled, remove the source file
		if encryptFileDeleteSource {
			if err := os.Remove(inputPath); err != nil {
				fmt.Println("Warning: failed to delete original file:", err)
			} else {
				fmt.Println("Original file deleted:", inputPath)
			}
		}

		fmt.Println("File encrypted successfully:", outputPath)
	},
}

func init() {
	// Define CLI flags
	encryptFileCmd.Flags().StringVarP(&encryptFilePassword, "password", "p", "", "Password for encryption")
	encryptFileCmd.Flags().BoolVar(&encryptFileDeleteSource, "delete-original", false, "Delete input file after encryption")
	encryptFileCmd.MarkFlagRequired("password")

	// Register with the root command
	rootCmd.AddCommand(encryptFileCmd)
}
