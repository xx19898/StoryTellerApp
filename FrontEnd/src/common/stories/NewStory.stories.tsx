import {StoryObj} from "@storybook/react";
import NewStory, { INewStory } from "./NewStory";

export default{
    component: NewStory,
    title: 'NewStory',
    decorators: [
        (NewStory) => (
            <div className="min-h-screen h-auto w-full flex flex-col justify-center items-center bg-special">
                <NewStory />
            </div>        )
    ]
}

type TypeNewStory = StoryObj<typeof NewStory>

const defaultArgs: INewStory = {
    username: 'test user',
}


export const Default: TypeNewStory = {
    args:defaultArgs,
}

