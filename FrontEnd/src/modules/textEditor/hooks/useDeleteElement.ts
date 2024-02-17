import useGetState from "./useGetElementState"
import { sendChangedStoryToServer } from "./useTextEditor"

export default () => {
    const { elementOrderArray,elementMap,setElementMap } = useGetState()

    async function deleteBlock(blockIdentifier:string){
        const newElementMap = new Map(elementMap)
        if(newElementMap.delete(blockIdentifier)){
            setElementMap(newElementMap)
            await sendChangedStoryToServer()
        }
    }

    return {deleteBlock}
}