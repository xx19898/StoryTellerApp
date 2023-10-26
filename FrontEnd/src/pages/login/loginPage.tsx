import { useEffect, useRef, useState } from "react"
import { gsap } from "gsap/src"
import { Transition } from "react-transition-group"
import { useForm } from "react-hook-form"
import ErrorComponent from "../../common/forms/ErrorComponent"
import { AxiosResponse, HttpStatusCode } from "axios"
import { useNavigate } from "react-router-dom"

interface ILoginForm{
    username:string,
    password:string,
}

type ISignInRequest = {
    httpCode:  number,
    authToken: string | undefined,
}

interface ILoginPage{
    login: (username:string,password:string) => Promise<{httpStatus:number}>,
    onSuccessfulLogin: () => void,
}

const LoginPage = ({login,onSuccessfulLogin}:ILoginPage) => {
    const {register,formState:{errors,isValid},getValues} = useForm<ILoginForm>({mode:"onChange",reValidateMode:"onChange"})

    useEffect(() => {
       setPageActive(true)
    },[])

    const [pageIsActive,setPageActive] = useState(false)
    const [error,setError] = useState<string | undefined>(undefined)

    const headerRef = useRef(null)
    const form = useRef(null)
    const usernameLabel = useRef(null)
    const usernameInput = useRef<HTMLElement | null>(null)
    const passwordLabel = useRef(null)
    const passwordInput = useRef<HTMLElement | null>(null)
    const button = useRef(null)
    const navigate = useNavigate()

    const {ref:usernameInputRefForValid,...usernameInputRest} = register('username',{minLength:{value:4,message:'Username too short'}})

    const {ref:passwordInputRefForValid,...passwordInputRest} = register('password',{minLength:{value:8,message:'Password too short'}})

    return (
        <div className="min-h-screen h-auto w-full bg-base text-white font-belanosima flex flex-col justify-center items-center">
            <Transition timeout={400} in={pageIsActive} mountOnEnter unmountOnExit nodeRef={form} onEnter={onEnter}>
            <form ref={form} className="bg-secondary sm:flex sm:flex-col sm:justify-center sm:items-center w-[90%] py-[4rem] px-4 lg:grid lg:grid-cols-3 gap-3 lg:grid-rows-3 rounded-md">
                <h2 ref={headerRef} className="col-start-0 col-span-3 text-xl mx-auto">Welcome, please <u className="text-primary decoration-transparent">login</u></h2>
                <label ref={usernameLabel} className="col-start-0 md:col-span-1 flex flex-col justify-center items-center">Username</label>
                <input {...usernameInputRest} ref={(e) => {
                    usernameInputRefForValid(e)
                    usernameInput.current = e
                    }} className="col-start-0 text-center input md:col-span-2 sm:w-full text-black"></input>
                {
                errors.username ? <div className="col-span-2 col-start-2"><ErrorComponent errorMessage='Username is too short' /></div> : null
                }
                <label ref={passwordLabel} className="col-start-0 col-span-1 flex flex-col justify-center items-center">Password</label>
                <input {...passwordInputRest} ref={(e) => {
                    passwordInputRefForValid(e)
                    passwordInput.current = e
                    }} type='password' className="col-start-0 text-center input sm:w-full md:col-span-2 text-black"></input>
                {
                errors.username && errors.username.type === 'password' ? <div className="col-span-2 col-start-2"><ErrorComponent errorMessage='Password is too short' /></div> : null
                }
                <button ref={button} onClick={(e) => {
                    e.preventDefault()
                    onSignInClick()
                }} disabled={!isValid} className="md:col-start-2 md:col-span-2 sm:w-1/2 bg-white text-black rounded-md py-4">Sign In</button>
                {
                    error ? <div className="md:col-start-1 col-span-3 sm:w-full flex justify-center items-center">
                        <ErrorComponent errorMessage={"There was an error when trying to log in, please check your credentials"}/>
                        </div> : null
                }
            </form>
            </Transition>
        </div>
    )

    async function onSignInClick(){
        console.log('got to onsigninclick')
        const result = await login(getValues('username'),getValues('password'))
        if(result.httpStatus == 202){
            onSuccessfulLogin()
        }else{
            setError('Error occured while trying to log in')
        }
    }

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

    function onExit(authToken:string){
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
        tl.eventCallback("onComplete",() => navigate("/storytellerLobby"))
        }
}

export default LoginPage