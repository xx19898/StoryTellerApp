package imagestorage

import (
	"os"
	"path/filepath"
)

func UserFolderForImagesExists(username string) bool {
	parent := getParentPath()
	userFolderPath := filepath.Join(parent, "IMAGES", "stories", username)

	return CheckThatDirectoryExists(userFolderPath)
}

func CreateUserDirInStoriesFolder(username string) bool {
	os.Chdir("../")
	os.Chdir("./IMAGES")
	os.Chdir("./stories")
	os.Mkdir(username, 0755)
	goBackToImageStorageDir()

	return true
}

func DeleteUserDirInStoriesFolder(username string) error {
	var err error

	err = os.Chdir("../")

	if err != nil {
		return err
	}

	err = os.Chdir("IMAGES")

	if err != nil {
		return err
	}

	err = os.Chdir("stories")

	if err != nil {
		return err
	}

	err = os.Remove(username)

	if err != nil {
		return err
	}

	goBackToImageStorageDir()

	return err
}
