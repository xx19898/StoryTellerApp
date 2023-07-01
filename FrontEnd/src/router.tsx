import { createBrowserRouter } from "react-router-dom";
import IntroPage from "./pages/introPage";
import LoginPage from "./pages/login/loginPage";
import SignUpPage from "./pages/signUp/signUpPage";


const router = createBrowserRouter([
    {
        path: "/welcome",
        element: <IntroPage/>
    },{
        path: "/login",
        element:<LoginPage />
    },{
        path: "/signUp",
        element:<SignUpPage />
    }
])

export default router