

interface IErrorComponent{
    errorMessage: string
}

const ErrorComponent = ({errorMessage}:IErrorComponent) => {

    return(
        <div className='bg-transparent rounded-md w-full h-auto flex flex-row justify-center items-center font-belanosima border-solid border-[3px] border-warning p-2 gap-3'>
            <p>{errorMessage}</p>
        </div>
    )
}

export default ErrorComponent