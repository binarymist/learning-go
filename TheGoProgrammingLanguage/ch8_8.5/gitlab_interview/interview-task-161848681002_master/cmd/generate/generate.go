package main

import (
	"flag"
	"fmt"
	"log"

	"gitlab.com/gl-technical-interviews/backend/go/internal/dupe/fileutil"
)

const (
	defaultRootDir    = "."
	defaultNumOfFiles = 100
	fileSize          = 1024
)

func main() {
	rootDir := flag.String("rootDir", defaultRootDir, "The root directory where dupe will start checking for duplicate files.")
	numOfFiles := flag.Int("numOfFiles", defaultNumOfFiles, "The number of files that you want to generate.")

	flag.Parse()

	createFiles(*rootDir, *numOfFiles)
}

func createFiles(dir string, numOfFiles int) {
	numOfDuplicateFiles := numOfFiles / 8
	numOfPartialMatchFiles := numOfFiles / 8
	numOfUniqueFiles := numOfFiles - ((numOfPartialMatchFiles * 2) + (numOfDuplicateFiles * 2))

	log.Printf("Number of duplicate files: %d", numOfDuplicateFiles*2)
	log.Printf("Numer of partial matching files: %d", numOfPartialMatchFiles*2)
	log.Printf("Number of unique files: %d", numOfUniqueFiles)

	for i := 0; i < numOfDuplicateFiles; i++ {
		fileName, err := fileutil.Create(dir, fileSize)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}

		err = fileutil.Copy(fileName, fmt.Sprintf("%s_%d", fileName, i))
		if err != nil {
			log.Fatalf("Failed to copy file: %v", err)
		}
	}

	for i := 0; i < numOfPartialMatchFiles; i++ {
		_, err := fileutil.CreatePartialMatch(dir, fileSize)
		if err != nil {
			log.Fatalf("Failed to create partial matching files: %v", err)
		}
	}

	// The rest of the files should be unique
	for i := 0; i < numOfUniqueFiles; i++ {
		_, err := fileutil.Create(dir, fileSize)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
	}
}
