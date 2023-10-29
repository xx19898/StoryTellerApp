
import { StoryObj } from "@storybook/react";
import LoginPage from "./loginPage";
import Story from "../../modules/story/Story";

export default {
    component: LoginPage,
    title: 'Login Page'
}

type Story = StoryObj<typeof LoginPage>

export const Default: Story = {
    args: {
        login: (username:string, password:string) => console.log({username,password}),
        error: 'Test Error',
        isLoading: true,
        loginSuccess: false,
        onSuccess: () => console.log('success')
    },
}

export const Success: Story = {
    args: {
        login: (username:string, password:string) => console.log({username,password}),
        error: undefined,
        isLoading: false,
        loginSuccess: true,
        onSuccess: () => console.log('success')
    },
}