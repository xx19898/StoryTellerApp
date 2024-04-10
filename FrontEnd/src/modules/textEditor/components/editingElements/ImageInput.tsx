import { useState } from "react"

interface IAddNewImage{
    setNewImage: (newImage:string) => Promise<void>,
}

function arrayBufferToBase64(buffer:ArrayBuffer){
    let binary = ''
    const bytes = new Uint8Array(buffer)
    const len = bytes.byteLength
    for(let i = 0; i < len; i++){
        binary += String.fromCharCode( bytes[i] )
    }
    return window.btoa(binary)
}

 const ImageInput = ({setNewImage}:IAddNewImage) => {
    const [image,setImage] = useState<string | undefined>(undefined)

    return(
    <div className="flex flex-col justify-center items-center gap-3 bg-white p-3 rounded-md shadow-md">
    <input onChange={ async (e) => {
        if(e.target.files){
            const buffer = await e.target.files[0].arrayBuffer()
            const data = arrayBufferToBase64(buffer)
            setImage(data)
        }
    }}
    type="file"
    className="bg-white p-3 rounded-md shadow-md"
    />
    <div>
    {
        image === undefined ? <p>No image yet</p> : <img src={'data:image/png;base64, ' + image}></img>
    }
    </div>
    </div>
    )
}

export default ImageInput