import { SlPencil } from "react-icons/sl"
import { SlPlus } from "react-icons/sl"


interface IToolbar{
    chooseEdit:() => void,
}

const EditingBlockManipulationToolbar = ({chooseEdit}:IToolbar) => {

    return(
        <div className="absolute right-0 top- w-max flex flex-row gap-2 justify-center items-center py-3 px-5 rounded-md border-white bg-darkestSecondary border-solid border-2">
            <SlPencil size='1.5rem' onClick={() => chooseEdit()}/>
            <SlPlus size='1.5rem'  />
        </div>
    )
}

export default EditingBlockManipulationToolbar