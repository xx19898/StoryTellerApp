import Comment, { IComment } from "../comments/comment"
import { IUser } from "../users/user"

export interface IStory{
    title: string,
    ID: number,
    content: string,
    owner: IUser,
    comments: IComment[],
}

const Story = ({ID,comments,content,owner,title}:IStory) => {
    return(
        <div className="bg-secondary text-white flex flex-col justify-between items-center font-belanosima">
            <section>
                <h2 className="text-xl font-semibold text-center my-5">
                    {title}
                </h2>
                <section className="px-5">
                    {content}
                </section>
                <div className="w-auto mx-5 h-1 bg-darkestSecondary rounded-[40px] my-5"></div>
                    <ul className="flex flex-col items-center justify-between gap-5 pb-5">
                        {
                            comments.map((comment) =>
                            <Comment
                            ID={comment.ID}
                            content={comment.content}
                            owner={comment.owner}
                            key={comment.ID} />)
                        }
                    </ul>
            </section>
        </div>
    )
}

export default Story