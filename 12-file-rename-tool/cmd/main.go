package main

import (
	"12-file-rename-tool/rename"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if err := removeTmp(); err != nil {
		return err
	}
	if err := setupTmp(); err != nil {
		return err
	}
	fmt.Println("Tmp created...")
	fmt.Println()
	fmt.Println("Reading...")
	if err := readAll(); err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Renaming...")
	if err := rename.Execute(); err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Reading...")
	if err := readAll(); err != nil {
		return err
	}
	fmt.Println()

	return nil
}

func readAll() error {
	return filepath.WalkDir(
		"tmp",
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			fileName := filepath.Base(path)
			fmt.Println(fileName)
			return nil
		},
	)
}

func removeTmp() error {
	if _, err := os.ReadDir("tmp"); os.IsNotExist(err) {
		return nil
	}

	return os.RemoveAll("tmp")
}

func setupTmp() error {
	err := os.Mkdir("tmp", 0755)
	if err != nil {
		return err
	}

	for i := 1; i < 5; i++ {
		fileName := fmt.Sprintf("birthday_00%v.txt", i)
		pathName := filepath.Join("tmp", fileName)
		_, err = os.Create(pathName)
		if err != nil {
			return err
		}
	}

	for i := 1; i < 6; i++ {
		fileName := fmt.Sprintf("christmas 2016 (%v of 100).txt", i)
		pathName := filepath.Join("tmp", fileName)
		_, err = os.Create(pathName)
		if err != nil {
			return err
		}
	}

	err = os.Mkdir(filepath.Join("tmp", "nested"), 0755)
	if err != nil {
		return err
	}

	for i := 8; i < 11; i++ {
		fileName := "n_0"
		if i < 10 {
			fileName += "0"
		}
		fileName += fmt.Sprintf("%v.txt", i)
		pathName := filepath.Join("tmp", "nested", fileName)
		_, err = os.Create(pathName)
		if err != nil {
			return err
		}
	}

	return nil
}
