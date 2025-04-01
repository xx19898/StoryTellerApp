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

func onOpeningBracketEncountered(currIndex *int, story []rune,openedTag *string)(bool,error){
	//currIndex at whatever comes next after <
	*currIndex++
	
	// openedTag == NONE => has to be legit opening tag (if not img tag)
	// tag name has to be after opening bracket
	// if img => src property HAS to be in place if not => > should follow the opening tag
	
	// openedTag != NONE => does not have to be legit opening/closing tag => 
	// in case of error IF tag is not enclosed with > just return the last position
	// => scroll until legit closing tag is found
	var tagNameBuilder strings.Builder
	var propertyBuilder strings.Builder
	closingTag := false 

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

		scrollToFirstNonSpaceChar(currIndex,story)

		if story[*currIndex] != '>'{
			if *openedTag != tagNameBuilder.String(){
				return false,errors.New(fmt.Sprint("Error at index %s: closing tag does not have same tag name as the opened one currently(%s)",strconv.Itoa(*currIndex),*openedTag))
			}
			return false,nil
		}
		
	}

	// TAG IS NOT A CLOSING TAG

	// reads tag 
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

	//TODO: from here on divide program into different cases for different tags: img, and others

	scrollToFirstNonSpaceChar(currIndex,story)

	tagClosed := story[*currIndex] == '>'
/*	
	TODO AFTER THE CLOSING > IS FOUND
	tagNameIsOK,_ := htmlTagIsAllowed(tagNameBuilder.String())
	
	if !tagNameIsOK && tagClosed{
		return false, errors.New(fmt.Sprintf("this is no tag at index %s",strconv.Itoa(*currIndex)))
	}
*/ 
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
		if(openedTag == "NONE"){
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
}


/*
func CheckStory(story string)(error){
	storyAsRuneArr := []rune(story)
	err := prelimCheckStory([]rune(story))
	if(err != nil){
		return err
	} 

	//Means opening tag already exists
	openedTag := "NONE"
	// status is NONE, OPENING_TAG, CLOSING_TAG, CONTENT
	STATUS := "NONE"

	_,allowedHtmlAttributesMap,err := GetAllowedElementsAndPropertiesMap()
	

	var tagTypeBuilder strings.Builder
	tagBuilt := false
	
	var htmlAttributeBuilder strings.Builder
	attributeBuilt := false

	// STATES: CLOSED, OPENED,(also need to know which tag is opened): 
	// if closed: next should be opening tag, 
	// if opened: scroll until <, 
	
	// MAKE THIS A FUNC:
	// extract tag type, extract all the properties
	// look if it is tag:
	// take in <, next one should be tag name!, check if tag is proper, 
	// check if tag has none attributes, 
	// check for attributes if needed, 
	// check that attribute is correct, 
	// if there are any chars after that attribute => error, 
	// scroll until >, 
	
	// set status to opened, 
	// set type of tag too close tag, 
	// scroll until next < => check that it is correct closing tag => 
	// if not scroll again until next < => so on... 
	
	
	for i := 0;i < len(storyAsRuneArr); i++{
		//
		if(storyAsRuneArr[i] == '<'){
			STATUS = "TAG"
			// checking if there is "/" after "<", then it is closing tag
			scrollToFirstNonSpaceChar(&i,storyAsRuneArr)
			if(i >= len(storyAsRuneArr)){
				break
			}
			if(storyAsRuneArr[i] == '/'){
				var closingTagBuilder strings.Builder

				for _,char := range storyAsRuneArr[i:]{
						if(char == ' '){
							i++
						}
					}
				
				if(openedTag == "NONE"){					

				}
			}
			for _,char := range storyAsRuneArr[i:]{
				charAsString := string(char)
				if(charAsString == " "){

				}else{

				}
				tagTypeBuilder.WriteRune(char)


				// tag identifier(opening)
				if char == '>' || (storyAsRuneArr[rightPointer] == ' ' && len(strings.TrimSpace(tagTypeBuilder.String())) != 0 ) {
					//CHECK if tag type builder is some kind of tag, parse all the way through to until the ">", check if properties are ok
					// NO EMBEDDING RULE!!!
					if(len(openedTags) != 0){
						return errors.New(fmt.Sprintf("Embedded html detected(%s)",tagTypeBuilder.String()))
					}

					openedTags = append(openedTags,strings.Trim(tagTypeBuilder.String()))
					tagTypeBuilder.Reset()
				//tag identifier(closing)
				}else if storyAsRuneArr[rightPointer] == '/'{
					if len(openedTags) == 0{
						return errors.New(fmt.Sprintf("Incorrect closing tag at character %s",rightPointer))
					}
				}else{
					fmt.Println("xd")
				}
			}
		}else{
			i++
		}
	}

	if(openedTag != "NONE"){
		return errors.New(fmt.Sprintf("ERROR! Story is malformed: last opened tag is unclosed: %s",openedTag))
	}
	
	// how to deal with <div></div></div>
	return nil
}
*/