package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check if directory path is provided as command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go /path/to/directory")
		return
	}

	// Get the directory path from command-line argument
	dir := os.Args[1]

	// Read the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Counter for serializing IDs
	counter := 1

	// Loop through files in the directory
	for _, file := range files {
		// Check if the file is an image (png, jpg, jpeg)
		ext := filepath.Ext(file.Name())
		if strings.EqualFold(ext, ".png") || strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
			// Generate the new file name with serialized ID
			newName := fmt.Sprintf("%s-%d%s", dir, counter, ext)

			// Rename the file
			oldPath := filepath.Join(dir, file.Name())
			newPath := filepath.Join(dir, newName)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Printf("Error renaming file %s to %s: %v\n", oldPath, newPath, err)
			} else {
				fmt.Printf("Renamed %s to %s\n", oldPath, newPath)
			}

			// Increment the counter
			counter++
		}
	}
}
