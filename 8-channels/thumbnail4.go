package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// thumbnail with wg.sync
func thumbnailConverter4(filename string) (string, error) {
	delay := (time.Second * 1) + time.Duration(rand.Float64()*1)
	time.Sleep(delay)
	if filename == "error" {
		return "", fmt.Errorf("invalid file")
	}
	return fmt.Sprintf("%s_thumbnail", filename), nil
}

func parallelGoRoutinesWithWaitGroup(files []string) ([]string, error) {
	thumbnailChannel := make(chan string, len(files))
	errChan := make(chan error, len(files))
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumbnail, err := thumbnailConverter4(f)
			fmt.Println(thumbnail)
			errChan <- err
			thumbnailChannel <- thumbnail
		}(file)
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(thumbnailChannel)
	}()
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}
	var thumbnails []string
	for thumbnail := range thumbnailChannel {
		thumbnails = append(thumbnails, thumbnail)
	}
	return thumbnails, nil
}
