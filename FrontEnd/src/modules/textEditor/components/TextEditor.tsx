import useTextEditor from "../hooks/useTextEditor";
import StoryEditor from "./storyVisualisation/StoryEditor";



//TODO: start implementing images
     export const TextEditor = () => {
        useTextEditor()

        return(
            <div className="w-[40%] py-2 px-4 min-h-screen h-auto
            flex flex-col gap-[3.5em] justify-start items-center
            text-white">
                <StoryEditor />
            </div>
            );
    }

