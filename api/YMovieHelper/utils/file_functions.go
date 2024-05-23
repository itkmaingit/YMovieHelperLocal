package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func SaveFile(r io.Reader, filePath string) error {
	// Ensure directory exists
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to file
	_, err = io.Copy(file, r)
	if err != nil {
		return err
	}

	return nil
}

// ReadFile reads the file at the given path and returns its content as []byte
func ReadFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("GetLocalFile: %w", err)
	}
	return data, nil
}

func MoveFile(oldPath string, newPath string) error {
	// Create the directory for the new file path if it doesn't exist
	newDir := filepath.Dir(newPath)
	if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
		return fmt.Errorf("MoveFile: %w", err)
	}

	// Rename (move) the file
	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Errorf("MoveFile: %w", err)
	}

	return nil
}

// DeleteFile deletes a file at the given path
func DeleteFile(filePath string) error {
	// Remove the file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("DeleteFile: %w", err)
	}
	return nil
}
