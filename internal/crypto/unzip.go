package crypto

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// UnzipFileToDir extracts a zip archive to the specified destination folder.
func UnzipFileToDir(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		path := filepath.Join(destDir, f.Name)

		if f.FileInfo().IsDir() {
			// Make directories as needed
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
			continue
		}

		// Create parent directories for the file
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		srcFile, err := f.Open()
		if err != nil {
			dstFile.Close()
			return err
		}

		_, err = io.Copy(dstFile, srcFile)

		dstFile.Close()
		srcFile.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
