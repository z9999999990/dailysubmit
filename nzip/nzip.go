package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func nmkfile(filepath string, zipwriter *zip.Writer) error {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("openfile err: %s", err)
		return err
	}
	defer f.Close()

	w, err := zipwriter.Create(filepath)
	if err != nil {
		fmt.Printf("create error: %s", err)
		return err
	}

	if _, err := io.Copy(w, f); err != nil {
		fmt.Printf("close error:%s", err)
		return err
	}
	return nil
}

func getfiles(path string, zipwriter *zip.Writer) error {

	files, _ := os.ReadDir(path)
	for _, file := range files {
		filepath := filepath.Join(path, file.Name())

		if file.IsDir() {
			getfiles(filepath, zipwriter)
		} else {
			if err := nmkfile(filepath, zipwriter); err != nil {
				return err
			}
		}
	}
	return nil
}

func nzip(args []string) {

	zipfile, err := os.Create(args[1])
	if err != nil {
		fmt.Printf("创建错误：%s", err)
	}
	defer zipfile.Close()

	zipwriter := zip.NewWriter(zipfile)
	defer zipwriter.Close()

	path := args[2]
	if err := getfiles(path, zipwriter); err != nil {
		fmt.Printf("nzip faild: %s", err)
		return
	}

	fmt.Println("nzip success!")
}

func umkfile(fpath string, file *zip.File) error {
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
	return nil
}

func unzipfiles(zipfile *zip.ReadCloser, path string) error {

	for _, file := range zipfile.File {
		fpath := filepath.Join(path, file.Name)

		if file.FileInfo().IsDir() {
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		if err := umkfile(fpath, file); err != nil {
			return err
		}
	}

	return nil
}

func unzip(args []string) {

	zipfile, err := zip.OpenReader(args[1])
	if err != nil {
		fmt.Printf("读取错误： %s", err)
		return
	}
	defer zipfile.Close()

	path := args[2]
	if err := unzipfiles(zipfile, path); err != nil {
		fmt.Printf("unzip faild: %s\n", err)
		return
	}

	fmt.Println("unzip success!")
}

func main() {

	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Usage: program <command> <zipfile> <docfile>")
		return
	}

	if args[0] == "-n" {
		nzip(args)
	}

	if args[0] == "-u" {
		unzip(args)
	}

}
