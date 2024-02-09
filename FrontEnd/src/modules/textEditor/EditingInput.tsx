import { useState } from "react"
import { SlActionUndo } from "react-icons/sl";
import { SlCheck } from "react-icons/sl";


interface IEditingInput{
    value: string,
    setValue: (val:string) => void,
}

const EditingInput = () => {
    const [inputValue,setInputValue] = useState<string>()
    const [originalValue,setOriginalValue] = useState<string>()

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