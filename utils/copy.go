package utils

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

var ErrDestNotExist error = errors.New("The folder you specified doesn't exist")

// Copy copies the src folder to another location.
// Implemented from: https://github.com/mactsouk/opensource.com/blob/master/cp1.go
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if sourceFileStat.Mode()&os.ModeSymlink != 0 {
		src, err = os.Readlink(src)
		if err != nil {
			log.Fatal(err)
		}
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func CopyDir(src string, destination string) ([]int64, error) {
	info, err := os.Stat(src)

	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(destination, info.Mode())
	if err != nil {
		log.Fatal(err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	sizeOperations := []int64{}
	for _, files := range entries {
		srcPath := filepath.Join(src, files.Name())
		destinationPath := filepath.Join(destination, files.Name())

		if files.IsDir() {
			_, err := CopyDir(srcPath, destinationPath)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		size, err := Copy(srcPath, destinationPath)
		if err != nil {
			log.Fatal(err)
		}
		sizeOperations = append(sizeOperations, size)
	}

	return sizeOperations, nil
}
