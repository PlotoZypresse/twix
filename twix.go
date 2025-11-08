package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/corona10/goimagehash"
)

type dup_img struct {
	imgPath1 string
	imgPath2 string
}

func main() {

	folderPath := os.Args[1]
	now := time.Now()
	checkDupes(1, folderPath)
	elapsed := time.Since(now)
	fmt.Println("Finding duplicates took: ", elapsed)
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
}

func pHashCompare(img1 *goimagehash.ImageHash, img2 *goimagehash.ImageHash) int {
	distance, _ := img1.Distance(img2)
	return distance
}

func storeImgHashes(hash []byte, fileName string, hashMap map[string]string) *dup_img {
	hashStr := string(hash)
	val, ok := hashMap[hashStr]
	if ok {
		return &dup_img{
			imgPath1: fileName,
			imgPath2: val,
		}

	} else {
		hashMap[hashStr] = fileName
		return nil
	}
}

func storePHashes(imgPHash *goimagehash.ImageHash, fileName string) {

}

// Function that compares all iamges from a folder.
// depending on the operation input it compares
// only the hashes, only the pHashes or both.
func checkDupes(operation int, folder string) {
	duplicateImgs := []dup_img{}

	switch operation {
	case 1: // only hash
		var hashMap = make(map[string]string)
		err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			imgBytes := readImgBytes(path)
			hash := hashImgBytes(imgBytes)
			val := storeImgHashes(hash, path, hashMap)

			if val != nil {
				duplicateImgs = append(duplicateImgs, *val)
			}

			return nil

		})
		check(err)
		prettyPrint(duplicateImgs)
	case 2: // only phash

	case 3: // phash and has

	default: // default is hash
	}
}

func prettyPrint(input []dup_img) {
	for _, item := range input {
		fmt.Println("Duplicate images found at:", item.imgPath1, "and", item.imgPath2)
	}
}
