package stories

import "errors"

func prelimCheckStory(story string) (bool, error) {
	if len(story) > 500 {
		return false, errors.New("Story too long")
	}
	if(len(story) == 0) {
		return false, errors.New("Story is empty")
	}
	return true, nil
}
func getTypeOfElement(element string) (string, error){
	
}
func SanitizeStoryHtmlString(unsntzd string) (string,error) {
	storyIsOk,err := prelimCheckStory(unsntzd)
	if(!storyIsOk || err != nil){
		return "",err
	} 
	
	let currElementType = ""
	//element type can only be h1,2,3..., p, img
	for index, char := range unsntzd {

	}
}

