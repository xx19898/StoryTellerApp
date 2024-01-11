import parse from 'html-react-parser';

interface IStory{
    htmlString: string,
}

export default ({htmlString}:IStory) => {
    return(
        <div className="flex flex-col justify-center items-center">
            {
                parse(htmlString)
            }
        </div>
    )
}