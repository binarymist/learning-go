package fileutil

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Specify the number of bytes you want to be different when create partial
// matching files.
const numOfDiffBytes = 4

// Create will create a file in a specific dir, for the specific size. The
// data inside of file is completely random.
func Create(dir string, size int) (string, error) {
	tmpFile, err := ioutil.TempFile(dir, "")
	if err != nil {
		return "", fmt.Errorf("failed to create tmp file inside of %s: %v", dir, err)
	}
	defer func() {
		_ = tmpFile.Close()
	}()

	fileContent := make([]byte, size)
	_, err = rand.Read(fileContent)
	if err != nil {
		return "", fmt.Errorf("failed to create random data for file: %v", err)
	}

	_, err = tmpFile.Write(fileContent)
	if err != nil {
		return "", fmt.Errorf("failed to write random data to file %s: %v", tmpFile.Name(), err)
	}

	return tmpFile.Name(), nil
}

// CreatePartialMatch will create two files that have the same size and the
// first few bytes but the final 4 bytes of the file are different. The 4 extra
// bytes are included into the specified size.
func CreatePartialMatch(dir string, size int) ([]string, error) {
	if !filepath.IsAbs(dir) {
		return nil, fmt.Errorf("cannot append to file, path is not absolute for: %s", dir)
	}

	// Create the identical files.
	originalFile, err := Create(dir, size-numOfDiffBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to create file inside %s: %v", dir, err)
	}

	cpFile := fmt.Sprintf("%s_partial", originalFile)
	err = Copy(originalFile, cpFile)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file inside %s: %v", dir, err)
	}

	err = appendToFile(originalFile, numOfDiffBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to append random data to %s: %v", originalFile, err)
	}

	err = appendToFile(cpFile, numOfDiffBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to append random data to %s: %v", originalFile, err)
	}

	return []string{originalFile, cpFile}, nil
}

func appendToFile(path string, size int) error {
	if !filepath.IsAbs(path) {
		return fmt.Errorf("cannot append to file, path is not absolute: %s", path)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", path, err)
	}
	defer func() {
		_ = f.Close()
	}()

	appendData := make([]byte, size)
	_, err = rand.Read(appendData)
	if err != nil {
		return fmt.Errorf("failed to create random data for file: %v", err)
	}

	if _, err := f.Write(appendData); err != nil {
		return fmt.Errorf("failed to append data to file %s: %v", path, err)
	}

	return nil
}
