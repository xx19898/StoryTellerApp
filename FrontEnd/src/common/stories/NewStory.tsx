import { TextEditor } from "../../modules/textEditor/textEditor"

export interface INewStory{
    username: string,
}

const NewStory = ({username}:INewStory) => {
    return(
        <div className="w-full min-h-screen bg-secondary h-auto text-white flex flex-col items-center justify-center">
            <h1 className="mx-auto text-center text-xl pt-5 font-bold ">Write a new Story</h1>
            <TextEditor />
        </div>
    )
}

export default NewStory