import { useEffect, useRef, useState } from "react"
import { SlActionUndo } from "react-icons/sl";
import { SlCheck } from "react-icons/sl";


interface IEditingInput{
    origValue: string,
    edit: (val:string,identifier:string) => void,
    identifier: string,
    stopEditing: () => void, 
}




const EditingInput = ({identifier,edit,origValue,stopEditing}:IEditingInput) => {
    const [inputValue,setInputValue] = useState<string>(origValue)
    const [originalValue,setOriginalValue] = useState<string>(origValue)
    const textAreaRef = useRef<HTMLTextAreaElement>(null)
    const [currHeight,setCurrHeight] = useState(200)

    useEffect(() => {
        textAreaRef.current?.focus()
        let len = origValue.length
        textAreaRef.current?.setSelectionRange(len,len)
    },[])

    function onChange(newVal:string){
        edit(newVal,identifier)
        if(textAreaRef.current?.scrollHeight) setCurrHeight(textAreaRef.current?.scrollHeight)
    }

    return(
        <div className="w-full h-auto p-4 flex flex-col justify-center items-center bg-darkerSecondary rounded-md">
            <textarea style={{height:`${currHeight}px`}} className="indent-4 p-2 w-full h-full text-black focus:outline-none rounded-md" 
            defaultValue={inputValue} 
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