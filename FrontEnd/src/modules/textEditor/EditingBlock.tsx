import { useState } from "react"
import EditingBlockManipulationToolbar from "./ManipulationToolbar"

interface IEditingBlock{
    type: 'paragraph' | 'title',
    content: string,
    identifier: string,
    chooseToEdit: (identifier:string) => void,
    chooseToDelete: (identifier: string) => Promise<void>,
}

const EditingBlock = ({type,content,chooseToEdit,identifier,chooseToDelete}:IEditingBlock) => {
    const [toolbarVisible,setToolbarVisible] = useState(false)

    function startEditing(){
        chooseToEdit(identifier)
    }

    return(
        <div className="relative" onMouseEnter={() => setToolbarVisible(true)} onMouseLeave={() => setToolbarVisible(false)}>
            {
                toolbarVisible && <EditingBlockManipulationToolbar chooseDelete={async () =>  await chooseToDelete(identifier)} chooseEdit={() => startEditing()} />
            }
            {
                type === 'paragraph' ? <p className="indent-4">{content}</p> : <h2 className="font-semibold text-lg py-2">{content}</h2>
            }
        </div>
    )
}

export default EditingBlock