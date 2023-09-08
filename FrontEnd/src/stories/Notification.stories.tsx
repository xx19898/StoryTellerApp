import { BiErrorCircle } from "react-icons/bi"
import Notification from "../common/forms/auth/Notification"
import {BsFillCheckCircleFill} from 'react-icons/bs'

export default{
    component: Notification,
    title: 'Notification',
    decorators: [
        (Story: JSX.Element) => (
          <div className="w-[200px] h-[200px] flex flex-col justify-center items-center">
            {Story}
          </div>
        ),],
    tags:['autodocs'],
}

export const Success = {
    args:{
        text: 'Sign Up successful, you can continue to the Login Page from button below.',
        icon: <BsFillCheckCircleFill/>,
        iconProperties:{
            size:'2em',color:'#3B429F'
        },
        background: '#2A1E5C'
    }
}

export const Failure = {
    args:{
        text: 'Sign Up unsuccessful, something wrong with the server.',
        icon: <BiErrorCircle/>,
        iconProperties:{
            size:'4em',color:'white'
        },
        background: "#3F3F37"
    }
}