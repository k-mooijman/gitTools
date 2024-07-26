package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func CountFiles() int {

	rootPath := "/"

	err := filepath.WalkDir(rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			if info.Name() == ".git" {
				fmt.Printf("variable - git path = %v is of type %T \n", path, path)
				return filepath.SkipDir
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("-> variable - err = %v is of type %T \n", err, err)
	}

	return 1
}
