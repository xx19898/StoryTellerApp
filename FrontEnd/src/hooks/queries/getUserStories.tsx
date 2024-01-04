import { BACKEND_URL } from "../../constants"


const useUserStories = (authToken:string) => {
    const { data } = await axios.get(
        `${BACKEND_URL}`
    )
}