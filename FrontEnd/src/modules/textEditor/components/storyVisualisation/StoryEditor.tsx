import { extractTypeAndContentOfHtmlElement } from "../../helpers/HtmlParsingElementUtilities"
import useGetState from "../../hooks/useGetElementState"
import AddNewBlock from "../addingNewElement/AddNewBlock"
import EditingBlock from "./TextBlock"


const StoryEditor = () => {
    const {elementMap,elementOrderArray} = useGetState()

    const {addNewImageBlock,addNewTextBlock} = useAddNewBlock()

    return(
        <ul>
            {
                elementOrderArray.map(identifier => {
                    const element = elementMap.get(identifier)
                    if(element === undefined) return null

                    const {contents,elementType} = extractTypeAndContentOfHtmlElement(element)

                    if(elementType === "paragraph") return <EditingBlock content={contents} identifier={identifier} type="paragraph"/>

                    if(elementType === 'title') return <EditingBlock content={contents} identifier={identifier} type="title"/>

                    return null
                })
            }
            <AddNewBlock />
            </li>
        </ul>
    )
}

export default StoryEditor