package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
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
			// Resize the image
			imagePath := filepath.Join(dir, file.Name())
			if err := resizeImage(imagePath); err != nil {
				fmt.Printf("Error resizing image %s: %v\n", imagePath, err)
				continue
			}

			// Get the new file name with serialized ID
			newName := fmt.Sprintf("%s-%d%s", dir, counter, ext)

			// Rename the file
			oldPath := imagePath
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

func resizeImage(imagePath string) error {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// Resize the image to 300x300 pixels using Lanczos resampling
	resizedImg := resize.Resize(300, 300, img, resize.Lanczos3)

	// Overwrite the original file with the resized image
	outFile, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode the resized image and write it to the output file
	err = jpeg.Encode(outFile, resizedImg, nil)
	if err != nil {
		return err
	}

	return nil
}
