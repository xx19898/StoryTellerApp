import { extractTypeAndContentOfHtmlElement } from "../../helpers/HtmlParsingElementUtilities"
import useAddNewBlock from "../../hooks/useAddNewBlock"
import useDeleteElement from "../../hooks/useDeleteElement"
import useGetState from "../../hooks/useGetElementState"
import AddNewBlock from "../addingNewElement/AddNewBlock"
import ImageBlock from "./ImageBlock"
import EditingBlock from "./TextBlock"


const StoryEditor = () => {
    const {elementMap,elementOrderArray} = useGetState()
    const {addNewImageBlock,addNewTextBlock} = useAddNewBlock()

    const {deleteBlock} = useDeleteElement()

    console.log({length:elementOrderArray.length})

    //TODO: start figuring out how to add images to the blogz
    return(
        <ul className="w-full flex flex-col justify-center items-center gap-[0.5em]">
            {
                elementOrderArray.map(identifier => {
                    const element = elementMap.get(identifier)
                    if(element === undefined) return null

                    const {contents,elementType} = extractTypeAndContentOfHtmlElement(element)

                    if(elementType === "paragraph") return <EditingBlock key={identifier} content={contents} identifier={identifier} type="paragraph"/>

                    if(elementType === 'title') return <EditingBlock key={identifier} content={contents} identifier={identifier} type="title"/>

                    if(elementType === 'image') return <ImageBlock key={identifier} content="" delete={deleteBlock} />

                    return null
                })
            }
            <AddNewBlock
            addNewBlock={addNewTextBlock}
            addNewImage={addNewImageBlock}
            />
        </ul>
    )
}

export default StoryEditor