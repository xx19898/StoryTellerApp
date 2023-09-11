import { useMutation } from "react-query"
import { BACKEND_URL } from "../../constants"
import axios from "axios"

const UseSignIn = () => {
    const {mutate:signIn,isLoading,data,error} = useMutation({
        mutationFn: (params: {username:string, password:string}) => {
            return axios.post(`${BACKEND_URL}/auth/login`,{
                username: params.username,
                password: params.password,
            })
        }
    })

    return {signIn,isLoading,data,error}
}

export default UseSignIn