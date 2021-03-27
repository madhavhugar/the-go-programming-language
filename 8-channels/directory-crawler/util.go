package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WalkDir(path string, filesize chan int64) {
	for _, file := range directoryEntries(path) {
		if file.IsDir() {
			p := filepath.Join(path, file.Name())
			WalkDir(p, filesize)
		} else {
			filesize <- file.Size()
		}
	}
}

func directoryEntries(file string) []os.FileInfo {
	files, err := ioutil.ReadDir(file)
	if err != nil {
		fmt.Println("Error reading directory: ", file)
		return nil
	}
	return files
}

func main() {
	filesize := make(chan int64)
	go func() {
		WalkDir("/home/maddy", filesize)
		close(filesize)
	}()

	var nfiles, nbytes int64
	for size := range filesize {
		nfiles++
		nbytes += size
	}
	fmt.Println("Total size:", nbytes, "with", nfiles, "files")
}
