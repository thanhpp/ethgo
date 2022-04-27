package explorerutil

import (
	"errors"
	"io/ioutil"
	"os"
	"time"
)

func LatestModFile(dir string) (os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("empty directory")
	}

	var (
		modTime time.Time
		f       os.FileInfo
	)

	for i := range files {
		if files[i].Mode().IsRegular() && files[i].ModTime().After(modTime) {
			modTime = files[i].ModTime()
			f = files[i]
		}
	}

	if f == nil {
		return nil, os.ErrNotExist
	}

	return f, nil
}
