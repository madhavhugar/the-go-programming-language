package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func WalkDirOne(path string, filesize chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, file := range directoryEntriesOne(path) {
		if file.IsDir() {
			fullpath := filepath.Join(path, file.Name())
			wg.Add(1)
			WalkDirOne(fullpath, filesize, wg)
		} else {
			filesize <- file.Size()
		}
	}
}

func directoryEntriesOne(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return files
}

func printDiskUsage(path string, numFiles, size int64) {
	fmt.Println("Directory:", path, "| Total size:", size, "| Total files:", numFiles)
}

var (
	root    = flag.String("root", ".", "anaylize a given path for disk usage")
	verbose = flag.Bool("v", false, "verbose mode")
)

func init() {
	flag.Parse()
}

func main() {
	roots := []string{*root}
	filesize := make(chan int64)

	go func() {
		var wg sync.WaitGroup
		for _, r := range roots {
			wg.Add(1)
			go WalkDirOne(r, filesize, &wg)
		}
		time.Sleep(time.Second)
		wg.Wait()
		close(filesize)
	}()

	var nfiles, nbytes int64
	if *verbose {
		t := time.NewTicker(time.Second * 1)
	loop:
		for {
			select {
			case <-t.C:
				printDiskUsage(*root, nfiles, nbytes)
			case size, ok := <-filesize:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			}
		}
		t.Stop()
	} else {
		for size := range filesize {
			nfiles++
			nbytes += size
		}
	}

	printDiskUsage(*root, nfiles, nbytes)
}
