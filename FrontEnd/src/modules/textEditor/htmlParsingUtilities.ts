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
    while(htmlString[i] <= '<') i++

    let j = i

    let elementType = ''

    while(htmlString[j] != '>'){
        elementType = elementType + htmlString[j]
        j++
    }
    let theElement = ''
    while(htmlString[i] != '>'){
        theElement = theElement + htmlString[i]
        i++
    }
    theElement = theElement + htmlString[i]

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

