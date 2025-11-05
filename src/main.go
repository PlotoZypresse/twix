package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	img1 := readImgBytes("test_images/a340-600.jpg")
	img2 := readImgBytes("test_images/a340-600_copy.jpg")

	img1hash := hashImgBytes(img1)
	img2hash := hashImgBytes(img2)

	compare := bytes.Compare(img1hash, img2hash)

	if compare == 0 {
		fmt.Println("Duplicate image")
	} else {
		fmt.Println("Not duplicate")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readImgBytes(imgPath string) []byte {
	data, err := os.ReadFile(imgPath)
	check(err)
	return data
}

func hashImgBytes(img []byte) []byte {
	hash := sha256.New()
	hash.Write(img)

	imgHash := hash.Sum(nil)
	return imgHash
}
