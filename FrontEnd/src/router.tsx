import { createBrowserRouter } from "react-router-dom";
import IntroPage from "./pages/introPage";
import LoginPage from "./pages/login/loginPage";
import SignUpPage from "./pages/signUp/signUpPage";
import SignUpPageContainer from "./pages/signUp/signUpPageContainer";


const router = createBrowserRouter([
    {
        path: "/welcome",
        element: <IntroPage/>
    },{
        path: "/login",
        element:<LoginPage />
    },{
        path: "/signUp",
        element:<SignUpPageContainer/>
    }
])

export default router