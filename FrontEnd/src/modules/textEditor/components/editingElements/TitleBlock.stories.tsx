import { StoryObj } from "@storybook/react";
import TitleBlock from "./TitleBlock";

export default{
    component:TitleBlock,
    title:'TitleBlock'
}

type Story = StoryObj<typeof TitleBlock>

const defaultArgs = {}

export const Default: Story = {
    args: defaultArgs,
}