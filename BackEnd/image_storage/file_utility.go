package imagestorage

import (
	"os"
)

/*

func checkThatFileIsImage(data []byte) bool {
	dataType := http.DetectContentType(data)

	isImage := (dataType == "image")

	fmt.Println(dataType)

	return isImage
}

*/

func CheckThatDirectoryExists(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}
