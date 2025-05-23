package crypto

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// EncryptFolder encrypts all files in a folder and writes them to outputDir.
// If recursive is true, it also processes subdirectories.
func EncryptFolder(inputDir, outputDir, password string, recursive bool) error {
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

		// Get relative path to maintain folder structure
		relPath, _ := filepath.Rel(inputDir, path)
		outputPath := filepath.Join(outputDir, relPath+".enc")

		// Ensure output directory exists
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		fmt.Println("Encrypting:", path, "→", outputPath)
		return EncryptFile(path, outputPath, password)
	})
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
