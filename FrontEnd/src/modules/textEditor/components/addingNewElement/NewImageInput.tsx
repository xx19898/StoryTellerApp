
import { useState } from "react";
import { CiTextAlignRight } from "react-icons/ci";
import { FaRegFileImage } from "react-icons/fa";

interface IAddNewBlock{
    addNewBlock:() => void,
    addNewImage: () => void,
}

const TextOrImage = ({addNewBlock,addNewImage}:IAddNewBlock) => {
    
    return(
        <div className="flex flex-row justify-center items-center gap-2">
            <CiTextAlignRight
            onClick={() => addNewBlock()}
            className="cursor-pointer" 
            size={'3rem'}  />
            <FaRegFileImage 
            className="cursor-pointer" 
            size={'3rem'} 
            onClick={() => addNewImage()}
            />
        </div>
    )
}


export default TextOrImage