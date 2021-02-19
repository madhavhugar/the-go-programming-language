package channels

import (
	"fmt"
	"time"
)

func thumbnailConverter(filepath string) (thumbnailPath string) {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s_thumbnail", filepath)
}
