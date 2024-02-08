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
<<<<<<< HEAD
            <div className="w-screen min-h-screen h-auto bg-base flex flex-col gap-[3.5em] justify-center items-center">
                <button>Title</button>
            </div>);
    }


    /*
    <ReactQuill
                onChange={onChange}
                className="w-[700px] p-10"

                modules={{
                    toolbar:
                    [
                        [{ header: [1, 2, false] }], 
                        ['bold', 'italic', 'underline'],
                        [{ 'script': 'sub'}, { 'script': 'super' }],
                        [{ 'indent': '-1'}, { 'indent': '+1' }],
                        [{ 'direction': 'rtl' }],
                        ['align'],
                        ['image']
                    ]
                    }}
                />
                <div className="w-[80%] p-10">
                    <Story htmlString={story}/>
                </div>
                <p className="w-[] h-auto py-10">
                    {
                        story
                    }
                </p>
    */
=======
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
>>>>>>> dada76c4c000ac56fe0d3beffe648fbd29d66e6a

