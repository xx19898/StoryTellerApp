import axios from "axios"
import LoginPage from "./loginPage"
import {useMutation} from 'react-query'
import { BACKEND_URL } from "../../constants"
import UseSignIn from "../../hooks/mutations/useSignIn"
import { useSetAtom } from "jotai"
import { authTokenAtom } from "../../atomStore"
import { z } from "zod"
import { useNavigate } from "react-router-dom"


const LoginPageContainer = () => {
    const {signIn,error,data,isLoading} = UseSignIn()
    const setAuthToken = useSetAtom(authTokenAtom)
    const navigate = useNavigate()
    const parsedError = error ? parseError(error) : undefined

    console.log({data})

    return(
        <LoginPage
        login={(username:string, password:string) => signIn({
            username:username,
            password:password
        })}
        error={parsedError}
        isLoading={true}
        loginSuccess={data != undefined}
        onSuccess={() => onSuccess(data)}
        />
    )

    function parseError(error:unknown):string{
        const loginErrSchema = z.object({
            response:z.object({
                data:z.object({
                    error: z.string()
                })
            })
        })

        const parsedErrorObject = loginErrSchema.parse(error)
        return parsedErrorObject.response.data.error
    }

    function onSuccess(data:unknown){
        const successfulResponseSchema = z.object({
            data:z.object({
                token: z.string()
            })
        })
        const parsedData = successfulResponseSchema.parse(data)
        console.log({parsedData})
        setAuthToken(parsedData.data.token)
        navigate("/storytellerLobby")
    }

}

export default LoginPageContainer