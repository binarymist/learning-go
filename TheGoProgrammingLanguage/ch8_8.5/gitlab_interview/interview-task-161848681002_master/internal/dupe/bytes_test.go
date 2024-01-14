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

	duplicateByBytes, err := findDuplicateByBytes(tmpDir, duplicateBySize)
	if err != nil {
		t.Fatalf("Got unexpected error when find duplicates by bytes: %v", err)
	}

	// We can assert that the exact map, but we get little benefit of doing so,
	// it's quicker to just test the length of the map and each key, and results
	// into the same assertion.
	if len(duplicateByBytes) != 2 {
		t.Errorf("Got unexpected length of on map wanted 2 got: %d", len(duplicateByBytes))
	}

	for _, duplicate := range duplicateByBytes {
		if len(duplicate) != 2 {
			t.Errorf("Got unexpected length of on slice wanted 2 got: %d", len(duplicateByBytes))
		}
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
