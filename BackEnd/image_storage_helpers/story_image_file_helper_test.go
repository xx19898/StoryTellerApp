package imagestoragehelpers

import (
	"bytes"
	"os"
	"testing"
)

// test taking contents of test image sun and creating new story image to user1/1/ and comparing contents of file to the original
func TestStoryImageCreating(t *testing.T) {
	testData, err := os.ReadFile("/backend/IMAGES/test_image_sun.jpg")

	if err != nil {
		t.Fatal(err)
	}

	err = CreateNewStoryImage("testUser2", uint(2), testData, "test_image_1.jpg")

	if err != nil {
		t.Fatal(err)
	}

	dataToCheck, err := os.ReadFile("/backend/IMAGES/stories/testUser2/2/test_image_1.jpg")

	if err != nil {
		t.Fatal("Could not read the created file")
	}

	if !bytes.Equal(dataToCheck, testData) {
		t.Fatal("Data does not match")
	}

	DeleteStoryImage("testUser2", uint(2), "test_image_1.jpg")
}
