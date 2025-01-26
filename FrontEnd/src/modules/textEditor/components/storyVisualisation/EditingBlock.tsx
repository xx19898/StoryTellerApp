import { useState } from 'react'
import EditingBlockManipulationToolbar from '../editingElements/ManipulationToolbar'
import UseSelectElement from '../../hooks/useSelectElement'
import useDeleteElement from '../../hooks/useDeleteElement'
import EditingInput from '../editingElements/EditingInput'

interface IEditingBlock {
	type: 'paragraph' | 'title' | 'image'
	content: string
	identifier: string
}

const EditingBlock = ({ type, content, identifier }: IEditingBlock) => {
	const [toolbarVisible, setToolbarVisible] = useState(false)
	const { selectElement, currentlyEditedElement } = UseSelectElement()
	const { deleteBlock } = useDeleteElement()

	if (currentlyEditedElement === identifier)
		return <EditingInput origValue={content} identifier={identifier} />

	return (
		<li
			className='relative p-8 h-auto flex flex-wrap justify-center items-center'
			onMouseEnter={() => setToolbarVisible(true)}
			onMouseLeave={() => setToolbarVisible(false)}
			key={identifier}
		>
			{toolbarVisible && (
				<EditingBlockManipulationToolbar
					chooseDelete={async () => await deleteBlock(identifier)}
					chooseEdit={() => selectElement(identifier)}
				/>
			)}
			{type === 'paragraph' ? (
				<p className='p-4 rounded-md bg-darkerSecondary indent-4 break-all'>
					{content} {type}
				</p>
			) : (
				<h2 className='font-semibold text-xl py-2 text-center'>
					{content} {type}
				</h2>
			)}
		</li>
	)
}

export default EditingBlock
