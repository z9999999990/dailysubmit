package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func getfiles(path string, w io.Writer) {
	files, _ := os.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			getfiles(path + "/" + file.Name())
		} else {
			f, err := os.Open(file.Name())
			if err != nil {
				fmt.Printf("openfile err: %s", err)
			}
			defer f.Close()

			if _, err := io.Copy(w, f); err != nil {
				fmt.Printf("close error:%s", err)
			}
		}
	}
}

func zipab() {}

func main() {

	args := os.Args[1:]

	zipfile, err := os.Create(args[0])
	if err != nil {
		fmt.Printf("创建错误：%s", err)
	}
	defer zipfile.Close()

	zipwriter := zip.NewWriter(zipfile)
	w, err := zipwriter.Create("ziptest.txt")
	if err != nil {
		fmt.Printf("create error: %s", err)
	}

	getfiles(".", w)
	zipwriter.Close()
}
