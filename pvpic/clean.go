package pvpic

import (
	"github.com/scottleedavis/go-exif-remove"
)

func Clean(img []byte) ([]byte,error) {
	clean, err := exifremove.Remove(img)
	
	if(err != nil) {
		return nil, err
	}

	return clean, nil
}
