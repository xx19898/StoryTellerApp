import { extractTypeAndContentOfHtmlElement } from "../../helpers/HtmlParsingElementUtilities"
import useGetState from "../../hooks/useGetElementState"
import EditingBlock from "./TextBlock"


const CurrentStory = () => {
    const {elementMap,elementOrderArray} = useGetState()

    console.log({elementOrderArray})

    return(
        <ul>
            {
                //TODO: wrap the p and h2 elements in <TextBlock/>
                elementOrderArray.map(identifier => {
                    const element = elementMap.get(identifier)
                    if(element === undefined) return null

                    const {contents,elementType} = extractTypeAndContentOfHtmlElement(element)

                    if(elementType === "paragraph") return <EditingBlock content={contents} identifier={identifier} type="paragraph"/>

                    if(elementType === 'title') return <EditingBlock content={contents} identifier={identifier} type="title"/>

                    return null
                })
            }
        </ul>
    )
}

export default CurrentStory