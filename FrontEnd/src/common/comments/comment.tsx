import { IUser } from "../users/user";


export interface IComment{
    content: string,
    ID: number,
    owner: IUser,
}

const Comment = ({content,owner}:IComment) => {
    return(
        <div className="w-[80%] h-max flex flex-row gap-10 justify-center align-center text-white p-2 bg-darkerSecondary rounded-md">
            <div className="flex flex-col justify-between items-center rounded-sm">
                <img
                src="https://variety.com/wp-content/uploads/2018/10/drive.jpg?w=1000&h=563&crop=1"
                className="w-10 h-10 rounded-[50%]"></img>
                <p className="text-center text-xl font-semibold">{owner.name}</p>
            </div>
            <p className="
            bg-darkestSecondary w-[80%] p-6
            flex flex-col justify-center items-start
            rounded-md shadow-[50px]">{content}</p>
        </div>
    )
}

export default Comment