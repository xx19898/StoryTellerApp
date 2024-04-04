import {Provider, atom, useAtom} from 'jotai'
import { myStore } from '../atomStore'
import { StoryObj } from '@storybook/react'

const counter = atom<number>(10)

const Counter = () => {
    const [counterValue,setCounterValue] = useAtom(counter)

    return(
        <div className="flex flex-col justify-center items-center gap-3">
        <p>Counter: {counterValue}</p>
        <button className="p-2 px-8 rounded-md bg-secondPrimary" onClick={() => setCounterValue(counterValue + 1)}>Plus</button>
        </div>
    )
}

export default{
    component:Counter ,
    title: 'Counter Story',
    decorators: [
        (CounterStory) => (
            <div className="min-h-screen h-auto w-full flex flex-col justify-center items-center bg-special">
                <Provider store={myStore}>
                    <CounterStory />
                </Provider>
            </div>
        )
    ]
}

export const Default: StoryObj<typeof Counter> = {
    args: {}
}



