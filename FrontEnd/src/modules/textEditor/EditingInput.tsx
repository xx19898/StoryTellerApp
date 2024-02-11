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

    useEffect(() => {
        textAreaRef.current?.focus()
        let len = origValue.length
        textAreaRef.current?.setSelectionRange(len,len)
    },[])

    function onChange(newVal:string){
        edit(newVal,identifier)
    }

    return(
        <div className="w-full flex flex-col justify-center items-center bg-darkerSecondary rounded-md p-3">
            <textarea className="indent-4 p-4 w-full h-auto text-black focus:outline-none rounded-md" 
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