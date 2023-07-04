import { ReactNode } from "react"
import { IconContext } from "react-icons"

interface ISuccessNotification{
    text: string,
    icon: ReactNode,
    iconProperties: {size:string,color:string},
    background: string,
}
const Notification = ({text,icon,iconProperties,background}:ISuccessNotification) => {
    
    return(
    <div className="w-full h-auto p-4 text-white font-belanosima text-xl flex flex-col justify-center items-center gap-3 rounded-md" style={{backgroundColor:background}}>
        <p>{text}</p>
        <IconContext.Provider value={{...iconProperties}}>
            {icon}
        </IconContext.Provider>
    </div>
    )
}


export default Notification