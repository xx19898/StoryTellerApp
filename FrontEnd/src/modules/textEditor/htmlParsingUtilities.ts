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

export function processHtmlString(htmlString:string){
    let i = 0

    let closingBracket = 0

    const arr = []
    const map = new Map()

    let htmlElement = ''

    while(i < htmlString.length){
        if(htmlString[i] === '>') closingBracket++

        if(closingBracket === 2){
            htmlElement = htmlElement + htmlString[i]
            closingBracket = 0
            const {element} = extractTypeAndContentOfHtmlElement(htmlElement)
            const newIdentifier = uuidv4()
            arr.push(newIdentifier)
            map.set(newIdentifier,element)
            htmlElement = ''
        }else{
            htmlElement = htmlElement + htmlString[i]
        }

        i++
    }

    return {htmlOrderArray:arr,htmlElementMap:map}
}

export function addHtmlElementIdentifier( elType:'h2' | 'p', content: string){
    return `<${elType}>${content}</${elType}>`
}

function getNewIdentifierForElement(){
    return uuidv4()
}

export function addNewParagraph(newParagraphContent:string,map:Map<string,string>,arr:Array<string>){
    const newElement = addHtmlElementIdentifier('p',newParagraphContent)
    const newIdentifier = getNewIdentifierForElement()
    arr.push(newIdentifier)
    map.set(newIdentifier,newElement)
}

export function setNewTitle(newTitleContent:string,map:Map<string,string>,arr:Array<string>){
    const newTitleElement = addHtmlElementIdentifier('h2',newTitleContent)
    const newIdentifier = getNewIdentifierForElement()
    map.delete(arr[0])
    arr[0] = newIdentifier
    map.set(newIdentifier,newTitleElement)
}

export function buildHtmlString(map:Map<string,string>,arr: Array<string>){
    let htmlString = ''
    for(const identifier of arr){
        if(map.get(identifier) !== undefined){
            htmlString = htmlString + map.get(identifier)
        }
    }
    return htmlString
}


