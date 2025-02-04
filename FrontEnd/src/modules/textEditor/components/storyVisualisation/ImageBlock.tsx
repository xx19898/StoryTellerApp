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
	const [currImage, setCurrImage] = useState<string | undefined>(undefined)

	const imageInputRef = useRef<HTMLInputElement>(null)
	let imageInputAlreadyOpenedOnce = useRef<boolean>(false)

	useEffect(() => {
		if (
			!currImage &&
			imageInputRef &&
			!imageInputAlreadyOpenedOnce.current
		) {
			console.log('clicked')
			imageInputRef.current!.click()
			imageInputAlreadyOpenedOnce.current = true
		}
	}, [])

	// visualize the data which is feeded through the input
	// implement image deletion, image replacement
	console.log({ currImage, identifier })
	return (
		<div>
			{currImage ? (
				<img src={currImage}></img>
			) : (
				<input
					type='file'
					onInput={(e) => {
						const file = e.target.files[0]
						const fr = new FileReader()
						fr.readAsDataURL(file)
						fr.onload = function () {
							if (fr.result) setCurrImage(fr.result as string)
						}
					}}
					ref={imageInputRef}
				></input>
			)}
		</div>
	)
}

export default ImageBlock
