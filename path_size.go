package code

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var result string
	var sum int64
	pathStat, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if (pathStat.IsDir() && all && strings.HasPrefix(pathStat.Name(), ".")) || (pathStat.IsDir() && !strings.HasPrefix(pathStat.Name(), ".")) {
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
			if (fileInfo.Mode().IsRegular() && all && strings.HasPrefix(pathStat.Name(), ".")) || (fileInfo.Mode().IsRegular() && !strings.HasPrefix(pathStat.Name(), ".")) {
				sum += fileInfo.Size()
			}
		}

	}

	if (pathStat.Mode().IsRegular() && all && strings.HasPrefix(pathStat.Name(), ".")) || (pathStat.Mode().IsRegular() && !strings.HasPrefix(pathStat.Name(), ".")) {
		sum += pathStat.Size()
	}

	if human {
		result = fmt.Sprintf("%s	%s", FormatSize(float64(sum)), path)
	}
	if !human {
		result = fmt.Sprintf("%vB	%s", sum, path)
	}
	return result, nil
}

func FormatSize(size float64) string {
	var res string
	format := [7]string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	var stage int

	for i := 0; size >= 1024; i++ {
		size = size / 1024
		stage++
	}

	res = fmt.Sprintf("%.1f%v", size, format[stage])

	return res
}
