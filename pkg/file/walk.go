package file

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func GetGitRepos(rootPath string) (folders []string) {
	err := filepath.WalkDir(rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
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
	return folders
}
