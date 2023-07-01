

const LoginPage = () => {
    
    return (
        <div className="min-h-screen h-auto w-full bg-base text-white font-belanosima flex flex-col justify-center items-center">
            <form className="bg-secondary w-[90%] py-[4rem] px-4 grid grid-cols-3 gap-2 grid-rows-3 rounded-md">
                <h2 className="col-start-0 col-span-3 text-xl mx-auto">Welcome, please <u className="text-primary decoration-transparent">login</u></h2>
                <label className="col-start-0 col-span-1 flex flex-col justify-center items-center">Username</label><input className="col-start-0 input col-span-2"></input>
                <label className="col-start-0 col-span-1 flex flex-col justify-center items-center">Password</label><input type='password' className="col-start-0 input col-span-2"></input>
                <button className="col-start-2 col-span-1 mt-4 bg-white text-black rounded-md py-4">Sign In</button>
            </form>
        </div>
    )
}

export default LoginPage