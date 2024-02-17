import { TextEditor } from "../../modules/textEditor/components/TextEditor"

export interface INewStory{
    username: string,
}

const NewStory = ({username}:INewStory) => {
    return(
        <div className="w-full min-h-screen bg-secondary h-auto text-white flex flex-col items-center justify-center">
            <TextEditor />
        </div>
    )
}

export default NewStory