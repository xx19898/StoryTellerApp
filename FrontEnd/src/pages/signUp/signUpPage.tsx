import { useEffect, useRef, useState } from "react"
import { Transition } from "react-transition-group"
import { gsap } from "../../gsap"
import { useForm, } from "react-hook-form"

interface ISignUpForm{
    username: string,
    password: string
}

const SignUpPage = () => {

    useEffect(() => {
        setPageIsActive(true)
    },[])

    const [pageIsActive,setPageIsActive] = useState(false)
    const [temp,setTemp] = useState(false)

    const {register,handleSubmit,formState:{errors,isValid}} = useForm<ISignUpForm>({reValidateMode:'onChange',mode:'onChange'})

    const headerRef = useRef(null)
    const usernameLabel = useRef(null)
    const usernameInput = useRef<HTMLElement | null>(null)
    const passwordLabel = useRef(null)
    const passwordInput = useRef<HTMLElement | null>(null)
    const button = useRef(null)
    const mainRef = useRef(null)
    const formComp = useRef(null)

    const {
        ref:refForUsernameInput,
        ...usernameInputRest
    } = register("username",{minLength:{message:'Length of the username should be at least 4 characters',value:4}})

    const {
        ref:refForPasswordInput,
        ...passwordInputRest
    } = register("password",{minLength:8})

    console.log({isValid:isValid})
    console.log({errors:errors})
    
    return(
        <div ref={mainRef} className="w-auto min-h-screen h-auto bg-base font-belanosima flex flex-col justify-center items-center text-white">
            <Transition in={pageIsActive} nodeRef={formComp} onEnter={onEnter} timeout={400}>
            <form ref={formComp} onSubmit={handleSubmit(onSubmit)} className="grid grid-cols-3 grid-rows-3 gap-2 justify-center items-center py-[4rem] w-[90%] rounded-md pl-4 pr-8 bg-secondary">
            <h2 ref={headerRef} className="text-xl col-span-3 mx-auto my-4">Enter your information</h2>
            <label ref={usernameLabel} className="mx-auto col-span-1 col-start-0">Username</label>
            
            <input {...usernameInputRest} defaultValue={''} ref={(e) => {
                refForUsernameInput(e)
                usernameInput.current = e
            }} className="input col-span-2 col-start-0 bg-white text-black"/>
            <label ref={passwordLabel} className="mx-auto col-span-1 col-start-0">Password</label>
            
            <input {...passwordInputRest} ref={(e) => {
                refForPasswordInput(e)
                passwordInput.current = e
            }}
                type={'password'} className="input col-span-2 bg-white text-black"/>
            
            <button ref={button} className="text-black hover:bg-primary col-span-3 mx-auto mt-4 btn-secondary bg-white py-4 px-8 rounded-md">Sign Up</button>
            
            </form>
            </Transition>
        </div>
    )

    function onEnter(){
        const timeline = gsap.timeline()
        timeline.
        from([
            formComp.current,usernameLabel.current,passwordLabel.current,
            headerRef.current,button.current,usernameInput.current,
            passwordInput.current
        ],{
            autoAlpha:0,x:'-100vw',stagger:0.03
        }).play() 
    }

    function onSubmit(){
        console.log('SUBMITTIN THE FORM')
    }
}

export default SignUpPage