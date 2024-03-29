import { useEffect, useRef, useState } from "react"
import { gsap } from "gsap/src"
import { Transition } from "react-transition-group"
import { useForm } from "react-hook-form"
import ErrorComponent from "../../common/forms/ErrorComponent"

interface ILoginForm{
    username:string,
    password:string,
}

interface ILoginPage{
    login: (username:string,password:string) => void,
    loginSuccess: boolean,
    error: string | undefined,
    isLoading: boolean,
    onSuccess: () => void
}

const LoginPage = ({login,error,loginSuccess,onSuccess}:ILoginPage) => {
    const {register,formState:{errors,isValid},getValues} = useForm<ILoginForm>({mode:"onChange",reValidateMode:"onChange"})

    useEffect(() => {
       setPageActive(true)
    },[])

    useEffect(() => {
        if(loginSuccess){
            onExit()
            setPageActive(false)
        }
    },[loginSuccess])

    const [pageIsActive,setPageActive] = useState(false)

    const headerRef = useRef(null)
    const form = useRef(null)
    const usernameLabel = useRef(null)
    const usernameInput = useRef<HTMLElement | null>(null)
    const passwordLabel = useRef(null)
    const passwordInput = useRef<HTMLElement | null>(null)
    const button = useRef(null)

    const {ref:usernameInputRefForValid,...usernameInputRest} = register('username',{minLength:{value:4,message:'Username too short'}})

    const {ref:passwordInputRefForValid,...passwordInputRest} = register('password',{minLength:{value:8,message:'Password too short'}})

    return (
        <div className="min-h-screen h-auto w-full bg-base text-white font-belanosima flex flex-col justify-center items-center">
            <Transition timeout={400} in={pageIsActive} mountOnEnter unmountOnExit nodeRef={form} onEnter={onEnter}>
            <form ref={form} className="bg-secondary sm:flex sm:flex-col sm:justify-center sm:items-center w-[90%] py-[4rem] px-4 lg:grid lg:grid-cols-3 gap-3 lg:grid-rows-3 rounded-md">
                <h2 ref={headerRef} className="col-start-0 col-span-3 text-xl mx-auto">Welcome, please <u className="text-primary decoration-transparent">login</u></h2>
                <label ref={usernameLabel} className="col-start-0 md:col-span-1 flex flex-col justify-center items-center">Username</label>
                <input  data-testid="username-input"  {...usernameInputRest} ref={(e) => {
                    usernameInputRefForValid(e)
                    usernameInput.current = e
                    }} className="col-start-0 text-center input md:col-span-2 sm:w-full text-black"></input>
                {
                errors.username ? <div className="col-span-2 col-start-2"><ErrorComponent errorMessage='Username is too short' /></div> : null
                }
                <label ref={passwordLabel} className="col-start-0 col-span-1 flex flex-col justify-center items-center">Password</label>
                <input data-testid="password-input" {...passwordInputRest} ref={(e) => {
                    passwordInputRefForValid(e)
                    passwordInput.current = e
                    }} type='password' className="col-start-0 text-center input sm:w-full md:col-span-2 text-black"></input>
                {
                errors.username && errors.username.type === 'password' ? <div className="col-span-2 col-start-2"><ErrorComponent errorMessage='Password is too short' /></div> : null
                }
                <button ref={button} onClick={(e) => {
                    e.preventDefault()
                    login(getValues('username'),getValues('password'))
                }} disabled={!isValid} className="md:col-start-2 md:col-span-2 sm:w-1/2 bg-white text-black rounded-md py-4">Sign In</button>
                {
                    error ? <div className="md:col-start-1 col-span-3 sm:w-full flex justify-center items-center">
                        <ErrorComponent errorMessage={error}/>
                        </div> : null
                }
            </form>
            </Transition>
        </div>
    )

    function onEnter(){
        const tl = gsap.timeline()
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

    function onExit(){
        const tl = gsap.timeline()
        tl.to(
            [
                form.current,
                headerRef.current,
                usernameLabel.current,usernameInput.current,
                passwordLabel.current,passwordInput.current,
                button.current,
            ],
            {x:'-40vw',autoAlpha:0,stagger:0.03}).play()
        tl.eventCallback("onComplete",() => onSuccess())
        }
}

export default LoginPage