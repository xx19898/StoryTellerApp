import axios from "axios"
import LoginPage from "./loginPage"
import {useMutation} from 'react-query'
import { BACKEND_URL } from "../../constants"


const LoginPageContainer = () => {
    const {mutate,data,error} = useMutation({
        mutationFn: ({username,password}:{username:string,password:string}) => {
            return axios.post(`${BACKEND_URL}/api/signIn`,{username:username,password:password})
        },
    })
    return(
        <LoginPage login={login} error={error} data={data}/>
    )

    function login(username:string,password:string){
        mutate({username:username,password:password})
    }
}

export default LoginPageContainer