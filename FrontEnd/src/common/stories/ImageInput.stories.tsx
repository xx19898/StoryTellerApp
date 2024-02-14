import {StoryObj} from "@storybook/react";
import AddNewImage from "../../modules/textEditor/addNewImage";

export default{
    component: AddNewImage,
    title: 'New image input',
    decorators: [
        (NewImageInput) => (
            <div className="min-h-screen h-auto w-full flex flex-col justify-center items-center">
                <NewImageInput />
            </div>)
    ]
}

type TypeNewStory = StoryObj<typeof AddNewImage>

const defaultArgs = {
}


export const Default: TypeNewStory = {
    args:defaultArgs,
}
