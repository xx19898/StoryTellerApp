import { useState } from "react"
import { CgAdd } from "react-icons/cg"
import TextOrImage from "./newImageInput"
 "./NewImageInput"

interface IAddNewBlock{
    addNewBlock:() => void,
    addNewImage: () => void,
}

export default ({addNewBlock,addNewImage}:IAddNewBlock) => {
    const [hovered,setHovered] = useState(false)

    return(
        <div
        className="bg-secondPrimary"
        onMouseEnter={() => setHovered(true)}
        onMouseLeave={() => setHovered(false)}>
            {
                hovered ? <TextOrImage 
                addNewBlock={addNewBlock}
                addNewImage={addNewImage}
                 /> : <CgAdd
                className="cursor-pointer" 
                size={'3rem'} 
                onClick={() => addNewBlock()}
                />
            }
        </div>
    )
}