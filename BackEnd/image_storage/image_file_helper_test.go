package imagestorage

import (
	"fmt"
	"os"
	"testing"
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
	CreateUserDirInStoriesFolder("tx")

	// resets working directory from IMAGES/stories to /image_storage
	os.Chdir("../")
	os.Chdir("../")
	os.Chdir("./image_storage")

	result := StoriesImagesFolderForUserExists("tx")

	if !result {
		t.Fatal("folder tx should exist inside IMAGES/stories")
	} else {
		fmt.Println("HERE")
		os.Chdir("..")
		os.Chdir("IMAGES")
		os.Chdir("stories")
		os.Remove("tx")
	}

	// resets working directory from IMAGES/stories to /image_storage
	os.Chdir("../")
	os.Chdir("../")
	os.Chdir("./image_storage")

	result = StoriesImagesFolderForUserExists("tx")

	if result {
		t.Fatal("folder tu should not exist inside IMAGES/stories")
	}
}

func TestThatStoryFolderExistanceCheckWorks(t *testing.T) {
	res := StoriesImagesFolderForStoryIdExists("testUser", 1)

	if !res {
		t.Fatal("Folder stories/testUser/1 should exist")
	}

	res = StoriesImagesFolderForStoryIdExists("testUser", 2)

	if res {
		t.Fatal("Folder stories/testUser/2 should not exist")
	}
}

// TODO: test creation of story folder inside IMAGES/username directory
func TestThatStoryFolderGetsCreatedAndDeletedCorrectly(t *testing.T) {
	res := StoriesImagesFolderForStoryIdExists("ts", 1)

	mydir, _ := os.Getwd()

	os.Chdir(mydir)

	if res {
		t.Fatal("Folder stories/ts/1 should not exist")
	}

	err := CreateStoryFolder("ts", 1)

	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not create folder ts/1 in stories")
	}

	os.Chdir(mydir)

	err = DeleteStoryFolder("ts", 1)

	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not delete folder ts/1 in stories")
	}

	os.Chdir(mydir)

	res = StoriesImagesFolderForStoryIdExists("ts", 1)

	if !res {
		t.Fatal("stories/ts/1 still exists after the deletion")
	}
}

/*

func TestThatStoryFolderGetsCreatedCorrectly(t *testing.T) {

}

// TODO: create functionality and test for creating new image under non existing user and image name
func TestThatImageCreatingForNonExistantUsersWorks(t *testing.T) {
	testImageBytes, err := os.ReadFile("../IMAGES/test_image_sun.jpg")

	if err != nil {
		t.Fatal("unable to read file")
	}

	//Create directory for said user

}

*/
