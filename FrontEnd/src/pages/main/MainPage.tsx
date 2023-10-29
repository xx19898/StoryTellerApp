
export interface IMainPage{
    stories: {
        name:string,
        id: number,
    }[],

    loggedInUser: string,
}

const MainPage = ({stories, loggedInUser}:IMainPage) => {

    console.log({stories,loggedInUser})

    return(
        <div className="
        bg-secondary min-h-screen h-auto
        w-full flex flex-col justify-center items-center
        sm:px-4">
            <ul className="
            bg-transparent outline outline-3 outline-darkestSecondary
            rounded-md w-full flex sm:flex-col sm:justify-center sm:items-center
            sm:text-4xl sm:px-2 sm:mx-2 sm:py-4">
            {
                stories.map(
                    (story) =>
                    <li className="
                    text-white my-2 bg-darkerSecondary w-full text-center
                    rounded-md py-3 drop-shadow-md">
                        {story.name}
                    </li>)
            }
            </ul>
        </div>
    )
}

export default MainPage