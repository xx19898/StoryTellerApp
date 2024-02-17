import { atom, useAtom } from "jotai"

export const currentlyEditedElementAtom = atom<undefined | string>('')

const UseSelectElement = () => {
    const [currentlyEditedElement,setCurrentlyEditedElement] = useAtom(currentlyEditedElementAtom)

    function selectElement(identifier:string | undefined){
        setCurrentlyEditedElement(identifier)
    }

    return {currentlyEditedElement,selectElement}
}

export default UseSelectElement