import { useEffect, useRef, useState } from "react";

import Story from "../../story/Story";
import EditingBlock from "./storyVisualisation/Block";
import EditingInput from "./editingElements/EditingInput";
import { CgAdd } from "react-icons/cg";
import { addHtmlElementIdentifier, buildHtmlString, extractTypeAndContentOfHtmlElement, getNewIdentifierForElement, processHtmlString, typeToTag } from "../helpers/HtmlParsingUtilities";
import AddNewBlock from "./addingNewElement/AddNewBlock";
import ImageInput from "./editingElements/ImageInput";



//TODO: start implementing images
     export const TextEditor = () => {
        
        
        return(
            <div className="w-auto max-w-[40%] py-2 px-4 min-h-screen h-auto
            flex flex-col gap-[3.5em] justify-start items-center
            text-white">
                <h1>Text editor</h1>
                <section ref={editSectionRef} className="flex gap-4 flex-col justify-start items-center text-md">
                {
                    elementOrderArray && elementOrderArray.map((identifier) => {
                        const el = elementMap?.get(identifier)
                        if(el != undefined){
                            const {contents,element,elementType} = extractTypeAndContentOfHtmlElement(el)
                            if(currentlyEditedElement === identifier){
                                if(elementType === 'paragraph') return <EditingInput identifier={identifier} edit={editBlock} stopEditing={stopEditing} origValue={contents} />
                                else if(elementType === 'image') return <ImageInput setNewImage={}/>
                            }
                            return <EditingBlock 
                                    content={contents} 
                                    type={elementType} 
                                    identifier={identifier}
                                    chooseToEdit={setCurrentlyEditedElement}
                                    chooseToDelete={deleteBlock}
                                    />
                        }                        
                })}
                <AddNewBlock addNewBlock={addNewBlock} />
                </section>
            </div>
            );
    }

