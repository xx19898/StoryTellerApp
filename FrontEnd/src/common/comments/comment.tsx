import { IUser } from "../users/user";


export interface IComment{
    content: string,
    ID: number,
    owner: IUser,
}

const Comment = ({content,owner}:IComment) => {
    return(
        <div className="w-max h-max flex flex-row gap-10 justify-center align-center text-white p-2 bg-secSpecial">
            <div className="flex flex-col justify-center items-center bg-darkerSecondary">
                <img
                src="https://variety.com/wp-content/uploads/2018/10/drive.jpg?w=1000&h=563&crop=1"
                className="w-20 h-20 rounded-[50%]"></img>
                <p className="text-center text-xl p-4">{owner.name}</p>
            </div>
            <p className="bg-darkestPrimary p-6 flex flex-col justify-center align-center">{content}</p>
        </div>
    )
}

export default Comment