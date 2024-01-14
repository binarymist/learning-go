package dupe

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var errIsNotDir = errors.New("path is not a directory")

type pathNotAbs struct {
	path string
}

func (p *pathNotAbs) Error() string {
	return fmt.Sprintf("%s is not an absolute path", p.path)
}

// Find will read each file and check if there are any duplicates for the
// specified directory if it exists. It starts by doing cheap operations to find
// a match like matching file sizes, then makes more in depth operation, by
// comparing the first few bytes of the files, if any matches are found a full
// checksum check is made.
func Find(dir string) (map[string][]string, error) {
	err := validDir(dir)
	if err != nil {
		return nil, err
	}

	dir, err = filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to find abolsute path %s: %v", dir, err)
	}

	duplicateBySize, err := findDuplicateByFileSize(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to find duplicates by size: %v", err)
	}

	inputCh := generateSizePipeline(duplicateBySize)

	duplicateByBytes1 := findDuplicateByBytes(dir, inputCh)
	duplicateByBytes2 := findDuplicateByBytes(dir, inputCh)
	duplicateByBytes3 := findDuplicateByBytes(dir, inputCh)

	duplicateByFirstBytes := merge(duplicateByBytes1, duplicateByBytes2, duplicateByBytes3)

	duplicateByChecksum1 := findDuplicateByChecksum(dir, duplicateByFirstBytes)
	duplicateByChecksum2 := findDuplicateByChecksum(dir, duplicateByFirstBytes)
	duplicateByChecksum3 := findDuplicateByChecksum(dir, duplicateByFirstBytes)

	duplicatesCh := merge(duplicateByChecksum1, duplicateByChecksum2, duplicateByChecksum3)

	duplicates := make(map[string][]string)

	for bucket := range duplicatesCh {
		for key, files := range bucket {
			duplicates[key] = files
		}
	}

	return duplicates, nil
}

// validDir will check if the path given exists, and is a directory.
func validDir(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !f.IsDir() {
		return errIsNotDir
	}

	return nil
}

func generateSizePipeline(duplicateBySize map[int64][]string) <-chan []string {
	out := make(chan []string)
	go func() {
		for _, bucket := range duplicateBySize {
			out <- bucket
		}

		close(out)
	}()

	return out
}

func merge(cs ...<-chan map[string][]string) <-chan map[string][]string {
	wg := sync.WaitGroup{}
	out := make(chan map[string][]string)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan map[string][]string) {
		for bucket := range c {
			out <- bucket
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
