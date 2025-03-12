package stories

import (
	"errors"
	"fmt"
	"strings"
)

func GetAllowedElementsAndPropertiesMap()(map[string][]string,error){
	allowedElementTagsWithProperties := map[string] []string{
		"h":[]string{},
		"p":[]string{},
		"img":[]string{"src"},}
	return allowedElementTagsWithProperties,nil
}
type elementsPropertiesFunctionType func() (map[string][]string,error)
/*

make(
	map[string][]string{
		"h":[]string{},
		"p":[]string{},
		"img":[]string{"src"},
}
)
*/



//checks that story is not empty nor too long
func prelimCheckStory(story []rune) error {
	if len(story) > 500 {
		return errors.New("Story too long")
	}
	if(len(story) == 0) {
		return errors.New("Story is empty")
	}
	return nil
}

func GetOpenerTag(story []rune,pointer int)(string,error,int){
	if string(story[pointer]) != "<"{
		return "",errors.New("First character is not \"<\""),0
	}
	if len(story[pointer:]) < 3{
		return "",errors.New("Cannot extract opener tag as the bit of story passed in is too short(smaller then 3 elements)"),0
	}
	var openingTag strings.Builder
	pointerLastPosition := pointer
	
	//pointer should be at '<'
	for i := pointer + 1; i < len(story);i++{
		pointerLastPosition++
		charAtHand := story[i]
		if string(charAtHand) == " " || string(charAtHand) == ">"{break}

		openingTag.WriteRune(charAtHand)
	}

	return openingTag.String(),nil,pointerLastPosition
}

// note: no children expected to be in any html node
func CheckStoryHtmlSynthaxis(story string) (error){
	trimmedElement := []rune(strings.TrimSpace(story))

	err := prelimCheckStory(trimmedElement)
	if err != nil{
		return err
	}

	if(len(trimmedElement) == 0){
		return errors.New("Element is empty")
	}
	

	//	get opener tag <div>, <p> etc if no opener => err
	//	GetOpenerTag func

	//	get properties

	//	check if properties are ok for the chosen opener

	//	what to do if there is closing tag inside the comments? 
	// 	ex: <div>and the closing tag for div should look like: </div> </div>

	//	get closer tag </div>, </p> etc if no closer => err

	//	move on to the next html element if there are any




	//TODO: check that <> </> are ok, what of escaped <'s and >'s?
	
	firstChar := string(trimmedElement[0])
	if(firstChar != "<") {
		fmt.Println("first char is  " + firstChar)
		return errors.New("First char is not \"<\" " + "but " + firstChar)
	}

	lastChar := string(trimmedElement[len(trimmedElement) - 1])
	if(lastChar != ">"){
		return errors.New("Last char " + "(" + lastChar + ")" + "is not \">\"")
	}

	// loop through characters, not bytes
	secondLastChar := string(trimmedElement[len(trimmedElement) - 2])
	if(secondLastChar != "/"){
		return errors.New("Second last char is not \"\\\"")
	}

	// check that html element is semantically proper: <> </>
	var checkString strings.Builder
	for _,char := range trimmedElement{
		for _,keyElement := range []rune{'>','<','/'}{
			if(char == keyElement){
				checkString.WriteRune(char)
			}
		}
	}
	if checkString.String() != "<></>"{
		return errors.New("the html element is not semantically proper: <></>")		
	}
	
	return nil
}

func GetTypeOfElement(element string,getElementsPropertiesMap elementsPropertiesFunctionType) (string, error){
	
	trimmedElement := []rune(strings.TrimSpace(element))
	  
	if(len(trimmedElement) == 0){
		return "",  errors.New("Element is empty")
	}
	
	/*
	allowedElementsAndPropertiesMap,err := getElementsPropertiesMap()
	if(err != nil){
		return "",errors.New("Could not get allowedElementsAndPropertiesMap")
	}
	*/

	//TODO: check that <> </> are ok, what of escaped <'s and >'s?
	
	firstChar := string(trimmedElement[0])
	if(firstChar != "<") {
		fmt.Println("first char is  " + firstChar)
		return "",errors.New("First char is not \"<\" " + "but " + firstChar)
	}

	lastChar := string(trimmedElement[len(trimmedElement) - 1])
	if(lastChar != ">"){
		return "",errors.New("Last char " + "(" + lastChar + ")" + "is not \">\"")
	}

	// loop through characters, not bytes
	secondLastChar := string(trimmedElement[len(trimmedElement) - 2])
	if(secondLastChar != "/"){
		return "",errors.New("Second last char is not \"\\\"")
	}

	// check that html element is semantically proper: <> </>
	var checkString strings.Builder
	for _,char := range trimmedElement{
		for _,keyElement := range []rune{'>','<','/'}{
			if(char == keyElement){
				checkString.WriteRune(char)
			}
		}
	}
	if checkString.String() != "<></>"{
		return "",errors.New("the html element is not semantically proper: <></>")		
	}

	var elementType strings.Builder

	for _, char := range trimmedElement{
		isPartOfElType := true
		for _, elementTypeEndingSymbol := range []rune {'>',' '}{
			if char == elementTypeEndingSymbol{
				isPartOfElType = false
				break
			}
		}
		if !isPartOfElType{
			break
		}
		elementType.WriteRune(char)
	}

	return elementType.String(),nil
}

/*
func SanitizeStoryHtmlString(unsntzd string) (string,error) {
	storyIsOk,err := prelimCheckStory(unsntzd)
	if(!storyIsOk || err != nil){
		return "",err
	} 
	
	var currElementType = ""
	
	//element type can only be h1,2,3..., p, img
	for index, char := range unsntzd {

	}
}
*/
