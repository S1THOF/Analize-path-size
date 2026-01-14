package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func calcSize(path string, recursive, human, all bool) (int64, error) {
	fileinfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !fileinfo.IsDir() {
		if shouldInclude(fileinfo.Name(), all) && fileinfo.Mode().IsRegular() {
			return fileinfo.Size(), nil
		}
	}

	entryes, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	for _, entry := range entryes {
		if shouldInclude(entry.Name(), all) {
			info, err := entry.Info()
			if err != nil {
				return 0, nil
			}
			size := info.Size()
			total += size
		}

		fullPath := filepath.Join(path, entry.Name())

		if recursive && entry.IsDir() {
			size, err := calcSize(fullPath, recursive, human, all)
			if err != nil {
				return 0, err
			}
			total += size
		} else if entry.IsDir() {
			fileinfo, err := entry.Info()
			if err != nil {
				return 0, err
			}
			if fileinfo.Mode().IsRegular() {
				total += fileinfo.Size()
			}
		}
	}

	return total, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := calcSize(path, recursive, human, all)
	if err != nil {
		return "", err
	}

	var result string
	if human {
		result = fmt.Sprintf("%s\t%s", FormatSize(float64(size)), path)
	} else {
		result = fmt.Sprintf("%vB\t%s", size, path)
	}
	return result, nil
}

func FormatSize(size float64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0

	for size >= 1024 && unitIndex < len(units)-1 {
		size /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%.0fB", size)
	}
	return fmt.Sprintf("%.1f%s", size, units[unitIndex])
}

func shouldInclude(name string, all bool) bool {
	return all || !strings.HasPrefix(name, ".")
}
