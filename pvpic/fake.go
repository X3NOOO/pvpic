package pvpic

import (
	"bytes"

	"github.com/sfomuseum/go-exif-update"
)

func Fake(img []byte, model map[string]interface{})([]byte, error){
	var faked []byte
	var img_reader = bytes.NewReader(img)

	var faked_writer = bytes.NewBuffer(faked)
	
	update.UpdateExif(img_reader, faked_writer, model)

	return faked, nil
}
