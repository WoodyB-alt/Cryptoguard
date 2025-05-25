package crypto

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

// EncryptFolder handles either zipped or recursive encryption of folders.
// If --zip is set, it compresses the folder, encrypts the zip, and optionally deletes the original.
func EncryptFolder(inputDir, outputDir, password string, recursive, deleteOriginal, zip bool) error {
	if zip {
		// Step 1: Define the .zip and .enc paths
		zipPath := filepath.Join(outputDir, filepath.Base(inputDir)+".zip")
		encryptedZipPath := zipPath + ".enc"

		// Ensure output folder exists
		if err := os.MkdirAll(filepath.Dir(zipPath), 0755); err != nil {
			return err
		}

		// Step 2: Zip the folder
		fmt.Println("Zipping folder:", inputDir)
		if err := ZipFolderToFile(inputDir, zipPath); err != nil {
			return err
		}

		// Step 3: Encrypt the .zip file
		fmt.Println("Encrypting zip:", zipPath)
		if err := EncryptFile(zipPath, encryptedZipPath, password); err != nil {
			return err
		}

		// Step 4: Optionally delete the original folder and zip file
		if deleteOriginal {
			if err := os.RemoveAll(inputDir); err != nil {
				fmt.Println("Warning: failed to delete original folder:", err)
			}
			if err := os.Remove(zipPath); err != nil {
				fmt.Println("Warning: failed to delete zip file:", err)
			}
		}

		fmt.Println("Zip encrypted successfully:", encryptedZipPath)
		return nil
	}

	// If --zip not set, fall back to standard folder encryption
	return standardFolderEncrypt(inputDir, outputDir, password, recursive, deleteOriginal)
}

// standardFolderEncrypt walks through the folder, encrypting each file individually.
// It supports recursive traversal and deletion of the original files.
// Also shows a progress bar.
func standardFolderEncrypt(inputDir, outputDir, password string, recursive, deleteOriginal bool) error {
	var files []string

	// Step 1: Collect all files (respecting recursive flag)
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

	// Step 3: Encrypt each file and update progress
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
			_ = os.Remove(path)
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

// DecryptZipFile decrypts an AES-GCM encrypted .zip.enc file and extracts it to a directory.
func DecryptZipFile(encPath, outputDir, password string, deleteOriginal bool) error {
	// Step 1: Decrypt to .zip file
	tmpZip := encPath[:len(encPath)-4] // remove .enc
	fmt.Println("Decrypting:", encPath)
	if err := DecryptFile(encPath, tmpZip, password); err != nil {
		return err
	}

	// Step 2: Unzip contents to output directory
	fmt.Println("Extracting zip:", tmpZip)
	if err := UnzipFileToDir(tmpZip, outputDir); err != nil {
		return err
	}

	// Step 3: Optionally delete the zip and .enc files
	if deleteOriginal {
		_ = os.Remove(tmpZip)
		_ = os.Remove(encPath)
		fmt.Println("Deleted original encrypted and zip files")
	}

	fmt.Println("Decryption and extraction complete:", outputDir)
	return nil
}
