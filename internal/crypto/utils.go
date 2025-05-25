package crypto

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// PrintError is a simple helper to format errors
func PrintError(prefix string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", prefix, err)
	}
}

// ZipFolderToFile compresses the contents of a folder (srcDir) into a .zip file (destZipPath).
// It preserves relative paths and skips subdirectories (folders themselves are not stored, only contents).
func ZipFolderToFile(srcDir, destZipPath string) error {
	// Create the destination zip file
	zipFile, err := os.Create(destZipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a zip writer
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// Walk through all files in the source folder
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories (zip only stores files)
		if info.IsDir() {
			return nil
		}

		// Calculate relative path for zip structure
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		// Create a new file entry in the archive
		zipWriter, err := archive.Create(relPath)
		if err != nil {
			return err
		}

		// Open the source file and copy its contents into the zip archive
		fsFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fsFile.Close()

		_, err = io.Copy(zipWriter, fsFile)
		return err
	})

	return err
}
