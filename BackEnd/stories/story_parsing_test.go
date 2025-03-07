package stories

import (
	"fmt"
	"testing"
)

func TestHtmlElementTypeExtraction(t *testing.T) {
	testElement := "<wrong>"
	_,err := GetTypeOfElement(testElement,GetAllowedElementsAndPropertiesMap)
	
	if err == nil {
		fmt.Printf("Function did not see the error in %s",testElement)
		t.Fatal(err.Error())
	}
	
	fmt.Print(err.Error())
	//test that <div><div>,<div></p> does not go through
}