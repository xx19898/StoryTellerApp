import { useEffect, useRef, useState } from "react"

const test = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut massa euismod, tristique nulla a, posuere nulla. Ut iaculis ultrices odio quis suscipit. In ac pretium neque, sit amet ultricies.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut massa euismod, tristique nulla a, posuere nulla. Ut iaculis ultrices odio quis suscipit. In ac pretium neque, sit amet ultricies.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut massa euismod, tristique nulla a, posuere nulla. Ut iaculis ultrices odio quis suscipit. In ac pretium neque, sit amet ultricies. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut massa euismod, tristique nulla a, posuere nulla. Ut iaculis ultrices odio quis suscipit. In ac pretium neque, sit amet ultricies.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut massa euismod, tristique nulla a, posuere nulla. Ut iaculis ultrices odio quis suscipit. In ac pretium neque, sit amet ultricies.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis ut massa euismod, tristique nulla a, posuere nulla. Ut iaculis ultrices odio quis suscipit. In ac pretium neque, sit amet ultricies.'

interface IResponsiveTextArea{
    onChange: (newContent:string) => Promise<void>,
    defaultContent: string,
    updating: boolean
}

//TODO: connect editing input functionality to responsive text area (onchange,updating)
const ResponsiveTextArea = ({onChange,defaultContent,updating}:IResponsiveTextArea) => {
    console.log({defaultContent})
    const [content, setContent] = useState(defaultContent)
    const [prevContent, setPrevContent] = useState(content)
    const [textAreaHeight, setTextAreaHeight] = useState<string | number>('auto')
    const textAreaRef = useRef<HTMLTextAreaElement>(null)

    useEffect(() => {
        textAreaRef.current?.focus()
        const len = defaultContent.length
        textAreaRef.current?.setSelectionRange(len,len)
        updateTextAreaHeight()
    },[])

    function updateTextAreaHeight(){
        if(content.length >= prevContent.length){
            if(textAreaRef.current?.scrollHeight && textAreaRef.current?.style.height){
                setTextAreaHeight('auto')
                setTextAreaHeight(textAreaRef.current?.scrollHeight + 'px')
            }
            return
        }

        if(textAreaRef.current?.scrollHeight && textAreaRef.current?.style.height){
            const newHeight = content.length / prevContent.length * textAreaRef.current?.scrollHeight + 'px'
            setTextAreaHeight(newHeight)
        }
    }

    return(
        <>
        <textarea
        onResize={() => updateTextAreaHeight()}
        spellCheck={false}
        defaultValue={content}
        ref={textAreaRef}
        className="indent-6 p-2 px-5 h-auto w-full text-white bg-secondary focus:outline-none rounded-md resize-none"
        onChange={async (e) => {
            setPrevContent(content)
            setContent(e.target.value)
            updateTextAreaHeight()
            await onChange(e.target.value)
        }}
        style={{height: textAreaHeight}}
        />
        </>

    )
}

export default ResponsiveTextArea