package pvpic

import (
	"bufio"
	"bytes"
	"time"

	"github.com/sfomuseum/go-exif-update"
)

func Fake(img []byte, model map[string]interface{}, current_time bool) ([]byte, error) {
	var faked bytes.Buffer
	var img_reader = bytes.NewReader(img)

	var faked_writer = bufio.NewWriter(&faked)

	
	if(current_time){
		// get current timestamp in yyyy:mm:dd hh:mm:ss.sss
		current_time := time.Now()
		layout := "2006:01:02 3:4:5:06"
		timestamp := current_time.Format(layout)
		
		// change model to add current timestamp
		model["DateTimeDigitized"] = timestamp
		model["DateTimeOriginal"] = timestamp
		model["DateTime"] = timestamp
		model["PreviewDateTime"] = timestamp
	}

	err := update.UpdateExif(img_reader, faked_writer, model)

	return faked.Bytes(), err
}
