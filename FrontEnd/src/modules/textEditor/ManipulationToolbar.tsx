import { SlPencil } from "react-icons/sl"
import { SlPlus } from "react-icons/sl"


const EditingBlockManipulationToolbar = () => {

    return(
        <div className="w-max flex flex-row gap-2 justify-center items-center py-3 px-5 rounded-md border-white border-solid border-2">
            <SlPencil size='1.5rem' />
            <SlPlus size='1.5rem'  />
        </div>
    )
}

export default EditingBlockManipulationToolbar