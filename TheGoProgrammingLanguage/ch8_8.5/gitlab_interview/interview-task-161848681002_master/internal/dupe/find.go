package dupe

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

	duplicateByBytes, err := findDuplicateByBytes(dir, duplicateBySize)
	if err != nil {
		return nil, fmt.Errorf("failed to find duplicates by bytes: %v", err)
	}

	duplicateFiles, err := findDuplicateByChecksum(dir, duplicateByBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to find duplicates by checksum: %v", err)
	}

	return duplicateFiles, nil
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
