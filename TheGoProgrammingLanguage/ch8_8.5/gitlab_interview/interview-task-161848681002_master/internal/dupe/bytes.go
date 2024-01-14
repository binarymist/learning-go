package dupe

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// byteRange is the amount of bytes we read from the file.
const byteRange = 10

func findDuplicateByBytes(basePath string, duplicateBySize map[int64][]string) (map[string][]string, error) {
	if !filepath.IsAbs(basePath) {
		return nil, &pathNotAbs{path: basePath}
	}

	filesByBytes := make(map[string][]string)

	for size, paths := range duplicateBySize {
		for _, path := range paths {
			firstBytes, err := getFirstBytes(filepath.Join(basePath, path))
			if err != nil {
				return nil, err
			}

			// Create a unique key with size and first bytes, so we don't match
			// two files that have the same firstBytes but different size.
			key := fmt.Sprintf("%d%s", size, firstBytes)

			files, ok := filesByBytes[key]

			if !ok {
				filesByBytes[key] = []string{path}
				continue
			}

			filesByBytes[key] = append(files, path)
		}
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
