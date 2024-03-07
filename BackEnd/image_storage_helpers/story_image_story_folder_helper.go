package imagestoragehelpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func StoriesImagesFolderForUserExists(username string) bool {
	parent := getParentPath()
	userFolderPath := filepath.Join(parent, "IMAGES", "stories", username)
	return CheckThatDirectoryExists(userFolderPath)
}

func StoriesImagesFolderForStoryIdExists(username string, storyId uint) bool {
	storyIdString := strconv.FormatUint(uint64(storyId), 10)
	parent := getParentPath()
	finalPath := filepath.Join(parent, "IMAGES", "stories", username, storyIdString)

	return CheckThatDirectoryExists(finalPath)
}

func CreateStoryFolder(username string, storyId uint) error {

	var err error

	if StoriesImagesFolderForUserExists(username) {
		err = os.Chdir("../")

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		err = os.Chdir("IMAGES")

		if err != nil {
			goBackToImageStorageDir()
			os.Chdir("./image")
			return err
		}

		err = os.Chdir("stories")

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		err = os.Chdir(username)

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		err = os.Mkdir(fmt.Sprint(storyId), 0755)

		if err != nil {
			goBackToImageStorageDir()
			return err
		}
		goBackToImageStorageDir()

		return err
	} else {
		CreateUserDirInStoriesFolder(username)

		currDir, _ := os.Getwd()
		os.Chdir(currDir)

		err = os.Chdir("../")

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		err = os.Chdir("IMAGES")

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		err = os.Chdir("stories")

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		err = os.Chdir(username)

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		err = os.Mkdir(fmt.Sprint(storyId), 0755)

		if err != nil {
			goBackToImageStorageDir()
			return err
		}

		goBackToImageStorageDir()

		return err
	}
}

func DeleteStoryFolder(username string, storyId uint) error {
	var err error
	err = os.Chdir("../")

	if err != nil {
		goBackToImageStorageDir()
		return err
	}

	err = os.Chdir("IMAGES")

	if err != nil {
		goBackToImageStorageDir()
		return err
	}

	err = os.Chdir("stories")

	if err != nil {
		goBackToImageStorageDir()
		return err
	}

	err = os.Chdir(username)

	if err != nil {
		goBackToImageStorageDir()
		return err
	}

	err = os.Remove(fmt.Sprint(storyId))

	goBackToImageStorageDir()

	return err
}
