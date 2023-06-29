import { createBrowserRouter } from "react-router-dom";
import IntroPage from "./pages/introPage";


const router = createBrowserRouter([
    {
        path: "/welcome",
        element: <IntroPage/>
    }
])

export default router