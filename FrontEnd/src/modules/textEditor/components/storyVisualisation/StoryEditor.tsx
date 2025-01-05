import { useRef } from 'react'
import { extractTypeAndContentOfHtmlElement } from '../../helpers/HtmlParsingElementUtilities'
import useAddNewBlock from '../../hooks/useAddNewBlock'
import useGetState from '../../hooks/useGetElementState'
import AddNewBlock from '../addingNewElement/AddNewBlock'
import EditingBlock from './EditingBlock'
import ImageBlock from './ImageBlock'

// implement drag and drop to change placement of blocks
const StoryEditor = () => {
	const { elementMap, elementOrderArray } = useGetState()
	console.log({ elementOrderArray })
	const { addNewImageBlock, addNewTextBlock } = useAddNewBlock()

	const imageInputRef = useRef<HTMLInputElement>(null)

	return (
		<ul className='flex flex-col justify-center items-center gap-4 w-[80%]'>
			{elementOrderArray.map((identifier) => {
				const element = elementMap.get(identifier)
				if (element === undefined) return null

				const { contents, elementType } =
					extractTypeAndContentOfHtmlElement(element)

				if (elementType === 'paragraph')
					return (
						<EditingBlock
							content={contents}
							identifier={identifier}
							type='paragraph'
						/>
					)

				if (elementType === 'title')
					return (
						<EditingBlock
							content={contents}
							identifier={identifier}
							type='title'
						/>
					)

				if (elementType === 'image')
					return (
						<ImageBlock
							content={contents}
							deleteBlock={() => console.log('delete')}
							editBlock={() => console.log('edit block')}
							identifier='xdd'
						/>
					)

				return null
			})}

			<AddNewBlock
				addNewBlock={addNewTextBlock}
				addNewImage={addNewImageBlock}
			/>
		</ul>
	)
}

export default StoryEditor
