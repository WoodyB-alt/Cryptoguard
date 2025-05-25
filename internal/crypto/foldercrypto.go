package crypto

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

// EncryptFolder encrypts all files in a given directory using AES-GCM.
// Supports recursive walking, optional deletion, and displays a progress bar.
func EncryptFolder(inputDir, outputDir, password string, recursive, deleteOriginal bool) error {
	// Step 1: Collect list of files to encrypt
	var files []string

	err := filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if path != inputDir && !recursive {
				return filepath.SkipDir
			}
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return err
	}

	// Step 2: Initialize progress bar
	bar := progressbar.NewOptions(len(files),
		progressbar.OptionSetDescription("Encrypting files"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(40),
		progressbar.OptionSetPredictTime(true),
	)

	// Step 3: Encrypt each file and update progress bar
	for _, path := range files {
		relPath, _ := filepath.Rel(inputDir, path)
		outputPath := filepath.Join(outputDir, relPath+".enc")

		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		if err := EncryptFile(path, outputPath, password); err != nil {
			return err
		}

		if deleteOriginal {
			if err := os.Remove(path); err != nil {
				fmt.Println("Warning: failed to delete", path, ":", err)
			}
		}

		_ = bar.Add(1)
	}

	return nil
}

// DecryptFolder decrypts all files in a folder and writes them to outputDir.
// If recursive is true, it also processes subdirectories.
func DecryptFolder(inputDir, outputDir, password string, recursive bool) error {
	return filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories unless recursive is true
		if d.IsDir() {
			if path != inputDir && !recursive {
				return filepath.SkipDir
			}
			return nil
		}

		// Only process .enc files
		if filepath.Ext(path) != ".enc" {
			return nil
		}

		// Strip ".enc" and preserve relative path
		relPath, _ := filepath.Rel(inputDir, path)
		outputPath := filepath.Join(outputDir, relPath[:len(relPath)-4]) // remove .enc

		// Ensure output directory exists
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		fmt.Println("Decrypting:", path, "→", outputPath)
		return DecryptFile(path, outputPath, password)
	})
}
