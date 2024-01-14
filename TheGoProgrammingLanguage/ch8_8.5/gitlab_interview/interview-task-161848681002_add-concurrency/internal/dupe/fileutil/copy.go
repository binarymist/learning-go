package fileutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Copy will copy the file from src to dst, the paths have to be absolute to
// ensure consistent behavior.
func Copy(src string, dst string) error {
	if !filepath.IsAbs(src) || !filepath.IsAbs(dst) {
		return fmt.Errorf("cannot copy src to dst paths not abosulte src: %s dst: %s", src, dst)
	}

	srcStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	if !srcStat.Mode().IsRegular() {
		return fmt.Errorf("failed to copy file %s not a regular file", src)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open file to copy: %v", err)
	}
	defer func() {
		_ = srcFile.Close()
	}()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create file to copy to for %s:  %v", src, err)
	}
	defer func() {
		_ = dstFile.Close()
	}()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file src: %s dst: %s err %v", src, dstFile.Name(), err)
	}

	return nil
}
