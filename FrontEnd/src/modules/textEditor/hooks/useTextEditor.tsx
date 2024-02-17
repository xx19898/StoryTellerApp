import { useEffect, useRef } from "react"
import { addHtmlElementIdentifier, buildHtmlString, extractTypeAndContentOfHtmlElement, processHtmlString, typeToTag } from "../helpers/HtmlParsingUtilities"
import { atom, useAtom, useSetAtom } from "jotai"


//TODO: create state and state changing interface with jotai
export const elementOrderArrayAtom =  atom<string[]>([])
export const elementMapAtom = atom(new Map())
            
            async function deleteBlock(blockIdentifier:string){
                const newElementMap = new Map(elementMap)
                if(newElementMap.delete(blockIdentifier)){
                    setElementMap(newElementMap)
                    await sendChangedStoryToServer()
                }
            }

            async function addNewBlock(blockType: 'image' | 'text' | 'title'){
                const newElementOrderArray = [...elementOrderArray]
                const newIdentifier = getNewIdentifierForElement()
                newElementOrderArray.push(newIdentifier)
                const newElementMap = new Map(elementMap)
                if(blockType === 'text'){
                    newElementMap.set(newIdentifier,'<p></p>')
                }else if(blockType === 'image'){
                    newElementMap.set(newIdentifier,'<p></p>')
                } 
                setElementOrderArray(newElementOrderArray)
                setElementMap(newElementMap)
                await sendChangedStoryToServer()
                setCurrentlyEditedElement(newIdentifier)
            }  

            function onClickOutsideEditSection(e:MouseEvent){
                if (editSectionRef.current && !editSectionRef.current.contains(e.target as Node)) {
                    stopEditing()
                }
            }

            function stopEditing(){
                setCurrentlyEditedElement(undefined)
            }
    
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
                }
                    )
            }

        