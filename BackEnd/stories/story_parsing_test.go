package stories

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetOpenerTag(t *testing.T){
	testElement := []rune("<d")
	_,err,_ := GetOpenerTag(testElement,0)
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not catch story being too short(<3)"))
	}

	testElement = []rune("x<div/>")
	_,err,_ = GetOpenerTag(testElement,0)
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not catch story being too short(<3)"))
	}

	testElement = []rune("<div>")
	res ,err,_ := GetOpenerTag(testElement,0)
	if err != nil{
		t.Fatalf(fmt.Sprintf("Unexpected error, %s",err.Error()))
	}
	if res != "div"{
		t.Fatalf("Incorrect opener tag extracted")
	}

	testElement = []rune("<p>")
	res ,err,_ = GetOpenerTag(testElement,0)
	if err != nil{
		t.Fatalf(fmt.Sprintf("Unexpected error, %s",err.Error()))
	}
	if res != "p"{
		t.Fatalf("Incorrect opener tag extracted")
	}


	testElement = []rune("<div >")
	res ,err,_ = GetOpenerTag(testElement,0)
	if err != nil{
		t.Fatalf(fmt.Sprintf("Unexpected error, %s",err.Error()))
	}
	if res != "div"{
		t.Fatalf("Incorrect opener tag extracted")
	}

	testElement = []rune("<div>")
	res ,err,_ = GetOpenerTag(testElement,0)
	if err != nil{
		t.Fatalf(fmt.Sprintf("Unexpected error, %s",err.Error()))
	}
	if res != "div"{
		t.Fatalf("Incorrect opener tag extracted")
	}

	testElement = []rune("<img src=\"xdfdsjfjsd\">")
	res ,err, pointersLastPosition := GetOpenerTag(testElement,0)
	if err != nil{
		t.Fatalf(fmt.Sprintf("Unexpected error, %s",err.Error()))
	}
	if res != "img"{
		t.Fatalf("Incorrect opener tag extracted")
	}
	if pointersLastPosition != 4{
		t.Fatal(fmt.Sprintf("Incorrect pointerlastposition, should be 3 but is %s",strconv.Itoa(pointersLastPosition)))
	} 
}
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
		t.Fatalf(fmt.Sprintf("Function did not see the error in %s ,missing / before the last >",testElement))
	}

	testElement = "<<div><div/>"
	_,err = GetTypeOfElement(testElement,GetAllowedElementsAndPropertiesMap)
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not see the error in %s ,there is no such html elements as wrong",testElement))
	}
}

func TestThatHtmlPropertiesAreCorrect(t *testing.T){
	testElement := "<div incorrectProperty=\"xddd\"></div>"
	err := CheckStoryHtmlSynthaxis(testElement)
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not see error inside %s",testElement))
	}

	testElement = "<div></div>"
	err = CheckStoryHtmlSynthaxis(testElement)
	if err != nil{
		t.Fatalf(fmt.Sprintf("Function saw error inside %s",testElement))
	}

	testElement = "<img src=\"www.storyteller.com/story_images/3999\"></img>"
	err = CheckStoryHtmlSynthaxis(testElement)
	if err != nil{
		t.Fatalf(fmt.Sprintf("Function did not see error inside %s",testElement))
	}	
}

func TestOriginForSrcTagInsideImages(t *testing.T){
	testElement := "www.scam.com/imgs/2q30_\""
	err := CheckOriginForImageSource(testElement,"www.storyteller.com")
	if err == nil{
		t.Fatalf(fmt.Sprintf("Function did not see error inside %s",testElement))
	}

	testElement = "www.storyteller.com/imgs/2q30_\""
	err = CheckOriginForImageSource(testElement,"www.storyteller.com")
	if err != nil{
		t.Fatalf(fmt.Sprintf("Function saw inside %s ,should be ok and no error",testElement))
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