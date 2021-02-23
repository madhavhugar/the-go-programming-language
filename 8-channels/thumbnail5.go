package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func thumbnailConverter5(file string) (string, error) {
	delay := time.Duration(time.Second*1) + time.Duration(rand.Float64()*1)
	time.Sleep(delay)
	if file == "error" {
		return "", fmt.Errorf("invalid file")
	}
	return fmt.Sprintf("%s_thumbnail", file), nil
}

func parallelGoRoutinesWithInputFromChannels(files chan string) error {
	wg := sync.WaitGroup{}
	errChan := make(chan error)

	for file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			_, err := thumbnailConverter5(f)
			errChan <- err
		}(file)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}
