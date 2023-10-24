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

    return(
        <SignUpPage
        httpStatus={data?.status}
        signUp={({username,password}:{username:string,password:string}) => signUp({username,password})} />
    )
}

export default SignUpPageContainer