import { TextEditor } from "../../modules/textEditor/textEditor"

export interface INewStory{
    username: string,
}

const NewStory = ({username}:INewStory) => {
    return(
        <div>
            <h1>Write a new Story</h1>
            <section className="w-[80%] py-5">
                <TextEditor />
            </section>
        </div>
    )
}

export default NewStory