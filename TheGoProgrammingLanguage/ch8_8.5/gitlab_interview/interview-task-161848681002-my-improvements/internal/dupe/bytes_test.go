package dupe

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"gitlab.com/gl-technical-interviews/backend/go/internal/dupe/fileutil"
)

func Test_findDuplicateByBytes(t *testing.T) {
	tmpDir, duplicateBySize, cleanFn := setupTestFindDuplicateByBytes(t)
	defer func() {
		err := cleanFn()
		if err != nil {
			t.Logf("Failed to clean up test %s: %v", t.Name(), err)
		}
	}()

	cancelFindDuplicate := cancellation{
		done: make(chan struct{}),
	}

	duplicateBySizeCh := make(chan []string)

	go func() {
		for _, fileNames := range duplicateBySize {
			duplicateBySizeCh <- fileNames
		}
		close(duplicateBySizeCh)
	}()

	duplicateByBytesCh := findDuplicateByBytes(tmpDir, duplicateBySizeCh, &cancelFindDuplicate)

	duplicateByBytes, ok := <-duplicateByBytesCh

	if cancelFindDuplicate.reason != nil {
		t.Fatalf("Got unexpected error when find duplicates by bytes: %v", cancelFindDuplicate.reason)
	}

	if !ok {
		t.Errorf("Channel is closed, should be open for first message")
	}
	if len(duplicateByBytes) != 2 {
		t.Errorf("Got unexpected length of duplicateByBytes map, wanted 2 got: %d", len(duplicateByBytes))
	}
	for _, duplicate := range duplicateByBytes {
		if len(duplicate) != 2 {
			t.Errorf("Got unexpected length of duplicate slice, wanted 2 got: %d", len(duplicate))
		}
	}

	message2, ok := <-duplicateByBytesCh
	if message2 != nil || ok {
		t.Errorf("Channel is not closed after receiving 2nd message. 2nd message is %#v", message2)
	}

}

func setupTestFindDuplicateByBytes(t *testing.T) (string, map[int64][]string, func() error) {
	tmpDir, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.Fatal("Failed to create tmp dir for testdata")
	}

	size := int64(1024)

	partialMatches1, err := fileutil.CreatePartialMatch(tmpDir, int(size))
	if err != nil {
		t.Fatalf("Failed to create partial match files: %v", err)
	}

	partialMatches2, err := fileutil.CreatePartialMatch(tmpDir, int(size))
	if err != nil {
		t.Fatalf("Failed to create partial match files: %v", err)
	}

	bucket := make(map[int64][]string, 1)

	bucket[size] = append(bucket[size], partialMatches1...)
	bucket[size] = append(bucket[size], partialMatches2...)

	for i, path := range bucket[size] {
		bucket[size][i] = filepath.Base(path)
	}

	sort.Strings(bucket[size])

	return tmpDir,
		bucket,
		func() error {
			return os.RemoveAll(tmpDir)
		}
}
