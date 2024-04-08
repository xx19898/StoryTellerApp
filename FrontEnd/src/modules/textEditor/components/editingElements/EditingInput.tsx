import { useCallback, useEffect, useRef, useState } from "react"
import { SlActionUndo } from "react-icons/sl";
import { SlCheck } from "react-icons/sl";
import _debounce from 'lodash.debounce'
import LoadingSpinner from "../LoadingSpinner";
import UseEditBlock from "../../hooks/useEditBlock";
import useEditBlock from "../../hooks/useEditBlock";
import ResponsiveTextArea from "./ResponsiveTextArea";

interface IEditingInput{
    origValue: string,
    identifier: string,
}

const EditingInput = ({identifier,origValue}:IEditingInput) => {
    const [inputValue,setInputValue] = useState<string>(origValue)
    const editingInputRef = useRef<HTMLDivElement>(null)
    const [storyUpdating,setStoryUpdating] = useState<boolean>(false)

    const {editBlock} = useEditBlock(editingInputRef)

    const debouncedMessage = useCallback(_debounce(async (newVal:string) => {
        setStoryUpdating(true)
        await editBlock(newVal,identifier)
        setInputValue(newVal)
        setStoryUpdating(false)
    },2000),[])

    async function onChange(newVal:string){
        await debouncedMessage(newVal)
    }

    return(
        <div ref={editingInputRef} className="w-full h-auto flex flex-col justify-center items-center rounded-md outline outline-1 outline-white">
            {
                storyUpdating
                ?
                <LoadingSpinner />
                :
                null
            }
            <ResponsiveTextArea defaultContent={origValue} onChange={onChange} updating={storyUpdating} />
        </div>
    )
}

export default EditingInput