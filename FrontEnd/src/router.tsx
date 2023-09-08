import { createBrowserRouter } from "react-router-dom";
import IntroPage from "./pages/introPage";
import LoginPage from "./pages/login/loginPage";
import SignUpPageContainer from "./pages/signUp/signUpPageContainer";
import { TextEditor } from "./modules/textEditor/textEditor";


const router = createBrowserRouter([
    {
        path: "/",
        element: <IntroPage/>
    },
    {
        path:"/textEditor",
        element: <TextEditor />
    },
    {
        path: "/login",
        element:<LoginPage login={(username:string,password:string) => console.log({username,password})}/>
    },{
        path: "/signUp",
        element:<SignUpPageContainer/>
    }
])

export default router