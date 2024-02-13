import { SlPencil } from "react-icons/sl"
import { AiOutlineClose } from "react-icons/ai";


interface IToolbar{
    chooseEdit:() => void,
    chooseDelete: () => Promise<void>,
}

const EditingBlockManipulationToolbar = ({chooseEdit,chooseDelete}:IToolbar) => {

    return(
        <div className="absolute right-0 top- w-max flex flex-row gap-2 justify-center items-center py-3 px-5 rounded-md border-white bg-darkestSecondary border-solid border-2">
            <SlPencil size='1.5rem' onClick={() => chooseEdit()}/>
            <AiOutlineClose size='1.5rem' onClick={() => chooseDelete()}/>
        </div>
    )
}

export default EditingBlockManipulationToolbar