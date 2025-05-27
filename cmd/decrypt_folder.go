package cmd

import (
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/spf13/cobra"
)

var (
	decryptFolderPassword  string
	decryptFolderRecursive bool
)

// decryptFolderCmd is the Cobra command for recursively decrypting all .enc files in a directory.
var decryptFolderCmd = &cobra.Command{
	Use:   "decrypt-folder [input_dir] [output_dir]",
	Short: "Decrypt all encrypted files in a folder (optionally recursive)",
	Args:  cobra.ExactArgs(2), // Requires exactly two positional args: input and output directories
	Run: func(cmd *cobra.Command, args []string) {
		inputDir := args[0]
		outputDir := args[1]

		if decryptFolderPassword == "" {
			fmt.Println("Error: password is required. Use -p flag.")
			return
		}

		// Call the decryption function from the crypto package
		err := crypto.DecryptFolder(inputDir, outputDir, decryptFolderPassword, decryptFolderRecursive)
		if err != nil {
			fmt.Println("Folder decryption error:", err)
			return
		}

		fmt.Println("Folder decrypted successfully to:", outputDir)
	},
}

func init() {
	// Bind flags to variables
	decryptFolderCmd.Flags().StringVarP(&decryptFolderPassword, "password", "p", "", "Password for decryption (required)")
	decryptFolderCmd.Flags().BoolVar(&decryptFolderRecursive, "recursive", false, "Recursively decrypt subdirectories")
	decryptFolderCmd.MarkFlagRequired("password")

	// Register the command under the root
	rootCmd.AddCommand(decryptFolderCmd)
}
