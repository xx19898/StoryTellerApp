package stories

import (
	"fmt"
	"strconv"
	"testing"
)


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


func TestHtmlAttributeParsing(t *testing.T){
	htmlAttribute := "src=\"www.google.com\""

	attributeName,attributeValue,_ := ParseHtmlAttribute(htmlAttribute)
	if attributeName != "src"{
		t.Fatalf(fmt.Sprintf("Incorrectly parsed html attribute. Should ve gotten \"src\" but got %s instead",htmlAttribute))
	}

	if attributeValue != "www.google.com"{
		t.Fatalf(fmt.Sprintf("Incorrectly parsed html attribute value. Should ve gotten \"www.google.com\" but got %s instead",attributeValue))
	}

}

func TestHtmlStringScrollingToFirstNonEmptyChar(t *testing.T){
	htmlAttribute := "  <div>xddddd"
	i := 0
	scrollToFirstNonSpaceChar(&i,[]rune(htmlAttribute))
	
	fmt.Println(i)
	if i != 2{
		t.Fatal(fmt.Sprintf("did not scroll through the empty space correctly: i is %s and not 2",strconv.Itoa(i) ))
	}

	htmlElement := "<div>xddd</div>"
	i = 4
	scrollToFirstNonSpaceChar(&i,[]rune(htmlElement))
	if i != 4{
		t.Fatal(fmt.Sprintf("did not scroll through the empty space correctly: i is %s and not 4",strconv.Itoa(i) ))
	}
}

func TestCheckingIfThereIsHtmlTagInSpecificPlace(t *testing.T){
	testStory := []rune("<div> xddd</div>")
	
	index := 0
	openedTag := "NONE"

	bracketEncounter,err := OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if index != 4{
		t.Fatal(fmt.Sprintf("Index should be %s but it is %s",strconv.Itoa(5),strconv.Itoa(index)))
	}
	if err != nil{
		t.Fatal("Error should be null")
	}
	if !bracketEncounter{
		t.Fatal("Should be true")
	}
	if openedTag != "div"{
		t.Fatal(fmt.Sprintf("opened tag should be div but got %s",openedTag))
	}
}