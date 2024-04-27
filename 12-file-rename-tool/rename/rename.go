package rename

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Execute() error {
	baseNameMap, err := getBaseNameCountMap()
	if err != nil {
		return err
	}

	for name, total := range baseNameMap {
		currFile := 1
		pattern := name + "_???.*"
		err := filepath.WalkDir(
			"tmp",
			func(path string, info fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				fileName := filepath.Base(path)
				match, err := filepath.Match(pattern, fileName)
				if err != nil {
					return err
				}
				if match {
					ext := strings.Split(info.Name(), ".")[1]
					newName := fmt.Sprintf(
						"%v (%v of %v).%v",
						name,
						currFile,
						total,
						ext,
					)
					newPath := filepath.Join(filepath.Dir(path), newName)
					fmt.Println(newPath)
					if err := os.Rename(path, newPath); err != nil {
						return err
					}
					currFile++
				}
				return nil
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func getBaseNameCountMap() (map[string]int, error) {
	pattern := "*_???.*"

	baseNameMap := make(map[string]int)

	err := filepath.WalkDir(
		"tmp",
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			fileName := filepath.Base(path)
			match, err := filepath.Match(pattern, fileName)
			if err != nil {
				return err
			}
			if match {
				baseName := strings.Split(fileName, "_")[0]
				if currCount, ok := baseNameMap[baseName]; ok {
					baseNameMap[baseName] = currCount + 1
					return nil
				}
				baseNameMap[baseName] = 1
			}
			return nil
		},
	)
	return baseNameMap, err
}
