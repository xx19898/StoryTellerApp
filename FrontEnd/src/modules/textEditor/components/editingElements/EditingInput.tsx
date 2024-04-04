import { useCallback, useEffect, useRef, useState } from "react"
import { SlActionUndo } from "react-icons/sl";
import { SlCheck } from "react-icons/sl";
import _debounce from 'lodash.debounce'
import LoadingSpinner from "../LoadingSpinner";
import UseEditBlock from "../../hooks/useEditBlock";
import useEditBlock from "../../hooks/useEditBlock";

interface IEditingInput{
    origValue: string,
    identifier: string,
}

const EditingInput = ({identifier,origValue}:IEditingInput) => {
    const [inputValue,setInputValue] = useState<string>(origValue)
    const origValueRef = useRef(origValue)
    const [textAreaHeight,setTextAreaHeight] = useState<string | number>('auto')
    const textAreaRef = useRef<HTMLTextAreaElement>(null)
    const [storyUpdating,setStoryUpdating] = useState<boolean>(false)

    const {editBlock} = useEditBlock(identifier,textAreaRef)

    
    useEffect(() => {
        textAreaRef.current?.focus()
        const len = origValue.length
        textAreaRef.current?.setSelectionRange(len,len)
        updateTextAreaHeight()
    },[])



    function updateTextAreaHeight(){
        if(textAreaRef.current?.scrollHeight && textAreaRef.current?.style.height){
            console.log({currHeight:textAreaRef.current?.scrollHeight})
            setTextAreaHeight(textAreaRef.current?.scrollHeight)
        }
    }

    const debouncedMessage = useCallback(_debounce(async (newVal:string) => {
        setStoryUpdating(true)
        console.log({newVal,identifier})
        await editBlock(newVal,identifier)
        setStoryUpdating(false)
        updateTextAreaHeight()
    },1000),[editBlock])

    async function onChange(newVal:string){
        await debouncedMessage(newVal)
    }

    return(
        <div className="w-full h-auto flex flex-col justify-center items-center rounded-md outline outline-1 outline-white">
            {
                storyUpdating
                ?
                <LoadingSpinner />
                :
                null
            }
            <textarea
            spellCheck={false}
            className="indent-4 p-2 h-auto w-full text-white bg-secondary focus:outline-none rounded-md resize-none"
            defaultValue={inputValue}
            style={{height: textAreaHeight === 'auto' ? 'auto' : `${textAreaHeight}px`}}
            onChange={async (e) => await onChange(e.target.value)}
            ref={textAreaRef}
            >
            </textarea>
        </div>
    )
}

export default EditingInput