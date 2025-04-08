package stories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)
func CheckOriginForImageSource(url string,correctSource string)(error){
	//check for length > 7 and < 50
	correctSourceLength := len([]rune(correctSource))

	urlOriginAsRuneSlice := []rune(url)[:correctSourceLength]
	urlOriginAsString := string(urlOriginAsRuneSlice)
	if( urlOriginAsString != correctSource) {
		fmt.Println(urlOriginAsString)
		return errors.New("Incorrect origin: " + urlOriginAsString)
	}

	return nil
}
func GetAllowedElementsAndPropertiesMap()([]string,map[string][]string){
	allowedElementTagsWithProperties := map[string] []string{
		"h":[]string{},
		"p":[]string{},
		"img":[]string{"src"},}	
	allowedElements := []string{"h","p","img"}

	return allowedElements,allowedElementTagsWithProperties
}
func htmlTagIsAllowed(element string)(bool, error){
	allowedElements,_ := GetAllowedElementsAndPropertiesMap()
	
	for _,allowedElement := range allowedElements{
		if allowedElement == element{
			return true,nil
		}
	}

	return false,nil
}

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

func scrollToFirstNonSpaceChar(curr *int  ,story []rune){
	fmt.Println("******")
	fmt.Println(string(story[*curr:]))
	fmt.Println("******")
	for _,char := range story[*curr:]{
		if(char == ' '){
			fmt.Println(fmt.Sprintf("curr char: %s",string(char)))	
			*curr = *curr + 1
		}else{
			break
		}
	}
}

func ParseHtmlAttribute(htmlAttributeString string)(string,string,error){
	splitByEqSign := strings.Split(htmlAttributeString,"=")

	if len(splitByEqSign) != 2{
		return "","",errors.New(fmt.Sprintf("Malformed Html Attribute: %s",htmlAttributeString))
	}

	return splitByEqSign[0],strings.ReplaceAll(splitByEqSign[1],"\"",""),nil
}

func OnOpeningBracketEncountered(currIndex *int, story []rune,openedTag *string)(bool,error){
	//currIndex at whatever comes next after <
	*currIndex++
	if *currIndex >= len(story){
		return false,errors.New(fmt.Sprintf("Unclosed < at index %s(end of the story)",strconv.Itoa(*currIndex)))
	}
	// openedTag == NONE => has to be legit opening tag (if not img tag)
	// tag name has to be after opening bracket
	// if img => src property HAS to be in place if not => > should follow the opening tag
	
	// openedTag != NONE => does not have to be legit opening/closing tag => 
	// in case of error IF tag is not enclosed with > just return the last position
	// => scroll until legit closing tag is found
	var tagNameBuilder strings.Builder

	scrollToFirstNonSpaceChar(currIndex,story)


	// Checks whether tag is a closing one
	if(story[*currIndex] == '/'){
		if *openedTag == "NONE"{
			return false,errors.New(fmt.Sprintf("Improper opening tag synthaxis at index %s",strconv.Itoa(*currIndex)))
		}

		//reading closing tag name
		for j := *currIndex;j < len(story);j++{
			*currIndex++
			char := story[j]
			if char == ' '{
				break
			}
			if char == '>'{
				break
			}else{
				tagNameBuilder.WriteRune(char)
			}
		}

		*currIndex++

		scrollToFirstNonSpaceChar(currIndex,story)
		if *currIndex >= len(story){
			return false,errors.New("")
		}
		if story[*currIndex] != '>'{
			return false,nil
		}

		if story[*currIndex] != '>'{
			if *openedTag != tagNameBuilder.String(){
				return false,errors.New(fmt.Sprintf("Error at index %s: closing tag does not have same tag name as the opened one currently(%s)",strconv.Itoa(*currIndex),*openedTag))
			}
			return false,nil
		}
		
	}

	if *openedTag != "NONE"{
		return false,errors.New(fmt.Sprintf("Error at index %s, there are two embedded tags. new tag opens after the last  one (%s) was not closed",strconv.Itoa(*currIndex),*openedTag))
	}
	// TAG IS NOT A CLOSING TAG

	// reads tag
	for{
		//checks that boundary is kept before grabbing the char to evade error
		if(*currIndex >= len(story)){
			return false, errors.New("Reached end of the story and did not encounter closing tag")
		}
		char := story[*currIndex]
		if(*currIndex == len(story) - 1){
			if(char) != '>'{
				return false,errors.New("Reached the end of the story and last char is not >")
			}
			return true,nil
		}
		if(char == ' ' || char == '>'){
			break
		}
		tagNameBuilder.WriteRune(char)
		*currIndex++
	} 


	switch tagType := tagNameBuilder.String(); tagType{
		case "img":
			if story[*currIndex] != '/'{
				return false,errors.New(fmt.Sprintf("Error at char %s, img element missing slash",strconv.Itoa(*currIndex)))
			}
			scrollToFirstNonSpaceChar(currIndex,story)
			if story[*currIndex] != '>'{
				return false,errors.New(fmt.Sprintf("Error at char %s, img element in incorrect form, should encounter> next, but got %s instead",strconv.Itoa(*currIndex),string(story[*currIndex])))
			}
			return true,nil 
		default:
			fmt.Println("****")
			fmt.Println("Currently at char " + string(story[*currIndex]))
			fmt.Println("****")
			scrollToFirstNonSpaceChar(currIndex,story)
			if story[*currIndex] != '>'{
				return false,errors.New(fmt.Sprintf("Error at char %s, html element is in incorrect form, should encounter > next, but got %s instead",strconv.Itoa(*currIndex),string(story[*currIndex])))
			}
			*openedTag = tagType 
			return true,nil		
	}
}


/*
func CheckStory(story string)(error){
	storyAsRuneArr := []rune(story)
	
	err := prelimCheckStory([]rune(story))
	if(err != nil){
		return err
	} 

	for{
		
	}


	return nil
}
	*/