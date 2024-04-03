import { deleteEmptyElements } from "../helpers/HtmlParsingUtilities"
import useGetState from "./useGetElementState"


function useRemoveEmptyElements(){
    const {
        elementMap,elementOrderArray,
        setElementMap,setElementOrderArray} = useGetState()

    function removeEmptyElements(){
        const newElementMap = new Map(elementMap)
        const newArr = [...elementOrderArray]

        const purifiedArr = deleteEmptyElements(newElementMap,newArr)
        setElementOrderArray(purifiedArr)
        setElementMap(newElementMap)
    }

    return {removeEmptyElements}
}


export default useRemoveEmptyElements