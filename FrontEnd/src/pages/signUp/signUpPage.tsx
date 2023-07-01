import { useRef } from "react"


const SignUpPage = () => {

    const headerRef = useRef(null)
    const usernameLabel = useRef(null)
    const usernameInput = useRef(null)
    const passwordLabel = useRef(null)
    const passwordInput = useRef(null)
    const button = useRef(null)
    const mainRef = useRef(null)
    const formComp = useRef(null)

    return(
        <div ref={mainRef} className="w-auto min-h-screen h-auto bg-base font-belanosima flex flex-col justify-center items-center text-white">
            <form ref={formComp} className="grid grid-cols-3 grid-rows-3 gap-2 justify-center items-center py-[4rem] w-[90%] rounded-md pl-4 pr-8 bg-secondary">
            <h2 ref={headerRef} className="text-xl col-span-3 mx-auto my-4">Enter your information</h2>
            <label ref={usernameLabel} className="mx-auto col-span-1 col-start-0">Username</label><input ref={usernameInput} className=" input col-span-2 col-start-0 col-end-2-3"/>
            <label ref={passwordLabel} className="mx-auto col-span-1 col-start-0">Password</label><input ref={passwordInput} className="input col-span-2"/>
            <button ref={button} className="text-black col-span-3 mx-auto mt-4 btn-secondary bg-white py-4 px-8 rounded-md">Sign Up</button>
            </form>
        </div>
    )
}

export default SignUpPage