package imagestorage

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

// func for checking whether folder with username exists
func StoriesImagesFolderForUserExists(username string) bool {
	parent := getParentPath()
	userFolderPath := filepath.Join(parent, "IMAGES", "stories", username)
	return CheckThatDirectoryExists(userFolderPath)
}

func UserFolderForImagesExists(username string) bool {
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

// func for creating folder with username
func CreateUserDirInStoriesFolder(username string) bool {
	os.Chdir("../")
	os.Chdir("./IMAGES")
	os.Chdir("./stories")
	os.Mkdir(username, 0755)
	return true
}

/*
func createStoryDir(username string, storyId uint) bool {

}

func CreateFolderForStoryIds(username string, storyId uint) {
	userFolderExists := UserFolderForImagesExists(username)

	if !userFolderExists {
		CreateUserDirInStoriesFolder(username)
	}

	createStoryDir(username, storyId)
}

*/

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

	return err
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

func CreateStoryFolder(username string, storyId uint) error {

	var err error

	if StoriesImagesFolderForUserExists(username) {
		err = os.Chdir("../")

		if err != nil {
			return err
		}

		x, _ := os.Getwd()

		fmt.Println("*******")
		fmt.Println(x)
		fmt.Println("*******")

		err = os.Chdir("IMAGES")

		if err != nil {
			os.Chdir("./image")
			return err
		}

		err = os.Chdir("stories")

		if err != nil {
			return err
		}

		err = os.Chdir(username)

		if err != nil {
			return err
		}

		err = os.Mkdir(fmt.Sprint(storyId), 0755)

		if err != nil {
			return err
		}

		return err
	} else {
		CreateUserDirInStoriesFolder(username)

		currDir, _ := os.Getwd()
		os.Chdir(currDir)

		currDir, _ = os.Getwd()

		fmt.Println("****")
		fmt.Println(currDir)
		fmt.Println("****")

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

		err = os.Chdir(username)

		if err != nil {
			return err
		}

		err = os.Mkdir(fmt.Sprint(storyId), 0755)

		if err != nil {
			return err
		}

		return err
	}
}

//func for creating func for creating folder with storyid name exists
