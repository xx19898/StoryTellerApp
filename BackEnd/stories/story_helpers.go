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
		"h1":[]string{},
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
		return "",errors.New("error. Index is out of boundaries")
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
		if char == '/' || char == '>' || char == ' ' || char == '<' {
			break
		}
		//fmt.Println(string(char))
		charSeq.WriteRune(char)
		if *index == len(story) - 1{
			break
		}
		*index++
	}
	//leaves index at last char of the grabbed sequence + 1
	return charSeq.String(),nil
}

//untested for now
func ScrollUntilNextOpeningBracket(story []rune, index *int)(error){
	for{
		if(*index >= len(story)){
			return errors.New("Got out of storys boundaries")
		}
		if *index == len(story) - 1{
			return errors.New("Reached story boundary")
		}
		if story[*index] == '<'{
			return nil
		}
		*index++
	}
}

func ParseProperties(index *int,story []rune,tag string) (map[string]string,error){
	propertiesMap := make(map[string]string)
	//fmt.Println(string(story))
	for{
		charSeq,err := GrabNextCharSeq(story,index)
		//fmt.Println("Last index " + string(story[*index]))
		if len([]rune(charSeq)) == 0{
			//fmt.Println(propertiesMap)
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
		//add here origin testing for img tags
		//can there be multiple attributes with the same name?
		propertiesMap[attribute] = value
	}

	return propertiesMap,nil
}

func ControlProperties(propertiesMap map[string]string, tag string) (error) {
	
	switch tag{
		case "img":
			if len(propertiesMap) == 0 || propertiesMap["src"] == ""{
				return errors.New("did not register obligatory src attribute on img tag")
			}
			
		default:
			//no attributes allowed
			if len(propertiesMap) != 0{
				return fmt.Errorf("error when controlling control properties on an element: tag is %s and it has unallowed properties %s",tag,propertiesMap)
			}		
	}

	return nil
}

func OnOpeningBracketEncountered(currIndex *int, story []rune,openedTag *string)(error){
	//currIndex at whatever comes next after <
	*currIndex++
	if *currIndex >= len(story){
		return fmt.Errorf("unclosed < at index %s(end of the story)",strconv.Itoa(*currIndex))
	}
	// openedTag == NONE => has to be legit opening tag (if not img tag)
	// tag name has to be after opening bracket
	// if img => src property HAS to be in place if not => > should follow the opening tag
	
	// openedTag != NONE => does not have to be legit opening/closing tag => 
	// in case of error IF tag is not enclosed with > just return the last position
	// => scroll until legit closing tag is found
	scrollToFirstNonSpaceChar(currIndex,story)

	// Checks whether tag is a closing one
	if(story[*currIndex] == '/'){
		if *openedTag == "NONE"{
			return fmt.Errorf("improper opening tag synthaxis at index %s",strconv.Itoa(*currIndex))
		}
		*currIndex++
		scrollToFirstNonSpaceChar(currIndex,story)
		//reading closing tag name
		closingTagName,err := GrabNextCharSeq(story,currIndex)
		if err != nil{
			return err
		}
		
		scrollToFirstNonSpaceChar(currIndex,story)
		if *currIndex >= len(story){
			return nil
		}
		if story[*currIndex] != '>'{
			return fmt.Errorf("error at index %s: encountered unexpected char, should be \">\" as the tag is marked to be a closing tag, but got %s instead (closing tag is: %s)", strconv.Itoa(*currIndex), string(story[*currIndex]), closingTagName)
		}

		if *openedTag != closingTagName{
			return fmt.Errorf("error at index %s: closing tag(%s) does not have same tag name as the opened one currently(%s)",strconv.Itoa(*currIndex),closingTagName,*openedTag)
		}
		//closing tag
		*openedTag = "NONE"
		return nil
	}
	
	if *openedTag != "NONE"{
		return fmt.Errorf("error at index %s, there are two embedded tags. new tag opens after the last  one (%s) was not closed",strconv.Itoa(*currIndex),*openedTag)
	}
	// TAG IS NOT A CLOSING TAG

	// reads tag
	tagName,err := GrabNextCharSeq(story,currIndex)

	if err != nil{
		return err
	}

	if len(tagName) == 0{
		return fmt.Errorf("error at char %s, there is no tag name, but opening and closing brackets",strconv.Itoa(*currIndex))
	}

	htmlTagIsCorrect,err := htmlTagIsAllowed(tagName)
	if !htmlTagIsCorrect || err != nil{
		return fmt.Errorf("error at char %s, improper tag name:%s",strconv.Itoa(*currIndex),tagName)
	}

	
	scrollToFirstNonSpaceChar(currIndex,story)
	propertiesMap,err := ParseProperties(currIndex, story, tagName)
	
	if err != nil{
		return fmt.Errorf("error when trying to parse html attributes: %s",err.Error())
	}
	err = ControlProperties(propertiesMap,tagName)
	if err != nil{
		return fmt.Errorf("error when trying to control that html attributes <are correct. Tag name is %s and error is %s ",tagName,err.Error())
	}

	if story[*currIndex] != '/' && story[*currIndex] != '>'{
		for{
			
			if story[*currIndex] != '/' && story[*currIndex] != '>'{
				break
			}
		}
	}


	//now index is behind the tag and all the properties

	// handles end of the tag
	switch tagType := tagName; tagType{
		case "img":
			if story[*currIndex] != '/'{
				return fmt.Errorf("error at char %s, img element missing slash",strconv.Itoa(*currIndex))
			}
			scrollToFirstNonSpaceChar(currIndex,story)
			if story[*currIndex] != '>'{
				return fmt.Errorf("error at char %s, img element in incorrect form, should encounter> next, but got %s instead",strconv.Itoa(*currIndex),string(story[*currIndex]))
			}
			return nil 
		default:
			scrollToFirstNonSpaceChar(currIndex,story)
			
			if story[*currIndex] != '>'{
				return fmt.Errorf("error at char %s, html element is in incorrect form, should encounter > next, but got %s instead",strconv.Itoa(*currIndex),string(story[*currIndex]))
			}
			*openedTag = tagType 
			return nil		
	}
}

func CheckStory(story string)(error){
	var err error
	storyAsRuneArr := []rune(story)
	openedTag := "NONE"
	index := 0

	for{
		if(index >= len(storyAsRuneArr)){
			return errors.New("reached the story boundary")
		}
		if(index == len(storyAsRuneArr) - 1){
			return nil
		}
		err = ScrollUntilNextOpeningBracket(storyAsRuneArr,&index)
		if err != nil{
			return err
		}
		err = OnOpeningBracketEncountered(&index,[]rune(story),&openedTag)
		if err != nil{
			return err
		}
	}
}