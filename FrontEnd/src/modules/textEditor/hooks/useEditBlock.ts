import { sendChangedStoryToServer } from "./useTextEditor"
import { useEffect } from "react"
import { addHtmlElementIdentifier, extractTypeAndContentOfHtmlElement, processHtmlString, typeToTag } from "../helpers/HtmlParsingUtilities"
import useGetState from "./useGetElementState"


const UseEditBlock = () => {
    const {
        elementMap,setElementMap,
        elementOrderArray,setElementOrderArray
    } = useGetState()

    
    //TO USE ONLY WHILE DEVELOPING WITH STORYBOOK, LATER IS REPLACED BY REACT QUERY
    useEffect(() => {
        const htmlString = '<h2>Title</h2><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi sollicitudin consequat condimentum. Suspendisse vitae libero et mi semper molestie. Suspendisse sed bibendum arcu. Suspendisse et aliquam tortor, eget sagittis lacus. Maecenas consectetur sollicitudin turpis, sed consequat felis mollis at. Nunc nec lectus condimentum, ultrices eros ut, auctor eros. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Fusce a sapien pharetra, pulvinar nibh ac, vestibulum lorem.</p><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam a velit lacinia, varius lorem id, euismod massa. Integer ornare varius congue. Pellentesque congue nulla quis mauris tincidunt, vel consectetur lorem.</p>'
        
        const {
            htmlElementMap,htmlOrderArray  
        } = processHtmlString(htmlString)
    
        setElementOrderArray(htmlOrderArray)
        setElementMap(htmlElementMap)
    },[])
    

    async function editBlock(newContent:string,blockIdentifier:string){
        const newElementMap = new Map(elementMap)
        const element = elementMap.get(blockIdentifier)
        if(element != undefined){
            const {elementType} = extractTypeAndContentOfHtmlElement(element)
            const tag = typeToTag(elementType)
            if(tag != 'unknown'){
                if(tag === 'img'){
                    console.log(`SETTING NEW PICTURE ON ${blockIdentifier}`)
                    /* 
                    HERE SHOULD BE FUNCTIONALITY FOR 
                    1) UPDATE FILE UNDER SAME IDENTIFIER NAME ON THE BACKEND SIDE WITH REACT QUERY
                    */    
                }else{
                    const newElement = addHtmlElementIdentifier(tag, newContent) 
                    newElementMap.set(blockIdentifier, newElement)
                    setElementMap(newElementMap)
                } 
            }
        }

        await sendChangedStoryToServer()
    }

    return {editBlock}
}
