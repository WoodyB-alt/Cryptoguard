package main

// Import the cmd package which houses all Cobra commands
import "github.com/WoodyB-alt/cryptoguard/cmd"

// main is the entry point of the application
func main() {
	// Calls the Execute function defined in cmd/root.go
	// which sets up and starts the Cobra CLI
	cmd.Execute()
}
