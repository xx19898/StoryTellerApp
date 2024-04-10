import {StoryObj} from "@storybook/react";
import NewStory, { INewStory } from "./NewStory";
import { Provider } from "jotai";
import { myStore } from "../../atomStore";
import StoryEditor from "../../modules/textEditor/components/storyVisualisation/StoryEditor";
import { TextEditor } from "../../modules/textEditor/components/TextEditor";

export default{
    component: StoryEditor,
    title: 'Story Editor',
    decorators: [
        (StoryEditor) => (

            <div className="min-h-screen h-auto w-full flex flex-col justify-center items-center bg-secondary">
                <Provider store={myStore}>
                    <TextEditor />
                </Provider>

            </div>)
    ]
}

type TypeNewStory = StoryObj<typeof StoryEditor>

const defaultArgs: INewStory = {
    username: 'test user',
}


export const Default: TypeNewStory = {
    args:defaultArgs,
}

