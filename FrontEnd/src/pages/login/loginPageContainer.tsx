import axios from "axios"
import LoginPage from "./loginPage"
import {useMutation} from 'react-query'
import { BACKEND_URL } from "../../constants"
import UseSignIn from "../../hooks/mutations/useSignIn"


const LoginPageContainer = () => {
    const {data,error,isLoading,signIn} = UseSignIn()
    return(
        <LoginPage login={(username:string, password:string) => signIn({username,password})} />
    )
}

export default LoginPageContainer