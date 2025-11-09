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

type store_phash struct {
	imgPHash *goimagehash.ImageHash
	filename string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: twix <folder path> [mode]")
		os.Exit(1)
	}

	folderPath := os.Args[1]
	flagIn := ""
	if len(os.Args) > 2 {
		flagIn = os.Args[2]
	}

	flag := inputFlag(flagIn)
	now := time.Now()
	checkDupes(flag, folderPath)
	elapsed := time.Since(now)
	fmt.Println("Finding duplicates took: ", elapsed)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func inputFlag(flag string) int {
	switch flag {
	case "-h":
		return 1
	case "-p":
		return 2
	case "-hp":
		return 3
	default:
		return 1
	}
}

// reads the bytes of the image input and returns them.
// takes a path to an image as input. Returns []bytes.
func readImgBytes(imgPath string) ([]byte, error) {
	data, err := os.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Function that hashes the bytes of an image.
// takes []bytes as input and returns []bytes.
func hashImgBytes(img []byte) []byte {
	hash := sha256.New()
	hash.Write(img)

	imgHash := hash.Sum(nil)
	return imgHash
}

// Function to store image hashes and the corresponding file name in a map.
// on collision both file names are stored in a struct and returned.
// Takes the file hash(bytes), the file name as a string and a map[string]string
// as input
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

// Function to create a perceptual hash (phash) of an image.
// Takes the bytes of an image as input. Decodes it to an Image interface,
// and creates an pHash value. To return bytes again, the code allocates a
// buffer *b*, creates a Writer to write to *b*. The img pHash is written
// as bytes to *b* and flushed to be returned
func pHashImgBytes(img []byte) (*goimagehash.ImageHash, error) {
	imgDecode, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		return nil, err
	}
	imgHash, _ := goimagehash.PerceptionHash(imgDecode)
	return imgHash, nil
}

// Function that compares all the phashes in a slice. It adds all below a specified distance
// threshold to a duplicate list that is passed.
func pHashCompare(pHashList []store_phash) []dup_img {
	duplicateList := []dup_img{}
	for i := 0; i < len(pHashList); i++ {
		for j := i + 1; j < len(pHashList); j++ {
			distance, _ := pHashList[i].imgPHash.Distance(pHashList[j].imgPHash)
			if distance <= 2 {
				img := dup_img{
					imgPath1: pHashList[i].filename,
					imgPath2: pHashList[j].filename,
				}
				duplicateList = append(duplicateList, img)
			}
		}
	}
	return duplicateList
}

// Function to store an images pHash value and filename in a slice.
// The slice is made up of store_phash(struct) elements.
func storePHashes(imgPHash *goimagehash.ImageHash, fileName string) *store_phash {
	return &store_phash{
		imgPHash: imgPHash,
		filename: fileName,
	}
	// pHashList = append(pHashList, img)
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

			if !isImage(path) {
				return nil
			}

			imgBytes, err := readImgBytes(path)
			if err != nil {
				return err
			}

			hash := hashImgBytes(imgBytes)
			val := storeImgHashes(hash, path, hashMap)

			if val != nil {
				duplicateImgs = append(duplicateImgs, *val)
			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
		}

		prettyPrint(duplicateImgs)
	case 2: // only phash
		pHashList := []store_phash{}
		err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			if !isImage(path) {
				return nil
			}

			imgBytes, err := readImgBytes(path)
			if err != nil {
				return err
			}

			pHash, err := pHashImgBytes(imgBytes)
			if err != nil {
				fmt.Printf("Warning: Skipping %s - %v\n", path, err)
				return nil
			}
			img := storePHashes(pHash, path)
			pHashList = append(pHashList, *img)

			return nil
		})

		if err != nil {
			fmt.Println(err)
		}

		duplicateImgs := pHashCompare(pHashList)
		prettyPrint(duplicateImgs)
	case 3: // phash and has
		fmt.Println("TODO - hash & phash")
	default: // default is hash
		var hashMap = make(map[string]string)
		err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			if !isImage(path) {
				return nil
			}

			imgBytes, err := readImgBytes(path)
			if err != nil {
				return err
			}
			hash := hashImgBytes(imgBytes)
			val := storeImgHashes(hash, path, hashMap)

			if val != nil {
				duplicateImgs = append(duplicateImgs, *val)
			}

			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
		prettyPrint(duplicateImgs)
	}
}

// Pretty prints the contents of a slice containing
// dup_img structs
func prettyPrint(input []dup_img) {
	for _, item := range input {
		fmt.Println("Duplicate images found at:", item.imgPath1, "and", item.imgPath2)
	}
}

func isImage(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}

}
