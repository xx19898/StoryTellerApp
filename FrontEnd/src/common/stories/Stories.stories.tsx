import { StoryObj } from "@storybook/react";

import Story, { IStory } from "./Story";

export default{
    component: Story,
    title: 'Story',
    decorators: [
        (Story) => (
            <div className="min-h-screen h-auto w-auto flex flex-col justify-center items-center bg-white">
                <Story />
            </div>
        )
    ]

}

type TypeStory = StoryObj<typeof Story>

const defaultArgs: IStory = {
    content: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis fermentum posuere quam, id tempus elit tristique sed. Vivamus cursus ullamcorper interdum. In semper pulvinar dolor, ac sollicitudin ante congue ultrices. Cras posuere lobortis condimentum. Proin molestie, metus a laoreet fringilla, mauris ipsum viverra metus, vel molestie lectus est at neque. In vel elementum ex. Sed ante nisl, ultrices sit amet enim sit amet, venenatis vestibulum justo. Praesent sollicitudin auctor orci et lacinia. Curabitur luctus feugiat arcu, eget gravida orci ullamcorper eu. Ut volutpat, lacus eget dictum varius, massa enim tincidunt tellus, vitae finibus massa purus in libero. Praesent sapien quam.',
    ID: 3,
    owner: {
        email:'owner@gmail.com',
        name: 'owner',
    },
    title: 'First Story',
    comments: [
        {
            content: 'test comment 1',
            ID: 1,
            owner: {
                email:'owner1@gmail.com',
                name: 'owner1'
            },
        },
        {
            content: 'test comment 2',
            ID: 2,
            owner: {
                email: 'owner2@gmail.com',
                name: 'owner2',
            }
        }
    ]
}

export const Default: TypeStory = {
    args:defaultArgs
}