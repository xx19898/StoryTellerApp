import { useState } from "react";
import ReactQuill from "react-quill";
import 'react-quill/dist/quill.snow.css';
import Story from "../story/Story";

     //TODO: 1) make function to parse html string and extract the <h1>title,
     export const TextEditor = () => {
        const [story, setStory] = useState('')

        return(
            <div className="w-screen min-h-screen h-auto bg-base flex flex-col gap-[3.5em] justify-center items-center">

            </div>);
    }

