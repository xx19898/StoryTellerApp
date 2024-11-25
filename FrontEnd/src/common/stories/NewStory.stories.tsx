<<<<<<< HEAD
import { StoryObj } from '@storybook/react'
import NewStory, { INewStory } from './NewStory'
import { Provider } from 'jotai'
import { myStore } from '../../atomStore'

export default {
	component: NewStory,
	title: 'NewStory',
	decorators: [
		(NewStory) => (
			<div className='min-h-screen h-auto w-full flex flex-col justify-center items-center bg-special'>
				<Provider store={myStore}>
					<NewStory />
				</Provider>
			</div>
		),
	],
=======
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
>>>>>>> 40c4d0b2a47243c9d95a8c25ba4ed35e49dcce72
}

type TypeNewStory = StoryObj<typeof StoryEditor>

const defaultArgs: INewStory = {
	username: 'test user',
}

export const Default: TypeNewStory = {
	args: defaultArgs,
}
