package imagestoragehelpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func FormFilepathForStoryImage(username string, storyId uint) string {
	return filepath.Join("backend", "IMAGES", "stories", username, fmt.Sprint(storyId))
}

func CreateNewStoryImageFile(filename string, targetDirectory string, data []byte) error {
	filepath := filepath.Join("/", targetDirectory, filename)

	w, err := os.Create(filepath)

	if err != nil {

		return err
	}

	defer w.Close()

	err = os.WriteFile(filepath, data, 0644)

	return err
}

func CheckIfStoryImageFileExists(username string, storyId uint, filename string) bool {

	if _, err := os.Stat(filepath.Join(username, fmt.Sprint(storyId), filename)); err == nil {
		return true
	}
	return false
}

func DeleteStoryImage(username string, storyId uint, filename string) error {
	err := os.Remove(filepath.Join(username, fmt.Sprint(storyId), filename))

	return err
}

func CreateNewStoryImage(username string, storyId uint, data []byte, filename string) error {
	// if story/username folder does not exist - create

	userFolderPath := filepath.Join("backend", "IMAGES", "stories", username)
	userFolderExists := CheckThatDirectoryExists(userFolderPath)

	if !userFolderExists {
		CreateUserDirInStoriesFolder(username)
	}

	// if story/username/storyId does not exist - create
	storyFolderPath := filepath.Join(userFolderPath, fmt.Sprint(storyId))
	storyFolderExists := CheckThatDirectoryExists(storyFolderPath)

	if !storyFolderExists {
		CreateStoryFolder(username, storyId)
	}

	// if story/username/storyId/filename exist - delete, create new file
	imageWithSameNameAlreadyExists := CheckIfStoryImageFileExists(username, storyId, filename)

	if imageWithSameNameAlreadyExists {
		DeleteStoryImage(username, storyId, filename)
	}

	filepath := FormFilepathForStoryImage(username, storyId)

	err := CreateNewStoryImageFile(filename, filepath, data)

	return err
}
