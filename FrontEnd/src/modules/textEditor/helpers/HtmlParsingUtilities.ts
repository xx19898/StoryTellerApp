import { v4 as uuidv4 } from 'uuid';
import { BACKEND_URL } from '../../../constants';
import { extractTypeAndContentOfHtmlElement } from './HtmlParsingElementUtilities';

export function composeHtmlElement(newContent:string, elementType: 'title' | 'paragraph'){
    const tagType = elementType === 'paragraph' ? 'p' : 'h2'

    const resultingString = `<${tagType}>${newContent}</${tagType}>`
    return resultingString
}

export function createNewImageElement(identifier:string){
    return `<img src={${BACKEND_URL}/images/stories/${identifier}}'></img>`
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

export function addHtmlElementIdentifier( elType:'h2' | 'p' | 'img', content: string){
    return `<${elType}>${content}</${elType}>`
}

export function getNewIdentifierForElement(){
    return uuidv4()
}

export function addNewParagraph(newParagraphContent:string,map:Map<string,string>, arr:Array<string>){
    const newElement = addHtmlElementIdentifier('p',newParagraphContent)
    const newIdentifier = getNewIdentifierForElement()
    arr.push(newIdentifier)
    map.set(newIdentifier,newElement)
}

export function addNewImage(imageContent:string,map:Map<string,string>, arr:Array<string>){
    const newElement = addHtmlElementIdentifier('img',imageContent)
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


