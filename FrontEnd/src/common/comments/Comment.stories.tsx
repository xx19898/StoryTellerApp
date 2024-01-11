import { StoryObj } from "@storybook/react";
import Comment, { IComment } from "./comment";


export default{
    component: Comment,
    title: 'Comment',
    decorators: [
        (Comment) => (
            <div className="h-screen flex flex-col justify-center items-center bg-black">
                <Comment/>
            </div>
        )
    ]
}

type TypeComment = StoryObj<typeof Comment>

const defaultArgs: IComment = {
    content: 'test content',
    ID: 2,
    owner: {
        email: 'owner@gmail.com',
        name: 'owner',
    },
}

export const Default: TypeComment = {
    args: defaultArgs,
}

