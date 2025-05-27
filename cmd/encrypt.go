package cmd

import (
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto" // Import encryption logic
	"github.com/spf13/cobra"                            // Cobra provides CLI parsing
)

// encryptCmd defines the "encrypt" subcommand for the CLI.
// Usage: cryptoguard encrypt -p "password" "text to encrypt"
var encryptCmd = &cobra.Command{
	Use:   "encrypt [text]",                       // How the command is used
	Short: "Encrypt a text message using AES-GCM", // Short description for help output
	Args:  cobra.ExactArgs(1),                     // Requires exactly one argument (the plaintext)

	// Run is the function that executes when the encrypt command is called
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the input text to encrypt
		plaintext := args[0]

		// Ensure a password is provided
		if password == "" {
			fmt.Println("Error: password is required. Use -p flag.")
			return
		}

		// Encrypt the plaintext using AES-GCM and PBKDF2-derived key
		ciphertext, err := crypto.EncryptAES(plaintext, password)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}

		// Output the encrypted (base64-encoded) ciphertext
		fmt.Println(ciphertext)
	},
}

// password stores the user-provided password via the CLI flag
var password string

// init runs when the package is initialized
// It sets up the password flag and registers the encrypt command with the root CLI
func init() {
	// Define the "-p" or "--password" flag to capture the user's password input
	encryptCmd.Flags().StringVarP(&password, "password", "p", "", "Password for encryption")

	// Add the encrypt command to the root CLI so it's available at runtime
	rootCmd.AddCommand(encryptCmd)
}
