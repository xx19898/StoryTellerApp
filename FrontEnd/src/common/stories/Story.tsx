import { IUser } from "../users/user"

export interface IStory{
    Title:string,
    ID: number,
    Content: string,
    Owner: IUser,
}

const Story = (storyParams:IStory) => {
    return(
        <div>
            <h2>{storyParams.Title}</h2>
        </div>
    )
}

export default Story