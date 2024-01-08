
export interface IStoryTab{
    ID:number,
    Owner: string,
    NumberComments: number,
    Title: string,
    NumberOfLikes: number,
}

const StoryTab = ({Owner,Title,NumberComments,NumberOfLikes}:IStoryTab) => {
    return (
        <li key={Title} className="bg-special w-full h-full flex flex-row gap-5 text-white justify-center align-center p-4 rounded-lg my-2">
            <p>{Owner}</p>
            <p>{Title}</p>
            <p>{NumberOfLikes}</p>
        </li>
    )

}

export default StoryTab