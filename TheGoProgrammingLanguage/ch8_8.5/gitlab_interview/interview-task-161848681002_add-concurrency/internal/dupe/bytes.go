package dupe

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// byteRange is the amount of bytes we read from the file.
const byteRange = 10

func findDuplicateByBytes(basePath string, duplicateBySize <-chan []string) <-chan map[string][]string {
	out := make(chan map[string][]string)

	go func() {
		for bucket := range duplicateBySize {
			matches, err := checkFirstBytes(basePath, bucket)
			if err != nil {
				log.Printf("Failed to check file first bytes: %v", err)
			}

			out <- matches
		}

		close(out)
	}()

	return out
}

func checkFirstBytes(basePath string, duplicateBySize []string) (map[string][]string, error) {
	if !filepath.IsAbs(basePath) {
		return nil, &pathNotAbs{path: basePath}
	}

	filesByBytes := make(map[string][]string)

	for _, path := range duplicateBySize {
		firstBytes, err := getFirstBytes(filepath.Join(basePath, path))
		if err != nil {
			return nil, err
		}

		files, ok := filesByBytes[firstBytes]

		if !ok {
			filesByBytes[firstBytes] = []string{path}
			continue
		}

		filesByBytes[firstBytes] = append(files, path)
	}

	for firstBytes, paths := range filesByBytes {
		if len(paths) != 1 {
			continue
		}

		delete(filesByBytes, firstBytes)
	}

	return filesByBytes, nil
}

func getFirstBytes(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	firstBytes := make([]byte, byteRange)
	if _, err := io.ReadFull(file, firstBytes); err != nil {
		return "", fmt.Errorf("failed to read %s: %v", path, err)
	}

	return string(firstBytes), nil
}
