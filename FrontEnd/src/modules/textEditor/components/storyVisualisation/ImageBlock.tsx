import { useEffect, useRef, useState } from 'react'
import UseSelectElement from '../../hooks/useSelectElement'

interface IImageBlock {
	identifier: string
	deleteBlock: () => void
	editBlock: () => void
	content: string
}

const ImageBlock = ({
	content,
	deleteBlock,
	editBlock,
	identifier,
}: IImageBlock) => {
	const { selectElement, currentlyEditedElement } = UseSelectElement()
	const [currImage, setCurrImage] = useState(undefined)

	const imageInputRef = useRef<HTMLInputElement>(null)
	let prop = useRef<boolean>(false)

	useEffect(() => {
		if (!currImage && imageInputRef) {
			console.log('clicked')
			imageInputRef.current!.click()
		}
	}, [])
	//create input element which opens itself on render, otherwise show chosen image

	return (
		<li>
			{currImage ? (
				<p>{currImage}</p>
			) : (
				<input type='file' ref={imageInputRef}></input>
			)}
		</li>
	)
}

export default ImageBlock
