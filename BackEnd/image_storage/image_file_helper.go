package imagestorage

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

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
	fmt.Println("*****")
	fmt.Println(userFolderPath)
	fmt.Println("*****")
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

func DeleteUserDirInStoriesFolder(username string) {
	os.Chdir("..")
	os.Chdir("IMAGES")
	os.Chdir("stories")
	os.Remove(username)
}

//func for creating func for creating folder with storyid name exists
