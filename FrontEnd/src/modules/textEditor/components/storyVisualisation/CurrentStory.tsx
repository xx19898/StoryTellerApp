import { extractTypeAndContentOfHtmlElement } from "../../helpers/HtmlParsingElementUtilities"
import useGetState from "../../hooks/useGetElementState"


const CurrentStory = () => {
    const {elementMap,elementOrderArray} = useGetState()

    return(
        <ul>
            {

                //TODO: wrap the p and h2 elements in <TextBlock/>
                elementOrderArray.map(identifier => {
                    const element = elementMap.get(identifier) 
                    if(element === undefined) return null

                    const {contents,elementType} = extractTypeAndContentOfHtmlElement(element)

                    if(elementType === "paragraph") return <p>{contents}</p>

                    if(elementType === 'title') return <h2>{contents}</h2>
                    
                    return null
                })
            }
        </ul>
    )
}

export default CurrentStory