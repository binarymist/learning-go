package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.com/gl-technical-interviews/backend/go/internal/dupe"
)

const defaultRootDir = "."

func main() {
	rootDir := flag.String("rootDir", defaultRootDir, "The root directory where dupe will start checking for duplicate files.")

	flag.Parse()

	startTime := time.Now()
	matches, err := dupe.Find(*rootDir)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
	if err != nil {
		log.Printf("Failed to find duplicates files: %v", err)
		os.Exit(1)
	}

	if len(matches) == 0 {
		log.Print("No duplicates found")
		os.Exit(0)
	}

	logger := log.New(os.Stdout, "", 0)

	logger.Printf("Found the following duplicates for %s:", *rootDir)
	for checksum, duplicates := range matches {
		logger.Printf("Checksum : %s", hex.EncodeToString([]byte(checksum)))
		for _, fileName := range duplicates {
			logger.Printf("\t file: %s", fileName)
		}
	}
}
