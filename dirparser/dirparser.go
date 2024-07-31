package dirparser

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type DirectoriesSizes struct {
	Directories []Directory `json:"directories"`
}

type Directory struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func GetDirSize(directories []string, isRecursive, isHumanReadable bool) (resp DirectoriesSizes, err error) {
	resp.Directories = make([]Directory, 0, len(directories))
	for _, path := range directories {
		var dirStats Directory
		fileInfo, err := os.Lstat(path)
		if err != nil {
			return
		}
		if !fileInfo.IsDir() {
			err = fmt.Errorf("%s is not a directory", fileInfo.Name())
		}

		dirStats.Name = fileInfo.Name()
		size, err := dirSize(path)
		if err != nil {
			err = fmt.Errorf("wile getting dir size for ")
			return
		}
		dirStats.Size = size
		resp.Directories = append(resp.Directories, dirStats)

	}

	return resp, err
}

func dirSize(path string) (int64, error) {
	var size int64
	var mu sync.Mutex

	// Function to calculate size for a given path
	var calculateSize func(string) error
	calculateSize = func(p string) error {
		fileInfo, err := os.Lstat(p)
		if err != nil {
			return err
		}

		// Skip symbolic links to avoid counting them multiple times
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		if fileInfo.IsDir() {
			entries, err := os.ReadDir(p)
			if err != nil {
				return err
			}
			for _, entry := range entries {
				if err := calculateSize(filepath.Join(p, entry.Name())); err != nil {
					return err
				}
			}
		} else {
			mu.Lock()
			size += fileInfo.Size()
			mu.Unlock()
		}
		return nil
	}

	// Start calculation from the root path
	if err := calculateSize(path); err != nil {
		return 0, err
	}

	return size, nil
}
