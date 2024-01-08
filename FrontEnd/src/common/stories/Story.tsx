export interface IStory{
    Title:string,
    ID: number,
    Content: string,
    Owner: string,
}

const Story = (storyParams:IStory) => {
    return(
        <li>
            <h2>{storyParams.Title}</h2>
        </li>
    )
}

export default Story