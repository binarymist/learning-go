package dupe

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"gitlab.com/gl-technical-interviews/backend/go/internal/dupe/fileutil"
)

func Test_findDuplicateByChecksum(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.Fatal("Failed to create tmp dir for testdata")
	}
	defer func() {
		err = os.RemoveAll(tmpDir)
		if err != nil {
			t.Logf("Failed to remove %s: %v", tmpDir, err)
		}
	}()

	f, err := fileutil.Create(tmpDir, 10)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	fCp := strconv.Itoa(time.Now().Nanosecond())
	err = fileutil.Copy(f, filepath.Join(tmpDir, fCp))
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	duplicateByFirstByte := map[string][]string{
		"10": {
			filepath.Base(f),
			filepath.Base(fCp),
		},
	}

	duplicateByChecksum, err := findDuplicateByChecksum(tmpDir, duplicateByFirstByte)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if len(duplicateByChecksum) != 1 {
		t.Errorf("Got unexpected len expected 1 got: %d", len(duplicateByChecksum))
	}

	for _, duplicate := range duplicateByChecksum {
		if len(duplicate) != 2 {
			t.Errorf("Got unexpected length of on slice wanted 2 got: %d", len(duplicateByChecksum))
		}
	}
}
