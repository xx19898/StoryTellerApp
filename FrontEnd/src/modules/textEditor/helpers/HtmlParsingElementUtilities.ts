import { tagToType } from "./HtmlElementTagUtlities"

export function isElementType(elementType: string): elementType is ('title' | 'paragraph' | 'image'){
    return (elementType === 'paragraph' || elementType === 'title' || elementType === 'image')
}

export function extractTypeAndContentOfHtmlElement(htmlString:string): {
    element: string,
    elementType: 'title' | 'paragraph' | 'image',
    contents: string,
}{
    let i = 0
    while(htmlString[i] != '<') i++

    let j = i
    j++

    let elementType = ''

    while(htmlString[j] != '>' && htmlString[j] != ' '){
        elementType = elementType + htmlString[j]
        j++
    }

    elementType = tagToType(elementType)

    if(!isElementType(elementType)){
        throw new Error(`Incorrect html element type: ${elementType}`)
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