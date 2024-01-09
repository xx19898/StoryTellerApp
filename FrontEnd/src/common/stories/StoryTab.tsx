import { LiaComments } from "react-icons/lia";
import { FaHeart } from "react-icons/fa";
import { v4 as uuidv4 } from 'uuid';


export interface IStoryTab{
    ID:number,
    Owner: string,
    NumberComments: number,
    Title: string,
    NumberOfLikes: number,
}

const StoryTab = ({Owner,Title,NumberComments,NumberOfLikes}:IStoryTab) => {
    return (
        <li key={uuidv4()} onClick={(e) => console.log('clicked')} className="bg-special hover:cursor-pointer w-full h-full flex flex-col gap-5 text-white justify-center align-center p-4 rounded-lg my-2">
            <p className='font-semibold text-xl text-center'>{Title}</p>
            <div className="w-full h-full flex flex-row gap-6 text-white justify-center align-center">
                <div className="flex flex-row justify-center align-center gap-2">
                    <FaHeart size={30}/>
                    <p>{NumberOfLikes}</p>
                </div>
                <div className="flex flex-row justify-center align-center gap-2">
                    <LiaComments size={30}/>
                    <p>{NumberComments}</p>
                </div>
                <div className="flex flex-row text-center align-center justify-center gap-2 ml-4">
                    <p >Author:</p>
                    <p className="font-semibold text-md relative text-center">{Owner}</p>
                </div>
            </div>
        </li>
    )
}

export default StoryTab