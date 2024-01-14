package dupe

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func findDuplicateByChecksum(basePath string, duplicateByFirstByte <-chan map[string][]string, cancel *cancellation) <-chan map[string][]string {
	out := make(chan map[string][]string)
	var wg sync.WaitGroup
	for fileNames := range duplicateByFirstByte {
		select {
		case <-cancel.done:
			// Channel closed due to error. No point in continuing to process.
		default:
			wg.Add(1)
			go func(fileNames map[string][]string) {
				defer wg.Done()
				matches, err := checkChecksum(basePath, fileNames)
				if err != nil {
					cancel.closeDone(fmt.Errorf("failed to find duplicates by checksum: %v", err))
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

func checkChecksum(basePath string, duplicateByFirstByte map[string][]string) (map[string][]string, error) {
	if !filepath.IsAbs(basePath) {
		return nil, &pathNotAbs{path: basePath}
	}

	duplicateByHash := make(map[string][]string)

	for _, paths := range duplicateByFirstByte {
		for _, path := range paths {
			sum, err := getChecksum(filepath.Join(basePath, path))
			if err != nil {
				return nil, err
			}

			duplicateFiles, ok := duplicateByHash[sum]

			if !ok {
				duplicateByHash[sum] = []string{path}
			}

			duplicateByHash[sum] = append(duplicateFiles, path)
		}
	}

	// Remove any sizes that only have 1 files attached to them.
	for size, paths := range duplicateByHash {
		if len(paths) != 1 {
			continue
		}

		delete(duplicateByHash, size)
	}

	return duplicateByHash, nil
}

func getChecksum(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	sha1Hash := sha1.New()
	if _, err := io.Copy(sha1Hash, file); err != nil {
		return "", err
	}
	sum := sha1Hash.Sum(nil)

	return string(sum), nil
}
