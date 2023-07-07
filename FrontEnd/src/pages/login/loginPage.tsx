import { useEffect, useRef, useState } from "react"
import { gsap } from "gsap/src" 
import { Transition } from "react-transition-group"


const LoginPage = ({login}:{login:(username:string,password:string) => void}) => {
    
    useEffect(() => {
       setPageActive(true)
    },[])

    const [pageIsActive,setPageActive] = useState(false)
    
    const headerRef = useRef(null)
    const form = useRef(null)
    const usernameLabel = useRef(null)
    const usernameInput = useRef(null)
    const passwordLabel = useRef(null)
    const passwordInput = useRef(null)
    const button = useRef(null)

    return (
        <div className="min-h-screen h-auto w-full bg-base text-white font-belanosima flex flex-col justify-center items-center">
            <Transition timeout={400} in={pageIsActive} mountOnEnter unmountOnExit nodeRef={form} onEnter={onEnter}>
            <form ref={form} className="bg-secondary w-[90%] py-[4rem] px-4 grid grid-cols-3 gap-2 grid-rows-3 rounded-md">
                <h2 ref={headerRef} className="col-start-0 col-span-3 text-xl mx-auto">Welcome, please <u className="text-primary decoration-transparent">login</u></h2>
                <label ref={usernameLabel} className="col-start-0 col-span-1 flex flex-col justify-center items-center">Username</label><input ref={usernameInput} className="col-start-0 input col-span-2"></input>
                <label ref={passwordLabel} className="col-start-0 col-span-1 flex flex-col justify-center items-center">Password</label><input ref={passwordInput} type='password' className="col-start-0 input col-span-2"></input>
                <button ref={button} className="col-start-2 col-span-1s mt-4 bg-white text-black rounded-md py-4">Sign In</button>
            </form>
            </Transition>
        </div>
    )

    function onEnter(){
        const tl = gsap.timeline()
        console.log('started')
        tl
        .from(
            [
                form.current,
                headerRef.current,
                usernameLabel.current,usernameInput.current,
                passwordLabel.current,passwordInput.current,
                button.current,
            ],
            {x:'-40vw',autoAlpha:0,stagger:0.03}).play()
        
    }
}

export default LoginPage