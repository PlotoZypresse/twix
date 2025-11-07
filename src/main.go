package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/corona10/goimagehash"
)

func main() {

	img1 := readImgBytes("test_images/a340-600.jpg")
	img2 := readImgBytes("test_images/a340-600_copy.jpg")
	img3 := readImgBytes("test_images/concord.jpg")

	img1hash := hashImgBytes(img1)
	img2hash := hashImgBytes(img2)
	// img3hash := hashImgBytes(img3)

	compare := bytes.Compare(img1hash, img2hash)

	if compare == 0 {
		fmt.Println("Duplicate image")
	} else {
		fmt.Println("Not duplicate")
	}

	img1pHash := pHashImgBytes(img1)
	// img2pHash := pHashImgBytes(img2)
	img3pHash := pHashImgBytes(img3)
	distance := pHashCompare(img1pHash, img3pHash)
	fmt.Printf("Distance between passed images: %v\n", distance)
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

func storeImgHashes(hash []byte, fileName string) {

}

// Function to create a perceptual hash (phash) of an image.
// Takes the bytes of an image as input. Decodes it to an Image interface,
// and creates an pHash value. To return bytes again, the code allocates a
// buffer *b*, creates a Writer to write to *b*. The img pHash is written
// as bytes to *b* and flushed to be returned
func pHashImgBytes(img []byte) *goimagehash.ImageHash {
	imgDecode, _, err := image.Decode(bytes.NewReader(img))
	check(err)
	imgHash, _ := goimagehash.PerceptionHash(imgDecode)

	return imgHash

	// var b bytes.Buffer
	// imgBuf := bufio.NewWriter(&b)
	// _ = imgHash.Dump(imgBuf)
	// imgBuf.Flush()
	// imgBytes := b.Bytes()
	// return imgBytes
}

func pHashCompare(img1 *goimagehash.ImageHash, img2 *goimagehash.ImageHash) int {
	distance, _ := img1.Distance(img2)
	return distance
}
