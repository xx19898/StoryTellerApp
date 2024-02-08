

interface IEditingBlock{
    type: 'paragraph' | 'title',
    content: string,
}

const EditingBlock = ({type,content}:IEditingBlock) => {
    return(
        <div>
            {
                type === 'paragraph' ? <p className="indent-4">{content}</p> : <h2 className="font-semibold text-lg py-2">{content}</h2>
            }
        </div>
    )
}

export default EditingBlock