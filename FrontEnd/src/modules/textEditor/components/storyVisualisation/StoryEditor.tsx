import { useEffect, useRef, useState } from 'react'
import { extractTypeAndContentOfHtmlElement } from '../../helpers/HtmlParsingElementUtilities'
import useAddNewBlock from '../../hooks/useAddNewBlock'
import useGetState from '../../hooks/useGetElementState'
import AddNewBlock from '../addingNewElement/AddNewBlock'
import EditingBlock from './EditingBlock'
import ImageBlock from './ImageBlock'
import { testMap } from '../../textEditorState'
import { useAtom } from 'jotai'
import GeneralInformation from './GeneralInformation'
import TitleEditField from '../editingElements/TitleEditField'

// implement drag and drop to change placement of blocks
const StoryEditor = () => {
	const { elementMap, elementOrderArray, setElementMap } = useGetState()
	const [map, setTestMap] = useAtom(testMap)
	const { addNewImageBlock, addNewTextBlock } = useAddNewBlock()
	const [authorIcon, setAuthorIcon] = useState(null)

	useEffect(() => {
		const newMapx = new Map(map)
		newMapx.set('xdd', 'xddd')
		setTestMap(newMapx)
		const newMap = new Map(elementMap)
		newMap.set('title', 'New story')
	}, [])

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
			<ul className='flex flex-col justify-center items-center gap-4 w-[80%]'>
				{elementOrderArray.map((identifier) => {
					const element = elementMap.get(identifier)

					if (!element) return null

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
