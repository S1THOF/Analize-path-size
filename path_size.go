package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	pathStat, err := os.Lstat(path)
	var sum int64
	if err != nil {
		return "", err
	}

	if pathStat.IsDir() {
		dir, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}

		for _, entry := range dir {
			if entry.IsDir() {
				continue
			}
			fileInfo, err := entry.Info()
			if err != nil {
				return "", err
			}
			if fileInfo.Mode().IsRegular() {
				sum += fileInfo.Size()
			}
		}

	} else {
		sum += pathStat.Size()
	}
	result := fmt.Sprintf("%vB	%s", sum, path)
	return result, nil
}
