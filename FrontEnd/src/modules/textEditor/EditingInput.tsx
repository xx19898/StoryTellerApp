import { useEffect, useRef, useState } from "react"
import { SlActionUndo } from "react-icons/sl";
import { SlCheck } from "react-icons/sl";
import LoadingSpinner from "./LoadingSpinner";

interface IEditingInput{
    origValue: string,
    edit: (val:string,identifier:string) => Promise<void>,
    identifier: string,
    stopEditing: () => void, 
}

const EditingInput = ({identifier,edit,origValue,stopEditing}:IEditingInput) => {
    const [inputValue,setInputValue] = useState<string>(origValue)
    const [originalValue,setOriginalValue] = useState<string>(origValue)
    const [textAreaHeight,setTextAreaHeight] = useState<string | number>('auto')
    const textAreaRef = useRef<HTMLTextAreaElement>(null)
    const [storyUpdating,setStoryUpdating] = useState<boolean>(false)

    useEffect(() => {
        textAreaRef.current?.focus()
        let len = origValue.length
        textAreaRef.current?.setSelectionRange(len,len)
        updateTextAreaHeight()
    },[])

    function updateTextAreaHeight(){
        if(textAreaRef.current?.scrollHeight && textAreaRef.current?.style.height){
            console.log({currHeight:textAreaRef.current?.scrollHeight})
            setTextAreaHeight(textAreaRef.current?.scrollHeight)
        }
    }

    async function onChange(newVal:string){
        setStoryUpdating(true)
        await edit(newVal,identifier)
        console.log('got here')
        setStoryUpdating(false)
        updateTextAreaHeight()
    }

    return(
        <div className="w-full h-auto flex flex-col justify-center items-center rounded-md outline outline-1 outline-white">
            {
                storyUpdating ? 
                <LoadingSpinner />
                :
                null
            }
            <textarea 
            className="indent-4 p-2 h-auto w-full text-white bg-secondary focus:outline-none rounded-md resize-none" 
            defaultValue={inputValue}
            style={{height: textAreaHeight === 'auto' ? 'auto' : `${textAreaHeight}px`}} 
            onChange={(e) => onChange(e.target.value)}
            ref={textAreaRef}
            >
            </textarea>
        </div>
    )
}

export default EditingInput