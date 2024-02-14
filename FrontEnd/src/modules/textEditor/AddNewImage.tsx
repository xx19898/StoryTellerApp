import { useState } from "react"


interface IAddNewImage{
    addNewImage: (newImage:string) => void,    
}

export default () => {
    const [image,setImage] = useState<string | undefined>(undefined)

    return(
    <div>
    <input onChange={async (e) => {
        if(e.target.files){
            const buffer = await e.target.files[0].arrayBuffer()
            
            console.log({buffer})
        }
    }} type="file"/>
    {
        image === 'undefined' ? <p>No image yet</p> : <img></img> 
    }
    </div>
    )
}