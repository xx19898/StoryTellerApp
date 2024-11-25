import { createBrowserRouter } from 'react-router-dom'
import IntroPage from './pages/introPage'
import LoginPage from './pages/login/loginPage'
import SignUpPageContainer from './pages/signUp/signUpPageContainer'
import MainPageContainer from './pages/main/MainPageContainer'
import LoginPageContainer from './pages/login/loginPageContainer'
import { TextEditor } from './modules/textEditor'

const router = createBrowserRouter([
	{
		path: '/',
		element: <IntroPage />,
	},
	{
		path: '/tellstory',
		element: <TextEditor />,
	},
	{
		path: '/storytellerLobby',
		element: <MainPageContainer />,
	},
	{
		path: '/login',
		element: <LoginPageContainer />,
	},
	{
		path: '/signUp',
		element: <SignUpPageContainer />,
	},
])

export default router
