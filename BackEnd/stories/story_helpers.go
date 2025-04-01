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
	for _,char := range story[*curr:]{
		if(char == ' '){	
			*curr = *curr + 1
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

	if(story[*currIndex] == '/'){
		// tag is a closing tag
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
	for j := *currIndex;j < len(story);j++{
		*currIndex++
		char := story[*currIndex]
		if char == ' '{
			break
		}
		if char == '<'{
			return false,errors.New("")
		}
		if char == '>'{
			break
		}else{
			tagNameBuilder.WriteRune(char)
		}
	}


	//TODO: from here on divide program into different cases for different tags: img, and others
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
			scrollToFirstNonSpaceChar(currIndex,story)
			if story[*currIndex] != '>'{
				return false,errors.New(fmt.Sprintf("Error at char %s, img element in incorrect form, should encounter> next, but got %s instead",strconv.Itoa(*currIndex),string(story[*currIndex])))
			}
			*openedTag = tagType 
			return true,nil		
	}

	/*
	scrollToFirstNonSpaceChar(currIndex,story) 

	tagClosed := story[*currIndex] == '>'

	TODO AFTER THE CLOSING > IS FOUND
	tagNameIsOK,_ := htmlTagIsAllowed(tagNameBuilder.String())
	
	if !tagNameIsOK && tagClosed{
		return false, errors.New(fmt.Sprintf("this is no tag at index %s",strconv.Itoa(*currIndex)))
	}

	//??
	if tagNameBuilder.String() != "img" && tagClosed{
		return true,nil
	}

	//tag is img and tag opened == NONE


	//if no attribute, opening tag is already parsed
	if(closingTag){
		if *openedTag == "NONE"{
			return false,errors.New(fmt.Sprintf("malformed story, found closing tag %s at index %s while there is no opening tag",tagNameBuilder.String(),strconv.Itoa(*currIndex)))
		}
		//if closing tag, close openedtag, if opened none => return error
		if tagClosed && closingTag{
			if *openedTag == tagNameBuilder.String(){
				*openedTag = "NONE"
				return true,nil
			}
			return false, errors.New(fmt.Sprintf("malformed story, wrong closing tag %s at index %s",&tagNameBuilder.String(),strconv.Itoa(*currIndex)))

		}
		return tagNameIsOK,nil
	}

	scrollToFirstNonSpaceChar(&currIndex,story)

	if(closingTag && story[*currIndex] != '>'){
		if(*openedTag == "NONE"){
			return false, errors.New(fmt.Sprintf("malformed story - no tag is opened, but there is none tag opening text at index %s",strconv.Itoa(*currIndex)))
		}
		return false, errors.New("malformed closing tag index " + strconv.Itoa(*currIndex) + " into the story")
	}

	//parsing html attribute value
	for j := *currIndex;j < len(story);j++{
		*currIndex++
		char := story[j]
		if char != ' '{
			propertyBuilder.WriteRune(char)
		}
	}
	return true,nil
	*/ 
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