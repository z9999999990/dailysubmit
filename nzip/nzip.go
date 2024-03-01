package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func unzipfiles(zipfile *zip.ReadCloser, path string) error {

	for _, file := range zipfile.File {
		fpath := filepath.Join(path, file.Name)

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		} else {

			w, err := os.Create(fpath)
			if err != nil {
				fmt.Printf("create error: %s", err)
				return err
			}
			defer w.Close()

			f, err := file.Open()
			if err != nil {
				fmt.Printf("openfile err: %s", err)
				return err
			}
			defer f.Close()

			if _, err := io.Copy(w, f); err != nil {
				fmt.Printf("copy error:%s", err)
				return err
			}
		}
	}

	return nil
}

func main() {

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: program <zipfile> <outputfile>")
		return
	}

	zipfile, err := zip.OpenReader(args[0])
	if err != nil {
		fmt.Printf("读取错误： %s", err)
		return
	}
	defer zipfile.Close()

	path := args[1]
	if err := unzipfiles(zipfile, path); err != nil {
		fmt.Printf("unzip faild: %s\n", err)
		return
	}

	fmt.Println("unzip success!")
}
