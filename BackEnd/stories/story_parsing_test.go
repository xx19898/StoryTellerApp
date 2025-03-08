package stories

import (
	"fmt"
	"testing"
)
func TestStoryHtmlSynthaxisIsCorrect(t *testing.T){
	testElement := "<wrong>"
	_,err := GetTypeOfElement(testElement,GetAllowedElementsAndPropertiesMap)
	
	if err == nil {
		t.Fatalf(fmt.Sprintf("Function did not see the error in %s",testElement))
	}
	
	fmt.Print(err.Error())
	
	testElement = "<div><div>"
	_,err = GetTypeOfElement(testElement,GetAllowedElementsAndPropertiesMap)
	fmt.Println(err.Error())
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not see the error in %s , missing / before the last >",testElement))
	}

	testElement = "<<div><div/>"
	_,err = GetTypeOfElement(testElement,GetAllowedElementsAndPropertiesMap)
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not see the error in %s ,there is no such html elements as wrong",testElement))
	}
}
func TestHtmlElementTypeExtraction(t *testing.T) {
	testElement := "<div><div/>"
	elementType,err := GetTypeOfElement(testElement, GetAllowedElementsAndPropertiesMap)
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not see the error in %s ,there is no such html elements as wrong",testElement))
	}
	if elementType != "div"{
		t.Fatalf("Incorrect element type is extracted")		
	}

}