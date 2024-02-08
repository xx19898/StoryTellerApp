

interface IEditingInput{
    value: string,
    setValue: (val:string) => void,
}

const EditingInput = () => {
    return(
        <div className="bg-special w-full flex justify-center items-center">
            <input className="indent-0 p-2 w-full text-black" value="lorum ipsum" >
            
            </input>
        </div>
    )
}

export default EditingInput