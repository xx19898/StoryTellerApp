package imagestorage

import (
	"os"
	"path/filepath"
)

func RetrieveFileFromImageFolder(filename string, subfolderName string) (*os.File, error) {
	path := filepath.Join("/backend", "IMAGES", subfolderName, filename)
	file, err := os.Open(path)
	return file, err
}
