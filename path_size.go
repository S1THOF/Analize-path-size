package psize

import (
	"fmt"
	"log"
	"os"
)

func GetSize(path string) {
	pathStat, err := os.Lstat(path)
	var result int64
	if err != nil {
		log.Fatal("Wrong the path")
	}

	if pathStat.IsDir() {
		dir, err := os.ReadDir(path)
		if err != nil {
			log.Fatal("Error reading dir")
		}

		fmt.Println("Path is dir", dir)

		for _, file := range dir {
			if !file.IsDir() {
				fileInfo, err := os.DirEntry.Info(file)
				if err != nil {
					log.Fatal("Can't read file info")
				}
				if fileInfo.Mode().IsRegular() {
					result += fileInfo.Size()
				}
			}
		}
	} else {
		result += pathStat.Size()
	}

	fmt.Printf("%vB	%s", result, path)
}
