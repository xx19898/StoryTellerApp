import { useState } from 'react'
import { extractTypeAndContentOfHtmlElement } from '../../helpers/HtmlParsingElementUtilities'
import useAddNewBlock from '../../hooks/useAddNewBlock'
import useGetState from '../../hooks/useGetElementState'
import AddNewBlock from '../addingNewElement/AddNewBlock'
import EditingBlock from './EditingBlock'
import ImageBlock from './ImageBlock'
import { useAtom } from 'jotai'
import GeneralInformation from './GeneralInformation'
import TitleEditField from '../editingElements/TitleEditField'
import useDragAndDrop from '../editingElements/DragAndDrop/useDragAndDrop'
import DragAndDropContainer from '../editingElements/DragAndDrop/dragAndDropContainer'
import { shadowElementArray } from '../editingElements/DragAndDrop/dragAndDropState'

interface IStoryEditor {
	updateStory: (
		elementMap: Map<string, string>,
		elemendOrderArray: string[]
	) => Promise<void>
	updateMainTitle: (newTitle: string) => Promise<void>
}
//{ updateStory, updateMainTitle }: IStoryEditor
const StoryEditor = () => {
	const { elementMap, elementOrderArray } = useGetState()
	const { addNewImageBlock, addNewTextBlock } = useAddNewBlock()

	const [authorIcon, setAuthorIcon] = useState(null)

	const { draggedElement, onDrag, onStopDrag } = useDragAndDrop()
	const [shadowElements, setShadowElements] = useAtom(shadowElementArray)

	const elementOrderArrayToVisualise = draggedElement
		? shadowElements
		: elementOrderArray

	/*
	console.log({
		elementOrderArrayToVisualise,
		shadowElements,
		draggedElement,
	})
		*/

	return (
		<>
			<GeneralInformation
				author={'Driver'}
				creationDate={new Date()}
				lastChangedDate={new Date()}
				key={'Driver'}
			>
				<TitleEditField
					initialTitle='Default title'
					setNewTitle={async (newTitle: string) => {
						return new Promise((resolve) => {
							setTimeout(() => {
								resolve({ status: 'ok' })
							}, 1000)
						})
					}}
				/>
			</GeneralInformation>
			<ul className='flex flex-col justify-center items-center gap-6 w-[80%]'>
				{elementOrderArrayToVisualise.map((identifier) => {
					const element = elementMap.get(identifier)

					if (!element) return null

					const { contents, elementType } =
						extractTypeAndContentOfHtmlElement(element)

					if (elementType === 'paragraph')
						return (
							<DragAndDropContainer
								identifier={identifier}
								key={identifier}
							>
								<EditingBlock
									content={contents}
									identifier={identifier}
									type='paragraph'
								/>
							</DragAndDropContainer>
						)

					if (elementType === 'title')
						return (
							<DragAndDropContainer
								identifier={identifier}
								key={identifier}
							>
								<EditingBlock
									content={contents}
									identifier={identifier}
									type='title'
								/>
							</DragAndDropContainer>
						)

					if (elementType === 'image')
						return (
							<DragAndDropContainer
								identifier={identifier}
								key={identifier}
							>
								<ImageBlock
									content={contents}
									deleteBlock={() => console.log('delete')}
									editBlock={() => console.log('edit block')}
									identifier={identifier}
								/>
							</DragAndDropContainer>
						)

					return null
				})}
				<li key='addNewBlock'>
					<AddNewBlock
						addNewBlock={addNewTextBlock}
						addNewImage={addNewImageBlock}
					/>
				</li>
			</ul>
		</>
	)
}

export default StoryEditor
