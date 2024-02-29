package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func getfiles(zipfile *zip.ReadCloser, path string) {
	for _, file := range zipfile.File {
		filepath := filepath.Join(path, file.Name)

		if file.FileInfo().IsDir() {
			_ = os.MkdirAll(filepath, os.ModePerm)
			continue
		} else {
			w, err := os.Create(filepath)
			if err != nil {
				fmt.Printf("create error: %s", err)
				continue
			}
			defer w.Close()

			f, err := zipfile.Open(file.Name)
			if err != nil {
				fmt.Printf("openfile err: %s", err)
				continue
			}
			defer f.Close()

			if _, err := io.Copy(w, f); err != nil {
				fmt.Printf("copy error:%s", err)
			}
		}

	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: program <zipfile> <outputpath>")
		return
	}

	zipfile, err := zip.OpenReader(args[0])
	if err != nil {
		fmt.Printf("读取错误： %s", err)
		return
	}
	defer zipfile.Close()

	path := args[1]

	getfiles(zipfile, path)
}
