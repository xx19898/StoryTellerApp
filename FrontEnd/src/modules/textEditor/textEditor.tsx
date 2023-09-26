import { useState } from "react";
import ReactQuill from "react-quill";
import 'react-quill/dist/quill.snow.css';
import Story from "../story/Story";

     export const TextEditor = () => {
        const [story,setStory] = useState('')

        function onChange(newStory:string){
            console.log({newValue: newStory})
            setStory(newStory)
        }
        return(
            <div className="w-full h-screen bg-base flex flex-col gap-[3.5em] justify-center items-center">
                <ReactQuill
                onChange={onChange}
                className="w-[80%] bg-secondPrimary p-10 pb-20"

                modules={{
                    toolbar:
                    [
                        [{ header: [1, 2, false] }],
                        ['bold', 'italic', 'underline'],
                        [{ 'script': 'sub'}, { 'script': 'super' }],
                        [{ 'indent': '-1'}, { 'indent': '+1' }],
                        [{ 'direction': 'rtl' }],
                        ['align']
                    ]
                    }}
                />
                <div className="w-[80%] p-10">
                    <Story htmlString={story}/>
                </div>
            </div>);
    }

