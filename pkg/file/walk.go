package file

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func GetGitRepos(rootPath string) (folders []string, err error) {

	startStat, err := os.Stat(rootPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat root directory: %w", err)
	}
	startDev := startStat.Sys().(*syscall.Stat_t).Dev

	err = filepath.WalkDir(rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			stat, err := info.Info()
			if err != nil {
				return fmt.Errorf("failed to stat directory %q: %w", path, err)
			}

			currDev := stat.Sys().(*syscall.Stat_t).Dev

			// If we've crossed to a different device, skip this directory
			if currDev != startDev {
				return filepath.SkipDir
			}
			if info.Name() == ".git" {
				gitPath, _ := strings.CutSuffix(path, ".git")
				folders = append(folders, gitPath)

				return filepath.SkipDir
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("-> variable - err = %v is of type %T \n", err, err)
	}
	return folders, nil
}
