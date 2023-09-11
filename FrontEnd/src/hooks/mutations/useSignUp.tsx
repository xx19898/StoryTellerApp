import axios from "axios"
import { useMutation } from "react-query"
import { BACKEND_URL } from "../../constants"


    const UseSignUp = () => {
        const {mutate:signUp,isLoading,data,error} = useMutation({
            mutationFn: (params:{ username:string, password:string}) => {
                return axios.post(`${BACKEND_URL}/auth/register`,
                    {
                        username: params.username,
                        password: params.password
                    })
            }
        })

        return {signUp,isLoading,data,error}
}


export default UseSignUp
