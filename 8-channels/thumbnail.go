package channels

import (
	"fmt"
	"time"
)

func thumbnailConverter(filepath string) (thumbnailPath string) {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s_thumbnail", filepath)
}

func howNotToWriteParallelGoRoutines(files []string) {
	for _, file := range files {
		// Will exit quickly
		// does not wait for the goroutines to finish
		go thumbnailConverter(file)
	}
}

func thumbnailConverter1(filepath string) (thumbnailPath string, err error) {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s_thumbnail", filepath), nil
}

// Parallel go routines without error handling
func parallelGoRoutinesWithUnbufferedChannel(files []string) []string {
	thumbnails := make(chan string)
	var pathnames []string

	for _, file := range files {
		go func(f string) {
			thumbnail, _ := thumbnailConverter1(f)
			thumbnails <- thumbnail
		}(file)
	}

	for range files {
		pathnames = append(pathnames, <-thumbnails)
	}
	return pathnames
}
