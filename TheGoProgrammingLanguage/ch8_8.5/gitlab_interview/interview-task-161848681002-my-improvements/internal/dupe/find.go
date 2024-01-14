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

type cancellation struct {
	done   chan struct{}
	reason error
	mu     sync.Mutex
}

func (d *cancellation) closeDone(err error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	select {
	case <-d.done:
		// Channel already closed
		fmt.Println("Channel already closed")
	default:
		close(d.done)
		d.reason = err
	}
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

	cancelFindDuplicate := cancellation{
		done: make(chan struct{}),
	}

	inputCh := generateSizePipeline(duplicateBySize)

	duplicateByBytes := findDuplicateByBytes(dir, inputCh, &cancelFindDuplicate)

	duplicatesCh := findDuplicateByChecksum(dir, duplicateByBytes, &cancelFindDuplicate)

	if cancelFindDuplicate.reason != nil {
		return nil, cancelFindDuplicate.reason
	}

	duplicates := make(map[string][]string)

	// The flow of channel messages:
	// generateSizePipeline out channel ->
	// findDuplicateByBytes duplicateBySize channel, out channel ->
	// findDuplicateByChecksum duplicateByFirstByte channel, out channel ->
	// duplicatesCh
	for bucket := range duplicatesCh {
		for firstBytesKey, fileNames := range bucket {
			duplicates[firstBytesKey] = fileNames
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
