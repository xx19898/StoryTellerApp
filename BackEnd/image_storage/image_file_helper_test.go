package imagestorage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestThatStoriesFolderExistanceFolderOperatesProperly(t *testing.T) {
	result := StoriesImagesFolderForUserExists("testUser")

	if !result {
		t.Fatal("testUser folder should exist inside IMAGES/stories/")
	}

	result = StoriesImagesFolderForUserExists("non existant")

	if result {
		t.Fatal("non existant should not exist inside IMAGES/stories")
	}
}

func TestThatFolderForStoryIdExistanceCheckOperatesProperly(t *testing.T) {
	result := StoriesImagesFolderForStoryIdExists("testUser", uint(1))

	if !result {
		t.Fatal("folder 1 should exist inside IMAGES/stories/testUser/1")
	}

	result = StoriesImagesFolderForStoryIdExists("testUser", uint(0))

	if result {
		t.Fatal("folder 0 should not exist inside IMAGES/stories/testUser/1")
	}
}

func TestThatFolderCreationForUserWorks(t *testing.T) {
	CreateUserDirInStoriesFolder("tu")
	result := StoriesImagesFolderForUserExists("tu")

	time.Sleep(10 * time.Second)

	if !result {
		t.Fatal("folder tu should exist inside IMAGES/stories")
	} else {
		fmt.Println("HERE")
		os.Chdir("..")
		os.Chdir("IMAGES")
		os.Chdir("stories")
		os.Remove("tu")
	}

	result = StoriesImagesFolderForUserExists("tu")

	if result {
		t.Fatal("folder tu should not exist inside IMAGES/stories")
	}
}
