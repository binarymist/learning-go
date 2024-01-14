package dupe

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"gitlab.com/gl-technical-interviews/backend/go/internal/dupe/fileutil"
)

func Test_findDuplicateByFileSize(t *testing.T) {
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

	sizeBucket1 := make([]string, 5)
	for i := 0; i < 5; i++ {
		f, err := fileutil.Create(tmpDir, 50)
		if err != nil {
			t.Fatalf("Failed to create tmp file: %v", err)
		}

		sizeBucket1[i] = filepath.Base(f)
	}

	sizeBucket2 := make([]string, 5)
	for i := 0; i < 5; i++ {
		f, err := fileutil.Create(tmpDir, 30)
		if err != nil {
			t.Fatalf("Failed to create tmp file: %v", err)
		}

		sizeBucket2[i] = filepath.Base(f)
	}

	_, err = fileutil.Create(tmpDir, 33)
	if err != nil {
		t.Fatalf("Failed to create tmp file: %v", err)
	}

	wantDuplicates := map[int64][]string{
		50: sizeBucket1,
		30: sizeBucket2,
	}

	gotDuplicates, err := findDuplicateByFileSize(tmpDir)
	if err != nil {
		t.Fatalf("Got unexpected error when check duplicate by file size: %v", err)
	}

	// Sort files so that we can run DeepEqual on slices.
	sort.Strings(sizeBucket1)
	sort.Strings(sizeBucket2)
	if !reflect.DeepEqual(wantDuplicates, gotDuplicates) {
		t.Errorf("Got unexpected duplicates want: %#+v got: %#+v", wantDuplicates, gotDuplicates)
	}
}
