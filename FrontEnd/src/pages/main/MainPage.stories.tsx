import { StoryObj } from "@storybook/react";
import MainPage, { IMainPage } from "./MainPage";


export default{
    component: MainPage,
    title: 'Main Page'
}

type Story = StoryObj<typeof MainPage>

const defaultArgs: IMainPage = {
    loggedInUser: 'User 1',
    stories: [
        {
            id: 1,
            name: 'Story 1'
        },
        {
            id: 2,
            name: 'Story 2',
        },
        {
            id: 3,
            name: 'Story 3',
        }
]
}

export const Default: Story = {
    args: defaultArgs,
}