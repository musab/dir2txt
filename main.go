package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Helper function to check if a file is binary.
// Here, we assume binary files are those that contain a null byte (0x00).
func isBinaryFile(filePath string) (bool, error) {
	data, err := os.ReadFile(filePath) // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return false, err
	}
	for _, b := range data {
		if b == 0 {
			return true, nil
		}
	}
	return false, nil
}

// Function to process files and print their contents.
func processFile(path string, info os.FileInfo) error {
	if info.IsDir() {
		return nil
	}

	// Check if the file is binary
	isBinary, err := isBinaryFile(path)
	if err != nil {
		return err
	}

	if !isBinary {
		data, err := os.ReadFile(path) // Use os.ReadFile instead of ioutil.ReadFile
		if err != nil {
			return err
		}

		fmt.Println("---")
		fmt.Printf("File: %s\n", path)
		fmt.Println("---")
		fmt.Println(string(data))
		fmt.Println()
	}

	return nil
}

// Main function to traverse the directory and process each file.
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a directory path")
	}

	dirPath := os.Args[1]

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return processFile(path, info)
	})
	if err != nil {
		log.Fatal(err)
	}
}
