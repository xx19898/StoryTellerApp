import {StoryObj} from "@storybook/react";
import NewStory, { INewStory } from "./NewStory";

export default{
    component: NewStory,
    title: 'NewStory',
    decorators: [
        (NewStory) => (
            <div className="min-h-screen bg-base h-auto w-auto flex flex-col justify-center items-center">
                <NewStory />
            </div>
        )
    ]
}

type TypeNewStory = StoryObj<typeof NewStory>

const defaultArgs: INewStory = {
    username: 'test user',
}


export const Default: TypeNewStory = {
    args:defaultArgs,
}

