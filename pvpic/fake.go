package pvpic

import (
	"bytes"
	"bufio"

	"github.com/sfomuseum/go-exif-update"
)

func Fake(img []byte, model map[string]interface{})([]byte, error){
	var faked bytes.Buffer
	var img_reader = bytes.NewReader(img)

	var faked_writer = bufio.NewWriter(&faked)
	
	update.UpdateExif(img_reader, faked_writer, model)

	return faked.Bytes(), nil
}
