package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cryptoguard",
	Short: "Cryptoguard 🔐 - AES-256 encryption & steganography CLI tool",
	Long: `Cryptoguard is a secure AES-GCM encryption tool for files, folders, and text.
It includes support for steganography, zipping, and password-based protection.`,
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
		cmd.Help()
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// printBanner prints the ASCII logo and CLI info
func printBanner() {
	fmt.Println(`
	/$$$$$$                                  /$$                                                           /$$
	/$$__  $$                                | $$                                                          | $$
	| $$  \__/  /$$$$$$  /$$   /$$  /$$$$$$  /$$$$$$    /$$$$$$   /$$$$$$  /$$   /$$ /$$$$$$   /$$$$$$  /$$$$$$$
	| $$       /$$__  $$| $$  | $$ /$$__  $$|_  $$_/   /$$__  $$ /$$__  $$| $$  | $$|____  $$ /$$__  $$/$$__  $$
	| $$      | $$  \__/| $$  | $$| $$  \ $$  | $$    | $$  \ $$| $$  \ $$| $$  | $$ /$$$$$$$| $$  \__/ $$  | $$
	| $$    $$| $$      | $$  | $$| $$  | $$  | $$ /$$| $$  | $$| $$  | $$| $$  | $$/$$__  $$| $$     | $$  | $$
	|  $$$$$$/| $$      |  $$$$$$$| $$$$$$$/  |  $$$$/|  $$$$$$/|  $$$$$$$|  $$$$$$/  $$$$$$$| $$     |  $$$$$$$
	\______/ |__/       \____  $$| $$____/    \___/   \______/  \____  $$ \______/ \_______/|__/      \_______/
						/$$  | $$| $$                           /$$  \ $$
						|  $$$$$$/| $$                          |  $$$$$$/
						\______/ |__/                           \______/      🔐

 Secure AES-256 CLI Encryption Tool by Blake Wood
 https://github.com/WoodyB-alt/cryptoguard
`)
}
