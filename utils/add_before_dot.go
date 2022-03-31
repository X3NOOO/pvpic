package utils

import (
	"strings"
)

func AddBeforeDot(name string, add string) string {
	// get extension of file
	name_slice := strings.Split(name, ".")
	extension := name_slice[len(name_slice)-1]

	// get name without extension
	name_without_extension := strings.Join(name_slice[:len(name_slice)-1], ".")

	new_name := name_without_extension + add + "." + extension

	return new_name
}