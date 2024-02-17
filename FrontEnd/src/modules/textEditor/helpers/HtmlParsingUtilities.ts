import { v4 as uuidv4 } from 'uuid';
import { BACKEND_URL } from '../../../constants';

export function composeHtmlElement(newContent:string, elementType: 'title' | 'paragraph'){
    const tagType = elementType === 'paragraph' ? 'p' : 'h2'

    const resultingString = `<${tagType}>${newContent}</${tagType}>`
    return resultingString
}

export function createNewImageElement(identifier:string){
    return `<img src={${BACKEND_URL}/images/stories/${identifier}}'></img>`
}

export function tagToType(tag:string){
    if(tag === 'h2') return 'title'
    if(tag === 'p') return 'paragraph'
    if(tag === 'img') return 'image'
    return 'unknown'
}

export function typeToTag(type: 'title' | 'paragraph' | 'image'){
    if(type === 'title') return 'h2'
    if(type === 'paragraph') return 'p'
    if(type === 'image') return 'img'
    return 'unknown'
}

function isElementType(elementType: string): elementType is ('title' | 'paragraph'){
    return (elementType === 'paragraph' || elementType === 'title')
}

export function extractTypeAndContentOfHtmlElement(htmlString:string): {
    element: string,
    elementType: 'title' | 'paragraph' | 'image',
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

export function getNewIdentifierForElement(){
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


