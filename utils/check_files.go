package utils

import (
	"errors"
	"os"
)

func CheckFiles(args []string)([]string, error){
	var files []string
	for _, arg := range args {
		if _, err := os.Stat(arg); !os.IsNotExist(err) {
			// append arg to files
			files = append(files, arg)
		}
	}
	
	// if no files found exit
	if len(files) == 0 {
		return nil, errors.New("no files found")
	}

	return files, nil
}