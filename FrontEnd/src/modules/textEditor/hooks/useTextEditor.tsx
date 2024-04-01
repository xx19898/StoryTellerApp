import { useEffect, useRef } from "react"
import { addHtmlElementIdentifier, buildHtmlString, processHtmlString } from "../helpers/HtmlParsingUtilities"
import { atom, useAtom, useSetAtom } from "jotai"
import useGetState from "./useGetElementState"


//TODO: create state and state changing interface with jotai
export const elementOrderArrayAtom =  atom<string[]>([])
export const elementMapAtom = atom(new Map())

const UseTextEditor = () => {
    //TO USE ONLY WHILE DEVELOPING WITH STORYBOOK, LATER IS REPLACED BY REACT QUERY
    const {
        setElementMap,
        setElementOrderArray
    } = useGetState()

    useEffect(() => {
        const htmlString = '<h2>Title</h2><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi sollicitudin consequat condimentum. Suspendisse vitae libero et mi semper molestie. Suspendisse sed bibendum arcu. Suspendisse et aliquam tortor, eget sagittis lacus. Maecenas consectetur sollicitudin turpis, sed consequat felis mollis at. Nunc nec lectus condimentum, ultrices eros ut, auctor eros. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Fusce a sapien pharetra, pulvinar nibh ac, vestibulum lorem.</p><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam a velit lacinia, varius lorem id, euismod massa. Integer ornare varius congue. Pellentesque congue nulla quis mauris tincidunt, vel consectetur lorem.</p>'

        const {
            htmlElementMap,htmlOrderArray
        } = processHtmlString(htmlString)

        setElementOrderArray(htmlOrderArray)
        setElementMap(htmlElementMap)
    },[])
}
    /*
    FOR LATER, NOTICE CLICK OUTSIDE OF COMPONENT

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
                setCurrentlyEditedElement(undefined)
            }

    */


export async function sendChangedStoryToServer(){
            //const newStoryString = buildHtmlString(elementMap,elementOrderArray)
            //SWAP FOR REAL MUTATION LATER
            await new Promise(resolve => {
                setTimeout(() => resolve("xd"),500)
            })
        }

export default UseTextEditor