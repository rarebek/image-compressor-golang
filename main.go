package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	file, err := os.Open("input.jpg")
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	outputFile, err := os.Create("compressed.jpg")
	if err != nil {
		fmt.Println("Error creating compressed image file:", err)
		return
	}
	defer outputFile.Close()

	jpegOptions := &jpeg.Options{Quality: 50}
	err = jpeg.Encode(outputFile, img, jpegOptions)
	if err != nil {
		fmt.Println("Error encoding image to JPEG:", err)
		return
	}

	fmt.Println("Image compressed and saved successfully as compressed.jpg")
}
