import { useState } from "react"
import EditingBlockManipulationToolbar from "../editingElements/ManipulationToolbar"
import UseSelectElement from "../../hooks/useSelectElement"
import useDeleteElement from "../../hooks/useDeleteElement"

interface IEditingBlock{
    type: 'paragraph' | 'title' | 'image',
    content: string,
    identifier: string,
}

const EditingBlock = ({type,content,identifier}:IEditingBlock) => {
    const [toolbarVisible,setToolbarVisible] = useState(false)
    const { selectElement } = UseSelectElement()
    const {deleteBlock} = useDeleteElement()

    return(
        <div className="relative" onMouseEnter={() => setToolbarVisible(true)} onMouseLeave={() => setToolbarVisible(false)}>
            {
                toolbarVisible && <EditingBlockManipulationToolbar chooseDelete={async () =>  await deleteBlock(identifier)} chooseEdit={() => selectElement(identifier)} />
            }
            {
                type === 'paragraph' ? <p className="indent-4">{content}</p> : <h2 className="font-semibold text-lg py-2">{content}</h2>
            }
        </div>
    )
}

export default EditingBlock