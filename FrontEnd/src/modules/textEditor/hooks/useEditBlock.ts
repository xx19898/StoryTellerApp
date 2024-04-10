import { sendChangedStoryToServer } from "./useTextEditor"
import { useEffect } from "react"
import { addHtmlElementIdentifier, deleteEmptyElements, processHtmlString} from "../helpers/HtmlParsingUtilities"
import useGetState from "./useGetElementState"
import UseSelectElement from "./useSelectElement"
import { extractTypeAndContentOfHtmlElement } from "../helpers/HtmlParsingElementUtilities"
import { typeToTag } from "../helpers/HtmlElementTagUtlities"
import useRemoveEmptyElements from "./useRemoveEmptyElements"


const useEditBlock = (editSectionRef:React.RefObject<HTMLDivElement>) => {
    const {
        elementMap,setElementMap,elementOrderArray,setElementOrderArray
    } = useGetState()

    const {selectElement} = UseSelectElement()
    const {removeEmptyElements} = useRemoveEmptyElements()

    //TODO: Write function in htmlparsingutility for deleting all the html elements whose contents are empty
    //TODO write custom hook for using that function on state and passing that function around connect that function to the onClickOutsideSectionElement
    //TODO connect that function to the onClickOutsideSectionElement
    useEffect(() => {
        document.addEventListener("mousedown", onClickOutsideEditSection)
        return () => {
            document.removeEventListener("mousedown", onClickOutsideEditSection)
        }
    },[elementMap,elementOrderArray])

    function onClickOutsideEditSection(e:MouseEvent){

        if (editSectionRef.current && !editSectionRef.current.contains(e.target as Node)) {
            console.log('%c CLICKED outside','color: red;')
            stopEditing()
            console.log({elementMapInOnClickOutside:elementMap})
            removeEmptyElements()
        }
    }

    function stopEditing(){
        selectElement(undefined)
    }

    async function editBlock(newContent:string,blockIdentifier:string){
        console.log('%c In the edit block','color: green;')
        console.log({elementMap})
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

export default useEditBlock