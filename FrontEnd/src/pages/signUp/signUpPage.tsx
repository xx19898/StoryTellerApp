import { useEffect, useRef, useState } from "react"
import { Transition } from "react-transition-group"
import { gsap } from "../../gsap"
import { useForm, } from "react-hook-form"
import ErrorComponent from "../../common/forms/ErrorComponent"
import { useDebouncedCallback } from "use-debounce"
import {BsFillPersonCheckFill, BsPersonFillAdd} from 'react-icons/bs'
import { IconContext } from "react-icons"

interface ISignUpForm{
    username: string,
    password: string,
}

const SignUpPage = (
    {checkIfUsernameIsTaken,signUp}:{
        checkIfUsernameIsTaken: (username:string) => Promise<boolean>,
        signUp: (username:string,password:string) => {status:number,message:string}
    }) => {

    useEffect(() => {
        setPageIsActive(true)
    },[])

    const [pageIsActive,setPageIsActive] = useState(false)
    const [signUpSuccess,setSignUpSuccess] = useState(false)
    const {register,handleSubmit,formState:{errors,isValid},setError,clearErrors,getValues} = useForm<ISignUpForm>({reValidateMode:'onChange',mode:'onChange'})

    const headerRef = useRef(null)
    const usernameLabel = useRef(null)
    const usernameInput = useRef<HTMLElement | null>(null)
    const passwordLabel = useRef(null)
    const passwordInput = useRef<HTMLElement | null>(null)
    const button = useRef(null)
    const mainRef = useRef(null)
    const formComp = useRef(null)
    const loginRedirectButton = useRef(null)

    const usernameTaken = useDebouncedCallback(
        async (username:string) => {
            if(username.length === 0) {
                clearErrors("username")
                setError('username',{type:'required',message:'This field cannot be empty'})
            }
            const result = await checkIfUsernameIsTaken(username)
            console.log({checkResult:result})
            if(result) clearErrors('username')
            else{
                setError('username',{type:'usernameTaken'})
            }
            return result
        },
        10
    )

    const {
        ref:refForUsernameInput,
        ...usernameInputRest
    } = register("username",{
        minLength:{message:'Length of the username should be at least 4 characters',value:4},
        required:{value:true,message:'This field cannot be empty'},
        validate:{
            usernameTaken,
        }})

    const {
        ref:refForPasswordInput,
        ...passwordInputRest
    } = register("password",{minLength:{value:8,message:'Password should contain at least 8 characters'},required:{value:true,message:'This field cannot be empty'}})

    console.log({isValid:isValid})
    console.log({errors:errors})

    return(
        <div ref={mainRef} className="w-auto min-h-screen h-auto bg-base font-belanosima flex flex-col justify-center items-center text-white">
            <Transition in={pageIsActive} nodeRef={formComp} onEnter={onEnter} timeout={400}>
            <form ref={formComp} onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-2 justify-center items-center py-[4rem] w-[90%] rounded-md pl-4 pr-8 bg-secondary">
            <h2 ref={headerRef} className="text-xl mx-auto my-4">Enter your information</h2>
            <label ref={usernameLabel} className="mx-auto col-start-0">Username</label>

            <input {...usernameInputRest} defaultValue={''} ref={(e) => {
                refForUsernameInput(e)
                usernameInput.current = e
            }} className="bg-white text-black"/>
            {
                errors.username && errors.username.type === 'usernameTaken' ? <div className="col-span-2 col-start-2"><ErrorComponent errorMessage='taken' /></div> : null
            }
            {
                (errors.username && errors.username.type != 'usernameTaken') ? <div className="col-span-2 col-start-2"><ErrorComponent errorMessage={errors.username.message as string} /></div> : null
            }

            <label ref={passwordLabel} className="mx-auto col-span-1 col-start-0">Password</label>

            <input {...passwordInputRest} ref={(e) => {
                refForPasswordInput(e)
                passwordInput.current = e
            }}
            type={'password'} className="input col-span-2 bg-white text-black"/>

            {
                errors.password ? <div className="col-span-2 col-start-2 w-full"><ErrorComponent errorMessage={errors.password.message as string} /></div> : null
            }


            {
                signUpSuccess ?
                <Transition timeout={400} in={signUpSuccess} onEnter={onEnterSignUpSuccess}>
                <div className="flex flex-col justify-center items-center bg-secSpecial rounded-md w-full p-3 col-span-2 col-start-2" ref={loginRedirectButton}>
                        <IconContext.Provider value={{size:'3em',color:'#D81E5B'}} >
                            <BsFillPersonCheckFill />
                        </IconContext.Provider>
                    <p className="indent-3 textarea-xl leading-5">Sign Up Succeeded! You can follow through to the Login page by clicking on the button below</p>
                </div>
                </Transition>
                :
                <button ref={button} className="text-black hover:bg-primary col-span-3 mx-auto mt-4 btn-secondary bg-white py-4 px-8 rounded-md">Sign Up <IconContext.Provider value={{color:'#3B429F',size:'2em'}}><BsPersonFillAdd /></IconContext.Provider> </button>
            }
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

    function onEnterSignUpSuccess(){
        const timeline = gsap.timeline()
        timeline.
        from(
            loginRedirectButton.current
        ,{
            autoAlpha:0
        })
    }

    function onSubmit(){
        const response = signUp(getValues('username'),getValues('password'))
        if(response.status === 201) setSignUpSuccess(true)
    }
}

export default SignUpPage