package cmd

import (
	"fmt"

	"github.com/WoodyB-alt/cryptoguard/internal/crypto"
	"github.com/spf13/cobra"
)

var decryptPassword string

// decryptCmd defines the "decrypt" CLI command for decrypting base64-encoded AES-GCM ciphertext.
var decryptCmd = &cobra.Command{
	Use:   "decrypt [ciphertext]",
	Short: "Decrypt a base64-encoded AES-GCM string",
	Args:  cobra.ExactArgs(1), // Requires exactly one ciphertext argument
	Run: func(cmd *cobra.Command, args []string) {
		cipherText := args[0]

		if decryptPassword == "" {
			fmt.Println("Error: password is required. Use -p flag.")
			return
		}

		// Call internal decryption logic
		plainText, err := crypto.DecryptAES(cipherText, decryptPassword)
		if err != nil {
			fmt.Println("Decryption error:", err)
			return
		}

		// Output the result
		fmt.Println(plainText)
	},
}

func init() {
	// Bind -p / --password flag
	decryptCmd.Flags().StringVarP(&decryptPassword, "password", "p", "", "Password used for decryption")
	decryptCmd.MarkFlagRequired("password")

	// Register command with root
	rootCmd.AddCommand(decryptCmd)
}
