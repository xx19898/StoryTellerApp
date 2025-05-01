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

	attributeName,attributeValue,err := ParseHtmlAttribute(htmlAttribute)
	if err != nil{
		t.Fatalf("should be no error here, but there is one: %s",err.Error())
	}
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

func TestGrabbingNextCharSeq(t* testing.T){
	testString := "<img src=\"wwww.storyteller.com\""
	index := 5
	

	word,err := GrabNextCharSeq([]rune(testString),&index)

	if index != len(testString) - 1{
		t.Fatalf("Index should be at %s, but is %s",strconv.Itoa(len(testString) - 1),strconv.Itoa(index))
	}

	if err != nil{
		t.Fatalf("Error should be nil, but it is %s",err.Error())
	}

	if word != "src=\"wwww.storyteller.com\""{
		t.Fatalf("word should be %s but got %s","src=\"wwww.storyteller.com\"",word)
	}
}

func TestCheckingAttributeIsCorrectForTheTag(t *testing.T){
	testAttribute := "src"
	tag := "img"

	attributeOk,error := AttributeAllowedForElement(tag,testAttribute)
	if error != nil{
		t.Fatalf("error when checking that attribute is okay for the chosen tag: %s",error.Error())
	}
	if !attributeOk{
		t.Fatalf("Attribute %s should be allowed for tag %s",testAttribute,tag)
	}


	testAttribute = "xdd"
	tag = "div"

	attributeOk,error = AttributeAllowedForElement(tag,testAttribute)

	if attributeOk{
		t.Fatalf("error when checking that %s is okay for the tag %s, should be registered as an error",testAttribute,tag)
	}

	testAttribute = "xdd"
	tag = "divxdd"

	attributeOk,error = AttributeAllowedForElement(tag,testAttribute)

	if error == nil{
		t.Fatalf("error when checking that %s is okay for the tag %s, tag is incorrect",testAttribute,tag)
	}
}

func TestParsingProperties(t * testing.T){
	testString := "<img src=\"www.storyteller.com\""

	index := 5
	properties, err := ParseProperties(&index,[]rune(testString),"img")

	if err != nil{
		t.Fatalf("Error thrown: %s", err.Error())
	}

	if properties["src"] != "www.storyteller.com"{
		t.Fatalf("Value respective of src should be www.storyteller.com but got %s instead",properties["src"])
	}
}

func TestControllingProperties(t *testing.T){
	testPropertiesMap := map[string]string{}

	err := ControlProperties(testPropertiesMap,"div")
	if err != nil{
		t.Fatalf("no error is expected when having div element without any properties but got %s instead",err.Error())
	}

	testPropertiesMap = map[string]string{
		"src":"www.google.com",
	}
	
	err = ControlProperties(testPropertiesMap, "div")
	if err == nil{
		t.Fatal("Error should occur when trying to check div element with property src")
	}

	err = ControlProperties(testPropertiesMap,"img")
	if err != nil{
		t.Fatalf("No error should occur when trying to check img element with attribute src")
	}

	testPropertiesMap = map[string]string{
	}
	err = ControlProperties(testPropertiesMap,"img")
	if err == nil{
		t.Fatalf("error should occure when trying to pass in img element without src attribute")
	}
}
func TestCheckingHtmlTag(t *testing.T){
 	testStory := []rune("<div>")
	
	index := 0
	openedTag := "NONE"
	var err error

	err = OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if index != 4{
		t.Fatalf("index should be %s but it is %s",strconv.Itoa(4),strconv.Itoa(index))
	} 
	if err != nil{
		t.Fatalf("error should be null,but got %s instead",err.Error())
	}
	
	if openedTag != "div"{
		t.Fatalf("opened tag should be div but got %s",openedTag)
	}

	testStory = []rune("<div")
	openedTag = "NONE"
	index = 0
	err = OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if err == nil{
		t.Fatalf("Error should be cast,but it is null.%s is an incorrect html tag",string(testStory))
	}
	
	testStory = []rune("<>")
	openedTag = "NONE"
	index = 0
	err = OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if err == nil{
		t.Fatalf("Error should be cast,but it is null.%s is an incorrect html tag",string(testStory))
	}

	testStory = []rune("<xddd>")
	openedTag = "NONE"
	index = 0
	err = OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if err == nil{
		t.Fatalf("Error should be cast,but it is null.%s is an incorrect html tag",string(testStory))
	}
	if index != len(testStory) - 1{
		t.Fatalf("Index should be %s, but it is at %s",strconv.Itoa(len(testStory) - 1),strconv.Itoa(index))
	}


	testStory = []rune("<xddd>")
	openedTag = "NONE"
	index = 0
	err = OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if err == nil{
		t.Fatalf("Error should be cast,but it is null.%s is an incorrect html tag",string(testStory))
	}
	if index != len(testStory) - 1{
		t.Fatalf("Index should be %s, but it is at %s",strconv.Itoa(len(testStory) - 1),strconv.Itoa(index))
	}

	testStory = []rune("<img>")
	openedTag = "NONE"
	index = 0
	err = OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if err == nil{
		t.Fatalf("Error should be cast,but it is null.%s should have an src property and a closing slash",string(testStory))
	}
	if index != len(testStory) - 1{
		t.Fatalf("Index should be %s, but it is at %s",strconv.Itoa(len(testStory) - 1),strconv.Itoa(index))
	}

	testStory = []rune("<h1>Hello World!</h1>")
	openedTag = "NONE"
	index = 0
	err = OnOpeningBracketEncountered(&index,testStory,&openedTag)

	if err != nil{
		t.Fatal(err.Error())
	}
}

func TestCheckingStory(t *testing.T){
	testStory := "<h1>Hello World!</h1><p>Today we gather here to rejoy our existance</p>"
	err := CheckStory(testStory)

	if err != nil{
		t.Fatalf("error should be nil when testing whether \" %s \" is a legit story, but got %s error instead ",testStory,err.Error())
	}

	testStory = "<div><div>Hello World!</div><div>Today we gather here to rejoy our existance</p>"
	err = CheckStory(testStory)
	
	if err == nil{
		t.Fatalf("error should be thrown when testing whether \" %s \" is a legit story(embedded elements present), but got nothing instead ",testStory)
	}

	testStory = "<div>Hello World<div>"
	err = CheckStory(testStory)
	if err == nil{
		t.Fatalf("error should be thrown when testing whether \" %s \" is a legit story(last div element is missing closing sign), but got nothing instead ",testStory)
	}

	testStory = "<div/>Hello World<div>"
	err = CheckStory(testStory)
	if err == nil{
		t.Fatalf("error should be thrown as first div tag has closing mark inside it")
	}

	testStory = "<h1>Main Title</h1><p>Main Content</p>"
	err = CheckStory(testStory)
	if err != nil{
		t.Fatalf("error should be nil but got %s instead (story is: %s)",err.Error(),testStory)
	}

	testStory = "<h1>Main Title</h1><p>Main Content"
	err = CheckStory(testStory)
	if err == nil{
		t.Fatalf("error should be thrown as the last p tag is not closed before the end of the story but got nil instead (story:%s)",testStory)
	}

	testStory = "Main Title</h1><p>Main Content</p>"
	err = CheckStory(testStory)
	if err == nil{
		t.Fatalf("error should be thrown as starting tag is not opened properly (%s) but got nil instead",testStory)
	}
}
