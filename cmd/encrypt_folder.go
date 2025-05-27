package cmd

import (
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/spf13/cobra"
)

var (
	encryptFolderPassword       string
	encryptFolderRecursive      bool
	encryptFolderDeleteOriginal bool
	encryptFolderZip            bool
)

// encryptFolderCmd defines the Cobra command for encrypting folders.
var encryptFolderCmd = &cobra.Command{
	Use:   "encrypt-folder <input_dir> <output_dir>",
	Short: "Encrypt all files in a folder using AES-GCM (optionally recursive, zipped, or delete source)",
	Args:  cobra.ExactArgs(2), // Input and output directories required
	Run: func(cmd *cobra.Command, args []string) {
		inputDir := args[0]
		outputDir := args[1]

		// Ensure password is provided
		if encryptFolderPassword == "" {
			fmt.Println("Error: password is required. Use -p flag.")
			return
		}

		// Call the core encryption logic
		err := crypto.EncryptFolder(
			inputDir,
			outputDir,
			encryptFolderPassword,
			encryptFolderRecursive,
			encryptFolderDeleteOriginal,
			encryptFolderZip,
		)

		if err != nil {
			fmt.Println("Folder encryption error:", err)
		} else {
			fmt.Println("Folder encrypted successfully:", outputDir)
		}
	},
}

func init() {
	// CLI flags for folder encryption
	encryptFolderCmd.Flags().StringVarP(&encryptFolderPassword, "password", "p", "", "Password for encryption (required)")
	encryptFolderCmd.Flags().BoolVar(&encryptFolderRecursive, "recursive", false, "Recursively encrypt files in subdirectories")
	encryptFolderCmd.Flags().BoolVar(&encryptFolderDeleteOriginal, "delete-original", false, "Delete original files/folders after encryption")
	encryptFolderCmd.Flags().BoolVar(&encryptFolderZip, "zip", false, "Zip the folder before encryption")
	encryptFolderCmd.MarkFlagRequired("password")

	// Register command
	rootCmd.AddCommand(encryptFolderCmd)
}
