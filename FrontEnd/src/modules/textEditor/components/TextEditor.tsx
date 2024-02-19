import useTextEditor from "../hooks/useTextEditor";
import CurrentStory from "./storyVisualisation/CurrentStory";



//TODO: start implementing images
     export const TextEditor = () => {
        useTextEditor()        
        
        return(
            <div className="w-auto max-w-[40%] py-2 px-4 min-h-screen h-auto
            flex flex-col gap-[3.5em] justify-start items-center
            text-white">
                <h1>Text editor</h1>
                <CurrentStory />
            </div>
            );
    }

