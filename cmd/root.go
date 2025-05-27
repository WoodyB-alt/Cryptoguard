package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cryptoguard",
	Short: "Cryptoguard 🔐 - AES-256 encryption & steganography CLI tool",
	Long: `Cryptoguard is a secure AES-GCM encryption tool for files, folders, and text, 
	with steganography, zipping, and password-based protection.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
