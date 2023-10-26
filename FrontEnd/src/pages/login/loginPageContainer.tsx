import axios from "axios"
import LoginPage from "./loginPage"
import {useMutation} from 'react-query'
import { BACKEND_URL } from "../../constants"
import UseSignIn from "../../hooks/mutations/useSignIn"
import { useSetAtom } from "jotai"
import { authTokenAtom } from "../../atomStore"


const LoginPageContainer = () => {
    const {signIn,error} = UseSignIn()
    const setAuthToken = useSetAtom(authTokenAtom)

    return(
        <LoginPage
        login={(username:string, password:string) => signIn({
            username:username,
            password:password
        })}
        error={error}
        setAuthToken={(newAtom) => setAuthToken(newAtom)}
        />
    )

    async function sleep(){
        return new Promise(resolve => setTimeout(resolve,1000))
    }

    async function mockSignIn(){
        await sleep()
        return {httpCode:400,authToken:undefined}
    }

}

export default LoginPageContainer