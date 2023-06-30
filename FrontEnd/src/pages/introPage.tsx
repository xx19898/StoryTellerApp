
import { gsap } from "../gsap";
import { useEffect, useRef, useState } from "react"
import {
    Transition,
} from 'react-transition-group';
import { Elastic } from "gsap/gsap-core";
import { useNavigate, useRoutes } from "react-router-dom";


const Welcome = () => {
    
    useEffect(() => {
        setPageVisible(true)
    },[])
    const mainRef = useRef(null)
    const headerRef = useRef(null)
    const buttonRef = useRef(null)
    const [pageVisible,setPageVisible] = useState(false)
    const navigate = useNavigate()
    const [actionHovered,setActionHovered] = useState(false)
    
    return(
        <>
        <Transition timeout={400} in={pageVisible} nodeRef={mainRef} unmountOnExit mountOnEnter onEntered={() => console.log('ENTERED')} onEnter={() => gsap.from(mainRef.current,{autoAlpha:0,x:"-100vw",stagger:0.3})} onExit={() => gsap.to(mainRef.current,{autoAlpha:0,x:"100vw"})}>
                <div  className="w-full min-h-screen bg-base flex flex-col justify-center items-center font-belanosima">
                    <div ref={mainRef} className="bg-secondary flex flex-col justify-center items-center card mx-10 py-[1em] shadow-3xl">
                        <h1 ref={headerRef} className="card-title text-4xl font-normal text-white my-4">Hello, <u className="text-primary decoration-transparent">stranger</u></h1>
                        <p className="text-white text-xl mx-10 mb-4">Do you have a <u className="text-primary decoration-transparent">story</u> to <u className="text-primary decoration-transparent">tell?</u> Or maybe you would be interested to read the stories of others then?
                        Either way, we've got you <u className="text-primary decoration-transparent">covered</u>. Sign in, or sign up and get going! </p>
                        //TODO: animate hover of call to action 
                        <button ref={buttonRef} onMouseEnter={() => setActionHovered(true)} className="btn-primary rounded-md px-8 py-4 text-lg mb-4">Get going!</button>
                    </div>
                </div>                
        </Transition>
        </>

    )
}

export default Welcome