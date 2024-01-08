import { StoryObj } from "@storybook/react";
import Story from "../../modules/story/Story";
import Stories, { IStories, IStoryTabs } from "./StoryTabs";
import { IStory } from "./Story";

export default{
    component: Stories,
    title: 'Stories'
}

type Story = StoryObj<typeof Stories>

const defaultArgs: IStoryTabs = {
    storyTabs:[
        {
            ID: 0,
            Owner: 'User 1',
            Title: "Story 1",
            NumberComments: 1,
            NumberOfLikes: 10,
        },
        {
            ID: 1,
            Owner: 'User 1',
            Title: "Story 2",
            NumberComments: 10,
            NumberOfLikes: 342,
        }
    ]
}

export const Default: Story = {
    args: defaultArgs,
}