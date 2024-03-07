package imagestoragehelpers

import (
	"os"
	"path/filepath"
)

func goBackToImageStorageDir() {
	os.Chdir("/backend/image_storage")
}

func CheckThatDirectoryExists(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func getParentPath() string {
	currPath, _ := os.Getwd()
	parent := filepath.Dir(currPath)

	return parent
}
