package imagestoragehelpers

import (
	"fmt"
	"testing"
)

func TestFolderForStoryIdExistanceCheckFunc(t *testing.T) {
	result := StoriesImagesFolderForStoryIdExists("testUser", uint(1))

	if !result {
		t.Fatal("folder 1 should exist inside IMAGES/stories/testUser/1")
	}

	result = StoriesImagesFolderForStoryIdExists("testUser", uint(0))

	if result {
		t.Fatal("folder 0 should not exist inside IMAGES/stories/testUser/1")
	}
}

func TestThatStoryFolderGetsCreatedAndDeletedCorrectly(t *testing.T) {
	res := StoriesImagesFolderForStoryIdExists("ts", 1)

	if res {
		t.Fatal("Folder stories/ts/1 should not exist")
	}

	err := CreateStoryFolder("ts", 1)

	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not create folder ts/1 in stories")
	}

	err = DeleteStoryFolder("ts", 1)

	if err != nil {
		fmt.Println(err)
		t.Fatal("Could not delete folder ts/1 in stories")
	}

	res = StoriesImagesFolderForStoryIdExists("ts", 1)

	if res {
		t.Fatal("stories/ts/1 still exists after the deletion")
	}

	DeleteUserDirInStoriesFolder("ts")
}
