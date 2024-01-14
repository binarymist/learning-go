package dupe

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"

	"gitlab.com/gl-technical-interviews/backend/go/internal/dupe/fileutil"
)

func Test_validDir(t *testing.T) {
	// Create tmp file.
	tmpFile, err := ioutil.TempFile("", t.Name())
	if err != nil {
		t.FailNow()
	}
	err = tmpFile.Close()
	if err != nil {
		t.FailNow()
	}

	// Create tmp dir.
	tmpDir, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.FailNow()
	}

	defer func() {
		err := os.Remove(tmpFile.Name())
		if err != nil {
			t.Logf("Failed to remove %s: %v", tmpFile.Name(), err)
		}
		err = os.RemoveAll(tmpDir)
		if err != nil {
			t.Logf("Failed to remove %s: %v", tmpDir, err)
		}
	}()

	test := fmt.Sprintf("./path/does/not/exist/%d", time.Now().Nanosecond())

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "path does not exist",
			path:    test,
			wantErr: true,
		},
		{
			name:    "path is file",
			path:    tmpFile.Name(),
			wantErr: true,
		},
		{
			name:    "directory exists",
			path:    tmpDir,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := validDir(tt.path)

			if tt.wantErr && gotErr == nil {
				t.Fail()
			}
		})
	}
}

func TestFind(t *testing.T) {
	wantDuplicateFiles, checksum, tmpDir, cleanUpFn := setupTestFind(t)
	defer cleanUpFn()

	matches, err := Find(tmpDir)
	if err != nil {
		t.Fatalf("Failed to find duplicates: %v", err)
	}

	if len(matches) != 1 {
		t.Errorf("Expected to find 1 match got: %d", len(matches))
	}

	if !reflect.DeepEqual(wantDuplicateFiles, matches[checksum]) {
		t.Errorf("Got unwated matches want:  %#+v got: %#+v", wantDuplicateFiles, matches[checksum])
	}
}

func setupTestFind(t *testing.T) ([]string, string, string, func()) {
	fileSize := 1024

	tmpDir, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.Fatal("Failed to create tmp dir for testdata")
	}

	baseFile, err := fileutil.Create(tmpDir, fileSize)
	if err != nil {
		t.Fatalf("Failed to create file for test: %v", err)
	}

	kbFileChecksum, err := getChecksum(baseFile)
	if err != nil {
		t.Fatalf("Failed to calculate checksum of file to validate test: %v", err)
	}

	// Create duplicate of base file to get a match.
	kbFileCp := strconv.Itoa(time.Now().Nanosecond())
	err = fileutil.Copy(baseFile, filepath.Join(tmpDir, kbFileCp))
	if err != nil {
		t.Fatalf("Failed to copy tmp file: %v", err)
	}

	// Create two files that are identical but are a full match.
	_, err = fileutil.CreatePartialMatch(tmpDir, fileSize)
	if err != nil {
		t.Fatalf("Failed to create partial matches")
	}

	wantDuplicateFiles := []string{filepath.Base(baseFile), filepath.Base(kbFileCp)}

	// Sort them so that we can use DeepEqual on the slice.
	sort.Strings(wantDuplicateFiles)

	return wantDuplicateFiles,
		kbFileChecksum,
		tmpDir,
		func() {
			err = os.RemoveAll(tmpDir)
			if err != nil {
				t.Logf("Failed to remove %s: %v", tmpDir, err)
			}
		}
}
