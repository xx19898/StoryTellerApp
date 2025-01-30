
export function swapBlocks(firstBlockIdentifier:string,secondBlockIdentifier:string,elementOrderArray: string[]){
    const modifiedOrderArray = [...elementOrderArray]
    
    const firstElementIndex = modifiedOrderArray.findIndex((blockIdentifier) => blockIdentifier === firstBlockIdentifier)
    const secondElementIndex = modifiedOrderArray.findIndex((blockIdentifier) => blockIdentifier === secondBlockIdentifier)

    modifiedOrderArray[firstElementIndex] = elementOrderArray[secondElementIndex]
    modifiedOrderArray[secondElementIndex] = elementOrderArray[firstElementIndex]

    return modifiedOrderArray
}

export function