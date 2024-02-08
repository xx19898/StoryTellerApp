import { useState } from "react";
import ReactQuill from "react-quill";
import 'react-quill/dist/quill.snow.css';
import Story from "../story/Story";
import EditingBlock from "./EditingBlock";


     export const TextEditor = () => {
        const [story, setStory] = useState('')

        const arr = [
            {elementType: 'title', contents: 'Title'},
            {elementType: 'paragraph', contents:'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi sollicitudin consequat condimentum. Suspendisse vitae libero et mi semper molestie. Suspendisse sed bibendum arcu. Suspendisse et aliquam tortor, eget sagittis lacus. Maecenas consectetur sollicitudin turpis, sed consequat felis mollis at. Nunc nec lectus condimentum, ultrices eros ut, auctor eros. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Fusce a sapien pharetra, pulvinar nibh ac, vestibulum lorem.'}
        ]

        return(
            <div className="w-auto max-w-[60%] py-2 px-4 min-h-screen h-auto
            flex flex-col gap-[3.5em] justify-start items-center
            text-white">
                <section className="w-3/4 min-h-screen flex flex-col justify-start items-center">
                {
                    arr.map((el) => {
                        if(el.elementType === 'title') return <EditingBlock content={el.contents} type={el.elementType}/>
                        else if(el.elementType === 'paragraph') return <EditingBlock content={el.contents} type={el.elementType}/>
                })
                }
                </section>
            </div>);
    }

