

//TODO: create a hook, which would take in array of refs and action function

import { useEffect } from "react"

// and check if the components refs been connected to have been clicked on
interface IUseDetectClickOutside{
    ref: React.RefObject<HTMLElement>,
    onClickOutside: () => void,
}

const useDetectClickOutside = ({onClickOutside,ref}:IUseDetectClickOutside) => {
    useEffect(() => {
        document.addEventListener("mousedown", onClickOutside)
        return () => {
            document.removeEventListener("mousedown", onClickOutside)
        }
    },[])
}

export default useDetectClickOutside