import { v4 as uuidv4 } from 'uuid';

type IHtmlStringParsed = {
    htmlStringMap: Map<string,string>,
    htmlStringArray: Array<string>
}

export function extractTypeAndContentOfHtmlElement(htmlString:string): {
    element: string,
    elementType: string,
    contents: string
}{
    let i = 0
    while(htmlString[i] != '<') i++

    let j = i
    j++

    let elementType = ''

    while(htmlString[j] != '>'){
        elementType = elementType + htmlString[j]
        j++
    }

    let theElement = ''
    let closingBracket = 0

    while(closingBracket < 2){
        theElement = theElement + htmlString[i]
        if(htmlString[i] == '>') closingBracket++
        i++
    }

    let contents = ''

    i = 0

    while(htmlString[i] != '>') i++

    i++

    while(htmlString[i] != '<'){
        contents = contents + htmlString[i]
        i++
    }

    return {
        element: theElement,
        elementType: elementType,
        contents: contents,
    }
}

