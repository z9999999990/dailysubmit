package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func getfiles(path string, zipwriter *zip.Writer) {
	files, _ := os.ReadDir(path)
	for _, file := range files {
		filepath := filepath.Join(path, file.Name())

		if file.IsDir() {
			getfiles(filepath, zipwriter)
		} else {
			f, err := os.Open(filepath)
			if err != nil {
				fmt.Printf("openfile err: %s", err)
			}
			defer f.Close()

			w, err := zipwriter.Create(filepath)
			if err != nil {
				fmt.Printf("create error: %s", err)
			}

			if _, err := io.Copy(w, f); err != nil {
				fmt.Printf("close error:%s", err)
			}
		}
	}

}

func main() {

	args := os.Args[1:]

	zipfile, err := os.Create(args[0])
	if err != nil {
		fmt.Printf("创建错误：%s", err)
	}
	defer zipfile.Close()

	zipwriter := zip.NewWriter(zipfile)
	defer zipwriter.Close()

	path := args[1]
	getfiles(path, zipwriter)
}
