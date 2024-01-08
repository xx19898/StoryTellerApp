import StoryTab, { IStoryTab } from "./StoryTab"

export interface IStoryTabs{
    storyTabs: IStoryTab[],
}

const Stories = ({storyTabs}:IStoryTabs) => {
    return(
        <ul className="bg-primary p-5 w-full">
        {
            storyTabs.map((storyTab) =>
            <StoryTab
            {...storyTab}
            />)
        }
        </ul>
    )
}

export default Stories