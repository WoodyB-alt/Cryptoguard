package cmd

import (
	"fmt"
	"os"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/spf13/cobra"
)

var (
	decryptZipPassword       string
	decryptZipDeleteOriginal bool
)

// decryptZipCmd defines the Cobra command to decrypt a .zip.enc file and extract its contents.
var decryptZipCmd = &cobra.Command{
	Use:   "decrypt-zip [encrypted.zip.enc] [output_dir]",
	Short: "Decrypt an encrypted .zip file and extract its contents",
	Args:  cobra.ExactArgs(2), // Requires encrypted file and output directory
	Run: func(cmd *cobra.Command, args []string) {
		encPath := args[0]
		outputDir := args[1]

		if decryptZipPassword == "" {
			fmt.Println("Error: password is required. Use -p flag.")
			return
		}

		// Perform decryption and extraction using internal logic
		err := crypto.DecryptZipFile(encPath, outputDir, decryptZipPassword, decryptZipDeleteOriginal)
		if err != nil {
			fmt.Println("DecryptZip error:", err)
			os.Exit(1)
		}

		fmt.Println("Zip archive decrypted and extracted to:", outputDir)
	},
}

func init() {
	// Register CLI flags
	decryptZipCmd.Flags().StringVarP(&decryptZipPassword, "password", "p", "", "Password to decrypt the zip file (required)")
	decryptZipCmd.Flags().BoolVar(&decryptZipDeleteOriginal, "delete-original", false, "Delete .zip and .enc files after successful extraction")
	decryptZipCmd.MarkFlagRequired("password")

	// Add this command to the root
	rootCmd.AddCommand(decryptZipCmd)
}
