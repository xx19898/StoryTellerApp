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
		return errors.New("Incorrect origin: " + urlOriginAsString)
	}

	return nil
}
func GetAllowedElementsAndPropertiesMap()([]string,map[string][]string){
	allowedElementTagsWithProperties := map[string] []string{
		"h":[]string{},
		"p":[]string{},
		"div":[]string{},
		"img":[]string{"src"},}	
	allowedElements := make([]string,len(allowedElementTagsWithProperties))

	i := 0
	for k := range allowedElementTagsWithProperties{
		allowedElements[i] = k
		i++
	}

	return allowedElements,allowedElementTagsWithProperties
}
func AttributeAllowedForElement(tag string, attribute string)(bool,error){
	_,allowedElementTagsWithProperties := GetAllowedElementsAndPropertiesMap()
	tagAllowed,err := htmlTagIsAllowed(tag)

	if err != nil{
		return false,err
	}
	if !tagAllowed{
		return false,fmt.Errorf("tag %s is not allowed",tag)
	}
	
	allowedAttributes := allowedElementTagsWithProperties[tag]
	
	if len(allowedAttributes) == 0{
		return false,nil
	}
	for _,allowedAttribute := range allowedAttributes{
		if attribute == allowedAttribute{
			return true,nil
		}
	}

	return false,nil
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
		}else{
			break
		}
	}
}

func ParseHtmlAttribute(htmlAttributeString string)(string,string,error){
	splitByEqSign := strings.Split(htmlAttributeString,"=")

	if len(splitByEqSign) != 2{
		return "","",fmt.Errorf("Malformed Html Attribute: %s",htmlAttributeString)
	}

	return splitByEqSign[0],strings.ReplaceAll(splitByEqSign[1],"\"",""),nil
}

func GrabNextCharSeq(story []rune,index *int)(string,error){
	if(*index >= len(story) || *index < 0){
		return "",errors.New("Error. Index is out of boundaries")
	}
	if(*index == len(story) - 1){
		return "",nil
	}
	var charSeq strings.Builder
	for{
		if *index >= len(story){
			break
		}
		char := story[*index]
		if char == '/' || char == '>' || char == ' ' {
			break
		}
		charSeq.WriteRune(char)
		if *index == len(story) - 1{
			break
		}
		*index++
	}
	return charSeq.String(),nil
}

func ParseProperties(index *int,story []rune,tag string) (map[string]string,error){
	propertiesMap := make(map[string]string)
	fmt.Println(string(story))
	for{
		charSeq,err := GrabNextCharSeq(story,index)
		fmt.Println("Last index " + string(story[*index]))
		if len([]rune(charSeq)) == 0{
			fmt.Println(propertiesMap)
			break
		}
		if err != nil{
			return propertiesMap,err
		}
		attribute,value,attributeParsingError := ParseHtmlAttribute(charSeq)
		if attributeParsingError != nil{
			return propertiesMap, attributeParsingError
		}
		attributeOk,err := AttributeAllowedForElement(tag,attribute)
		if !attributeOk{
			return propertiesMap, fmt.Errorf("%s is not a proper attribute for %s tag",attribute,tag)
		}
		//can there be multiple attributes with the same name?
		propertiesMap[attribute] = value
	}

	return propertiesMap,nil
}

func OnOpeningBracketEncountered(currIndex *int, story []rune,openedTag *string)(bool,error){
	//currIndex at whatever comes next after <
	*currIndex++
	if *currIndex >= len(story){
		return false,fmt.Errorf("unclosed < at index %s(end of the story)",strconv.Itoa(*currIndex))
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
			return false,fmt.Errorf("improper opening tag synthaxis at index %s",strconv.Itoa(*currIndex))
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
				return false,fmt.Errorf("error at index %s: closing tag does not have same tag name as the opened one currently(%s)",strconv.Itoa(*currIndex),*openedTag)
			}
			return false,nil
		}
	}

	if *openedTag != "NONE"{
		return false,fmt.Errorf("error at index %s, there are two embedded tags. new tag opens after the last  one (%s) was not closed",strconv.Itoa(*currIndex),*openedTag)
	}
	// TAG IS NOT A CLOSING TAG

	// reads tag
	for{
		//checks that boundary is kept before grabbing the char to evade error
		if(*currIndex >= len(story)){
			return false, fmt.Errorf("reached end of the story and did not encounter closing tag")
		}
		char := story[*currIndex]
		if(*currIndex == len(story) - 1){
			if(char) != '>'{
				return false,fmt.Errorf("reached the end of the story and last char is not >")
			}
			if len(tagNameBuilder.String()) == 0{
				return false, fmt.Errorf("error at char %s, there is no tag name, but opening and closing brackets",strconv.Itoa(*currIndex))
			}
			htmlTagIsCorrect,err := htmlTagIsAllowed(tagNameBuilder.String())
			if !htmlTagIsCorrect || err != nil{
				return false, fmt.Errorf("error at char %s, improper tag name:%s",strconv.Itoa(*currIndex),tagNameBuilder.String())
			}
			*openedTag = tagNameBuilder.String()
			return true,nil
		}
		if(char == ' ' || char == '>'){
			break
		}
		tagNameBuilder.WriteRune(char)
		*currIndex++
	}

	if len(tagNameBuilder.String()) == 0{
		return false,fmt.Errorf("error at char %s, there is no tag name, but opening and closing brackets",strconv.Itoa(*currIndex))
	}

	htmlTagIsCorrect,err := htmlTagIsAllowed(tagNameBuilder.String())
	if !htmlTagIsCorrect || err != nil{
		return false, fmt.Errorf("error at char %s, improper tag name:%s",strconv.Itoa(*currIndex),tagNameBuilder.String())
	}

	//propertiesMap := make(map[string]string)
	// parse the properties
	
	//TODO: parse all the properties until next char is > or /
	scrollToFirstNonSpaceChar(currIndex,story)

	if story[*currIndex] != '/' && story[*currIndex] != '>'{
		for{
			
			if story[*currIndex] != '/' && story[*currIndex] != '>'{
				break
			}
		}
	}


	switch tagType := tagNameBuilder.String(); tagType{
		case "img":
			if story[*currIndex] != '/'{
				return false,fmt.Errorf("error at char %s, img element missing slash",strconv.Itoa(*currIndex))
			}
			scrollToFirstNonSpaceChar(currIndex,story)
			if story[*currIndex] != '>'{
				return false,fmt.Errorf("error at char %s, img element in incorrect form, should encounter> next, but got %s instead",strconv.Itoa(*currIndex),string(story[*currIndex]))
			}
			return true,nil 
		default:
			scrollToFirstNonSpaceChar(currIndex,story)
			
			if story[*currIndex] != '>'{
				return false,fmt.Errorf("error at char %s, html element is in incorrect form, should encounter > next, but got %s instead",strconv.Itoa(*currIndex),string(story[*currIndex]))
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