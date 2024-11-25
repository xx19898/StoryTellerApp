import { extractTypeAndContentOfHtmlElement } from '../../helpers/HtmlParsingElementUtilities'
import useAddNewBlock from '../../hooks/useAddNewBlock'
import useGetState from '../../hooks/useGetElementState'
import AddNewBlock from '../addingNewElement/AddNewBlock'
import EditingBlock from './EditingBlock'

const StoryEditor = () => {
	const { elementMap, elementOrderArray } = useGetState()

	const { addNewImageBlock, addNewTextBlock } = useAddNewBlock()
	console.log({ elementMap, elementOrderArray })
	return (
		<ul className='flex flex-col justify-center items-center gap-4 w-[80%]'>
			{elementOrderArray.map((identifier) => {
				const element = elementMap.get(identifier)
				console.log({ element })
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
