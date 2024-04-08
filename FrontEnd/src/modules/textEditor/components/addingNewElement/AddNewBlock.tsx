import { useState } from "react"
import { CgAdd } from "react-icons/cg"
import TextOrImage from "./NewImageInput"
 "./NewImageInput"

interface IAddNewBlock{
    addNewBlock:() => void,
    addNewImage: () => void,
}
//TODO: implement and test adding new textblock and adding new image block
const AddNewBlock = ({addNewBlock,addNewImage}:IAddNewBlock) => {
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

export default AddNewBlock