import { useMutation } from "react-query";
import { BACKEND_URL } from "../../constants";
import SignUpPage from "./signUpPage";
import axios from "axios";
import UseSignUp from "../../hooks/mutations/useSignUp";


const SignUpPageContainer = () => {

    const {signUp,error,data} = UseSignUp()

    function sleep(ms:number){
        return new Promise(resolve => setTimeout(resolve, ms))
    }

    function mockFunction(username: string){
        const mockAnswer = sleep(500).then(() => {
            console.log({ username:username })
            const evalResult = username === 'freeUser'
            console.log({evalResult})
            return evalResult
        })
        return mockAnswer
    }

    function registerFunction(username:string,password:string){
        console.log(`MAKING REQUEST TO ${BACKEND_URL}`)
        if(password === 'correct' && username === 'test') return {status:201,message:'Allt bra'}
        return {status:500,message: 'Could not create password, sorry'}
    }

    return(
        <SignUpPage
        httpStatus={data?.status}
        signUp={({username,password}:{username:string,password:string}) => signUp({username,password})} />
    )
}

export default SignUpPageContainer