package channels

import (
	"fmt"
	"math/rand"
	"time"
)

// thumbnail with error handler:
// (ii) where go routine leak does not happen

func thumbnailConverter2(file string) (string, error) {
	delay := time.Second*1 + time.Duration(rand.Float64()*1)
	time.Sleep(delay)
	if file == "error" {
		return "", fmt.Errorf("invalid file")
	}
	return fmt.Sprintf("%s_thumbnail", file), nil
}

func parallelGoRoutinesWithBufferedChannel(files []string) ([]string, error) {
	type thumbnailResponse struct {
		filename string
		err      error
	}
	// declare a buffered channel
	thumbnails := make(chan thumbnailResponse, len(files))

	for _, file := range files {
		go func(f string) {
			filename, err := thumbnailConverter2(f)
			// non-blocking operation, because of buffered channel
			thumbnails <- thumbnailResponse{filename, err}
		}(file)
	}

	var pathnames []string
	for range files {
		res := <-thumbnails
		if res.err != nil {
			// no goroutine leaks even if we return mid-way,
			// we are using buffered channels,
			// no goroutines are left hanging even after main exits
			// since we do not have a blocking operation above
			return nil, res.err
		}
		pathnames = append(pathnames, res.filename)
	}
	return pathnames, nil
}
