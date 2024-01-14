package dupe

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// byteRange is the amount of bytes we read from the file.
const byteRange = 10

func findDuplicateByBytes(basePath string, duplicateBySize <-chan []string, cancel *cancellation) <-chan map[string][]string {
	out := make(chan map[string][]string)
	var wg sync.WaitGroup
	for fileNames := range duplicateBySize {
		select {
		case <-cancel.done:
			// Channel closed due to error. No point in continuing to process.
		default:
			wg.Add(1)
			go func(fileNames []string) {
				defer wg.Done()
				matches, err := checkFirstBytes(basePath, fileNames)
				if err != nil {
					cancel.closeDone(fmt.Errorf("failed to find duplicates by bytes: %v", err))
					return
				}
				out <- matches
			}(fileNames) // Go 1.22 was supposed to fix the for loop iteration variable issue, but I don't think it went in yet.
		}
	}
	go func() {
		wg.Wait()
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
