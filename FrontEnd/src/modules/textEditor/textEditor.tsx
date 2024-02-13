import { useEffect, useRef, useState } from "react";

import Story from "../story/Story";
import EditingBlock from "./EditingBlock";
import EditingInput from "./EditingInput";
import { addHtmlElementIdentifier, buildHtmlString, extractTypeAndContentOfHtmlElement, processHtmlString, typeToTag } from "./htmlParsingUtilities";


     export const TextEditor = () => {
        const [elementOrderArray,setElementOrderArray] = useState<string[]>([])
        const [elementMap,setElementMap] = useState<Map<string,string>>(new Map())

        const [currentlyEditedElement,setCurrentlyEditedElement] = useState<string | undefined>(undefined)

        const editSectionRef = useRef<HTMLBaseElement>(null)

        useEffect(() => {
            document.addEventListener("mousedown", onClickOutsideEditSection)
            return () => {
                document.addEventListener("mousedown", onClickOutsideEditSection)
            }
        },[])
        
        useEffect(() => {
            const htmlString = '<h2>Title</h2><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi sollicitudin consequat condimentum. Suspendisse vitae libero et mi semper molestie. Suspendisse sed bibendum arcu. Suspendisse et aliquam tortor, eget sagittis lacus. Maecenas consectetur sollicitudin turpis, sed consequat felis mollis at. Nunc nec lectus condimentum, ultrices eros ut, auctor eros. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Fusce a sapien pharetra, pulvinar nibh ac, vestibulum lorem.</p><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam a velit lacinia, varius lorem id, euismod massa. Integer ornare varius congue. Pellentesque congue nulla quis mauris tincidunt, vel consectetur lorem.</p>'
            
            const {
                htmlElementMap,htmlOrderArray  
            } = processHtmlString(htmlString)
            setElementOrderArray(htmlOrderArray)
            setElementMap(htmlElementMap)
            
        },[])
        

        async function editBlock(newTextContext:string,blockIdentifier:string){
            
            const newElementMap = new Map(elementMap)
            const element = elementMap.get(blockIdentifier)
            if(element != undefined){
                const {elementType} = extractTypeAndContentOfHtmlElement(element)
                const tag = typeToTag(elementType)
                if(tag != 'unknown'){
                    const newElement = addHtmlElementIdentifier(tag,newTextContext) 
                    newElementMap.set(blockIdentifier,newElement)
                    setElementMap(newElementMap)
                }
            }

            await sendChangedStoryToServer()
        }

        async function deleteBlock(blockIdentifier:string){
            const newElementMap = new Map(elementMap)
            if(newElementMap.delete(blockIdentifier)){
                setElementMap(newElementMap)
                await sendChangedStoryToServer()
            }
        }

        async function sendChangedStoryToServer(){
            const newStoryString = buildHtmlString(elementMap,elementOrderArray)
            //SWAP FOR REAL MUTATION LATER
            await new Promise(resolve => {
                setTimeout(() => resolve("xd"),3000)
            }
                )
        }  

        function onClickOutsideEditSection(e:MouseEvent){
            if (editSectionRef.current && !editSectionRef.current.contains(e.target as Node)) {
                stopEditing()
            }
        }

        function stopEditing(){
            setCurrentlyEditedElement(undefined)
        }
        
        return(
            <div className="w-auto max-w-[40%] py-2 px-4 min-h-screen h-auto
            flex flex-col gap-[3.5em] justify-start items-center
            text-white">
                <h1>Text editor</h1>
                <section ref={editSectionRef}  className="flex gap-4 flex-col justify-start items-center text-md">
                {
                    elementOrderArray && elementOrderArray.map((identifier) => {
                        const el = elementMap?.get(identifier)
                        if(el != undefined){
                            const {contents,element,elementType} = extractTypeAndContentOfHtmlElement(el)
                            if(currentlyEditedElement === identifier) return <EditingInput identifier={identifier} edit={editBlock} stopEditing={stopEditing} origValue={contents} />
                            return <EditingBlock 
                                    content={contents} 
                                    type={elementType} 
                                    identifier={identifier}
                                    chooseToEdit={setCurrentlyEditedElement}
                                    chooseToDelete={deleteBlock}
                                    />
                        }                        
                })
                }
                </section>
            </div>);
    }

