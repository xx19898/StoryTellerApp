package stories

import (
	"fmt"
	"testing"
)

func TestHtmlElementTypeExtraction(t *testing.T) {
	testElement := "<wrong>"
	_,err := GetTypeOfElement(testElement,GetAllowedElementsAndPropertiesMap)
	
	if err != nil {
		fmt.Printf("Element is malformed or type is not supported: %s",err)
		fmt.Print(err.Error())
		t.Fatal(err.Error())
	}
}