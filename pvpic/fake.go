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
		layout := "2006:01:02 03:04:05"
		timestamp := current_time.Format(layout)
		
		// change model to add current timestamp
		model["DateTimeDigitized"] = timestamp
		model["DateTimeOriginal"] = timestamp
		model["DateTime"] = timestamp
		model["PreviewDateTime"] = timestamp
	// } else {
		//// change model to add 00 timestamp
		// model["DateTimeDigitized"] = "0000:00:00 00:00:00"
		// model["DateTimeOriginal"] = "0000:00:00 00:00:00"
		// model["DateTime"] = "0000:00:00 00:00:00"
		// model["PreviewDateTime"] = "0000:00:00 00:00:00"
	}

	err := update.UpdateExif(img_reader, faked_writer, model)

	return faked.Bytes(), err
}
