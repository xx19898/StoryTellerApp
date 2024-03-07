package imagestoragehelpers

import (
	"testing"
)

func TestStoriesUserFolderExistanceFunc(t *testing.T) {
	result := StoriesImagesFolderForUserExists("testUser")

	if !result {
		t.Fatal("testUser folder should exist inside IMAGES/stories/")
	}

	result = StoriesImagesFolderForUserExists("non existant")

	if result {
		t.Fatal("non existant should not exist inside IMAGES/stories")
	}
}

func TestThatFolderCreationForUserWorks(t *testing.T) {
	CreateUserDirInStoriesFolder("tx")

	result := StoriesImagesFolderForUserExists("tx")

	if !result {
		t.Fatal("folder tx should exist inside IMAGES/stories")
	}

	DeleteUserDirInStoriesFolder("tx")

	result = StoriesImagesFolderForUserExists("tx")

	if result {
		t.Fatal("folder tu should not exist inside IMAGES/stories")
	}
}
