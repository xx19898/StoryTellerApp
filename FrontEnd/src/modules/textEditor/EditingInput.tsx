import { useEffect, useState } from "react"
import { SlActionUndo } from "react-icons/sl";
import { SlCheck } from "react-icons/sl";


interface IEditingInput{
    origValue: string,
    setValue: (val:string,identifier:string) => void,
    identifier: string,
    stopEditing: () => void, 
}


const EditingInput = ({identifier,setValue,origValue,stopEditing}:IEditingInput) => {
    const [inputValue,setInputValue] = useState<string>()
    const [originalValue,setOriginalValue] = useState<string>()

    useEffect(() => {
        setOriginalValue(origValue)
    },[])

    function onChange(newVal:string){
        setInputValue(newVal)
        setValue(newVal,identifier)
    }

    //TODO: clicks dont get registered to the section element, fix

    return(
        <div className="w-full flex flex-col justify-center items-center bg-darkerSecondary rounded-md p-3">
            <input className="indent-4 p-2 w-full text-black focus:outline-none rounded-md" value={inputValue} onChange={(e) => setInputValue(e.target.value)}>
            </input>
            <section className="flex flex-row gap-5 pt-3 justify-center items-center">
                <SlActionUndo size='1.5rem'/>
                <SlCheck size='1.5rem'/>
            </section>
        </div>
    )
}

export default EditingInput