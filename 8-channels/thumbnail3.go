package channels

import (
	"fmt"
	"math/rand"
	"time"
)

// thumbnail with error handler:
// (i) where go routine leak happens

func thumbnailConverter3(file string) (string, error) {
	delay := time.Second*1 + time.Duration(rand.Float64()*1)
	time.Sleep(delay)
	if file == "error" {
		return "", fmt.Errorf("invalid file")
	}
	return fmt.Sprintf("%s_thumbnail", file), nil
}

func parallelGoRoutinesWithUnBufferedChannelWithGoRoutineLeak(
	files []string,
) ([]string, error) {
	type thumbnailResponse struct {
		filename string
		err      error
	}
	// declare a unbuffered channel
	thumbnails := make(chan thumbnailResponse)

	for _, file := range files {
		go func(f string) {
			filename, err := thumbnailConverter3(f)
			// sending here is blocking operation
			thumbnails <- thumbnailResponse{filename, err}
		}(file)
	}

	var filenames []string
	for range files {
		response := <-thumbnails
		if response.err != nil {
			// on error will cause goroutine leak
			// goroutines will block on syncronous receive above
			return nil, response.err
		}
		filenames = append(filenames, response.filename)
	}
	return filenames, nil
}
