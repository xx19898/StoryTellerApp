import { useEffect, useRef, useState } from "react"

function sleep(ms:number){
    return new Promise(resolve => setTimeout(resolve,ms))
}

interface ITitleBlock{
    setTitle: (newTitle:string) => Promise<void>,
}

const TitleBlock = () => {
    const [currVal,setCurrVal] = useState('Hello')
    const [textAreaHeight,setTextAreaHeight] = useState<string | number>('auto')
    const [edit,setEdit] = useState(false)
    const textAreaRef = useRef<HTMLTextAreaElement>(null)

    useEffect(() => {
        textAreaRef.current?.focus()
        const len = currVal.length
        textAreaRef.current?.setSelectionRange(len,len)
        updateTextAreaHeight()
    },[])

    function updateTextAreaHeight(){
        if(textAreaRef.current?.scrollHeight && textAreaRef.current?.style.height){
            console.log({currHeight:textAreaRef.current?.scrollHeight})
            setTextAreaHeight(textAreaRef.current?.scrollHeight)
        }
    }

    return(
        <div className="font-belanosima font-bold text-lg tracking-widest text-center 
        p-2 px-6 w-fit
        outline outline-solid 
        hover:cursor-pointer hover:outline-4"
            onClick={(e) => setEdit(true)}>
            {
                edit ? <textarea className="p-0 w-fit" defaultValue={currVal}></textarea> 
                : 
                currVal
            }
        </div>
    )
}

export default TitleBlock