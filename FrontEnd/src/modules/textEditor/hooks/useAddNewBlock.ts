import { addNewImage, addNewParagraph } from "../helpers/HtmlParsingUtilities"
import useGetState from "./useGetElementState"
import UseSelectElement from "./useSelectElement"

const UseAddNewBlock = () => {
    const {elementMap,elementOrderArray,setElementMap,setElementOrderArray} = useGetState()
    const {currentlyEditedElement,selectElement} = UseSelectElement()

    function AddNewTextBlock(){
        const newElMap = new Map(elementMap)
        const newElArray = [...elementOrderArray]

        addNewParagraph("",newElMap,newElArray)
        setElementMap(newElMap)
        setElementOrderArray(newElArray)

        const lastIdentifier = newElArray.slice(-1)[0]
        selectElement(lastIdentifier)
    }

    function AddNewImageBlock(){
        const newElMap = new Map(elementMap)
        const newElArray = [...elementOrderArray]

        addNewImage("",newElMap,newElArray)
        setElementMap(newElMap)
        setElementOrderArray(newElArray)
    }
}