package stories

import (
	"errors"
	"strings"
)

func getAllowedElementsAndPropertiesMap()(map[string][]string,error){
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
func prelimCheckStory(story string) (bool, error) {
	if len(story) > 500 {
		return false, errors.New("Story too long")
	}
	if(len(story) == 0) {
		return false, errors.New("Story is empty")
	}
	return true, nil
}

func getTypeOfElement(element string,getElementsPropertiesMap elementsPropertiesFunctionType) (string, error){
	trimmedElement := strings.TrimSpace(element)
	  
	if(len(trimmedElement) == 0){
		return "",  errors.New("Element is empty")
	}
	
	allowedElementsAndPropertiesMap,err := getElementsPropertiesMap()
	if(err != nil){
		return "",errors.New("Could not get allowedElementsAndPropertiesMap")
	}
	//HERE
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
