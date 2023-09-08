import { useState } from "react";
import ReactQuill from "react-quill";
import 'react-quill/dist/quill.snow.css';

     export const TextEditor = () => {
        const [value,setValue] = useState('')
        console.log({value})
        function onChange(newval:string){
            console.log({newValue: newval})
            setValue(newval)
        }
        return(
            <div className="w-full h-screen bg-base flex flex-col justify-center items-center">
                <ReactQuill
                onChange={onChange}
                className="w-[80%] bg-secondPrimary"

                modules={{
                    toolbar:
                    [
                        [{ header: [1, 2, false] }],
                        ['bold', 'italic', 'underline'],
                        [{ 'script': 'sub'}, { 'script': 'super' }],
                        [{ 'indent': '-1'}, { 'indent': '+1' }],
                        [{ 'direction': 'rtl' }],
                        ['image', 'code-block'],
                    ]
                    }}/>
            </div>);
    }

