import { BACKEND_URL } from "../../../constants"


export function composeHtmlElement(newContent: string, elementType: 'title' | 'paragraph'){
    const tagType = elementType === 'paragraph' ? 'p' : 'h2'

    const resultingString = `<${tagType}>${newContent}</${tagType}>`
    return resultingString
}

export function createNewImageElement(identifier:string){
    return `<img src={${BACKEND_URL}/images/stories/${identifier}}'></img>`
}