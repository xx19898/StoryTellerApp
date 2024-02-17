import { useEffect, useRef } from "react"
import UseSelectElement from "./useSelectElement"


export default () => {
    const {selectElement} = UseSelectElement()
    const editSectionRef = useRef<HTMLBaseElement>(null)

    useEffect(() => {
        document.addEventListener("mousedown", onClickOutsideEditSection)
        return () => {
            document.addEventListener("mousedown", onClickOutsideEditSection)
        }
    },[])

    function onClickOutsideEditSection(e:MouseEvent){
                if (editSectionRef.current && !editSectionRef.current.contains(e.target as Node)) {
                    stopEditing()
                }
            }

    function stopEditing(){
        selectElement(undefined)
    }
}
    