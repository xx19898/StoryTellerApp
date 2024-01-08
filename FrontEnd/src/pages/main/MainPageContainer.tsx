import { useEffect } from "react"
import MainPage from "./MainPage"


const MainPageContainer = () => {
    useEffect(() => {
        console.log("Loading initial data")
    },[])
    return(
        <MainPage
        loggedInUser=""
        onClickMyProfileBut={() => console.log('clicked')}
        onClickNewStoryBut={() => console.log('clicked')}
        onClickStoryBut={() => console.log('clicked')}
        stories={[{id:2,name:'xd'}]}
        />
    )
}

export default MainPageContainer