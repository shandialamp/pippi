package storage

import (
	"os"
)

func MakeDirectory(directory string) error {
	return os.MkdirAll(directory, 0755)
}

func DeleteDirectory(directory string) error {
	return os.RemoveAll(directory)
}

func Directorys(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	directorys := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			directorys = append(directorys, file.Name())
		}
	}
	return directorys, nil
}
