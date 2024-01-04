import gsap from "gsap"
import { useEffect, useRef } from "react"

export interface IMainPage{
    stories: {
        name:string,
        id: number,
    }[],

    onClickNewStoryBut: () => void,
    onClickMyProfileBut: () => void,
    onClickStoryBut:(storyId:number) => void,
    loggedInUser: string,
}

const MainPage = ({stories, loggedInUser, onClickNewStoryBut, onClickMyProfileBut, onClickStoryBut}:IMainPage) => {
    /*
    useEffect(() => {
        onEnter()
    },[])
    */
    const headerRef = useRef(null)
    const storiesListRef = useRef(null)
    const myAccountButRef = useRef(null)
    const createAStoryButRef = useRef(null)

    return(
        <div className="
        text-white
        bg-secondary min-h-screen h-auto
        w-full flex flex-col justify-center items-center
        sm:px-8">
            <h1 ref={headerRef} className="font-bold text-[48px] my-10">StoryTeller</h1>
            <button ref={myAccountButRef} className="sm:w-full bg-secondPrimary py-4 text-4xl rounded-md my-4">My Account</button>
            <button ref={createAStoryButRef} className="sm:w-full bg-darkerPrimary py-4 text-4xl rounded-md my-4">Write a story</button>
            <ul ref={storiesListRef} className="
            bg-transparent
            rounded-md w-full flex sm:flex-col sm:justify-center sm:items-center
            sm:text-4xl sm:py-4">
            {
                "STORIES"

                /*
                stories.map(
                    (story) =>
                    <li className="
                    text-white my-2 bg-darkerSecondary w-full text-center
                    rounded-md py-3 drop-shadow-md">
                        {story.name}
                    </li>)
                */
            }
            </ul>
        </div>
    )

    function onEnter(){
        const tl = gsap.timeline()
        tl
        .from([
            headerRef.current,
            storiesListRef.current,
            myAccountButRef.current,
            createAStoryButRef.current,
        ],
        {
            y: '-20vh',autoAlpha:0,stagger:0.1
        }
        ).play()
    }

    function onExit(navFunc: () => void){
        const tl = gsap.timeline()
        tl
        .from([
            createAStoryButRef.current,
            myAccountButRef.current,
            storiesListRef.current,
            headerRef.current,
        ],
        {
            y: '-20vh',autoAlpha:0,stagger:0.1
        }
        ).play()

        tl.eventCallback('onComplete',() => navFunc())
    }
}

export default MainPage