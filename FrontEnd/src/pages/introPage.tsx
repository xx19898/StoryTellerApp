
import { gsap } from "../gsap";
import { useEffect, useRef, useState } from "react"

import {
    Transition,
} from 'react-transition-group';
import {FaRegArrowAltCircleUp,FaUserPlus} from "react-icons/fa"
import { IconContext } from "react-icons";
import { useNavigate} from "react-router-dom";


const Welcome = () => {

    useEffect(() => {
        setPageVisible(true)
    },[])

    const mainRef = useRef(null)
    const headerRef = useRef(null)
    const buttonRef = useRef(null)
    const loginButtonRef = useRef(null)
    const signUpButtonRef = useRef(null)
    const actionButtonRef = useRef(null)
    const buttonsRef = useRef(null)
    const registerButtonRef = useRef(null)

    const [nextPage,setNextPage] = useState<string>('')
    const [pageVisible,setPageVisible] = useState(false)
    const navigate = useNavigate()
    const [action,setAction] = useState(false)

    return(
        <>
        <Transition timeout={400} in={pageVisible} nodeRef={mainRef} unmountOnExit mountOnEnter
        onEnter={() => gsap.from(mainRef.current,{autoAlpha:0,x:"-100vw"})}
        onExit={() => handleExit()}>
                <div  className="w-full min-h-screen bg-base flex flex-col justify-center items-center font-belanosima">
                    <div ref={mainRef} className=" p-4 h-auto bg-secondary flex flex-col justify-center items-center rounded-md mx-10 py-[10%] shadow-3xl">
                        <h1 ref={headerRef} className="text-4xl font-normal text-white my-4">Hello, <u className="text-primary decoration-transparent">stranger</u></h1>
                        <p className="indent-4 text-white text-xl w-auto mb-4">Do you have a <u className="text-primary decoration-transparent">story</u> to <u className="text-primary decoration-transparent">tell?</u> Or maybe you would be interested to read the stories of others?
                        Either way, we've got you <u className="text-primary decoration-transparent">covered</u>. Sign in or sign up and get going! </p>
                        <button ref={actionButtonRef} onClick={() => setAction(true)}
                        className="btn-neutral rounded-md px-8 py-4 text-lg mb-4">
                            Get going!
                        </button>
                        <Transition timeout={500} in={action} nodeRef={buttonRef}  unmountOnExit mountOnEnter onEnter={() => clickedActionButton()}>
                        <IconContext.Provider value={{size:"2em",color:"white"}}>
                            <div className="flex flex-row h-auto bg-darkerSecondary w-[90%] rounded-md justify-between items-center text-white px-4 py-4" ref={buttonRef}>
                                <button ref={loginButtonRef} onClick={() => handleLoginClick()} className="bg-darkestSecondary rounded-md py-2 px-[3em] flex flex-col gap-1 justify-center items-center">
                                    Login<FaRegArrowAltCircleUp/>
                                </button>
                                <button ref={registerButtonRef} onClick={() => handleSignUpClick()} className="bg-darkestSecondary rounded-md py-2 px-[3em] flex flex-col justify-center gap-1 items-center">
                                    Sign Up<FaUserPlus/>
                                </button>
                            </div>
                        </IconContext.Provider>
                        </Transition>
                    </div>
                </div>
        </Transition>
        </>
    )
    function clickedActionButton(){
        console.log('xd')
        const timeline = gsap.timeline()
        timeline
        .to(actionButtonRef.current,{background:'#D81E5B',duration:0.5,color:"white"})
        .from(buttonRef.current,{autoAlpha:0,height:0,duration:0.5},0).play()
    }

    function handleLoginClick(){
        setNextPage('/login')
        setPageVisible(false)
    }
    function handleSignUpClick(){
        setNextPage('/signUp')
        setPageVisible(false)
    }
    function handleExit(){
        const timeline = gsap.timeline({onComplete:() => navigate(nextPage)})
        timeline.to(mainRef.current,{autoAlpha:0,x:"100vw"}).play()
    }
}

export default Welcome