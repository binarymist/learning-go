package dupe

import (
	"os"
	"path/filepath"
)

// findDuplicateByFileSize will check the whole directory, and group each file
// into a bucket of the same size. This will not look into sub directories, only
// at files in the root level.
func findDuplicateByFileSize(dir string) (map[int64][]string, error) {
	if !filepath.IsAbs(dir) {
		return nil, &pathNotAbs{path: dir}
	}

	filesBySize := make(map[int64][]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// Ignore if it's a directory.
		if info.IsDir() {
			return nil
		}

		size, ok := filesBySize[info.Size()]

		// Check if we have size already tracked.
		if !ok {
			filesBySize[info.Size()] = []string{info.Name()}
			return nil
		}

		// Size found append to slice.
		filesBySize[info.Size()] = append(size, info.Name())

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Remove any sizes that only have 1 files attached to them.
	for size, paths := range filesBySize {
		if len(paths) != 1 {
			continue
		}

		delete(filesBySize, size)
	}

	return filesBySize, nil
}
