

interface IImageBlock{
    identifier: string,
    delete: () => void,
    edit: () => void,
    content: string,
}
const ImageBlock = ({delete, edit, identifier, content}:IImageBlock) => {
    const { selectElement,currentlyEditedElement } = UseSelectElement()

    return(
    <li>

    </li>
    )
}

export default ImageBlock