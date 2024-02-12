import { useEffect, useRef, useState } from "react"
import { SlActionUndo } from "react-icons/sl";
import { SlCheck } from "react-icons/sl";

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
        <div className="w-full h-auto p-4 flex flex-col justify-center items-center bg-darkerSecondary rounded-md">
            {
                storyUpdating ? <p>UPDATING</p>: null
            }
            <textarea 
            className="indent-4 p-2 h-auto w-full text-black focus:outline-none rounded-md" 
            defaultValue={inputValue}
            style={{height: textAreaHeight === 'auto' ? 'auto' : `${textAreaHeight}px`}} 
            onChange={(e) => onChange(e.target.value)}
            ref={textAreaRef}
            >
            </textarea>
            <section className="flex flex-row gap-5 pt-3 justify-center items-center">
                <SlActionUndo size='1.5rem'/>
                <SlCheck size='1.5rem'/>
            </section>
        </div>
    )
}

export default EditingInput