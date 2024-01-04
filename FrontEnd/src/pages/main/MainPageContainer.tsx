import { useEffect } from "react"
import MainPage from "./MainPage"


const MainPageContainer = () => {
    useEffect(() => {
        console.log("Loading initial data")
    },[])
    return(
        <MainPage
        loggedInUser=""
        onClickMyProfileBut={}
        onClickNewStoryBut={}
        onClickStoryBut={}
        stories={}
        />
    )
}

export default MainPageContainer