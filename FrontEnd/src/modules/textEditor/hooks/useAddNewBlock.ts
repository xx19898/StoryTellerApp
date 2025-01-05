import {
	addNewImage,
	addNewParagraph,
	addNewTitle,
} from '../helpers/HtmlParsingUtilities'
import useGetState from './useGetElementState'
import UseSelectElement from './useSelectElement'

const useAddNewBlock = () => {
	const {
		elementMap,
		elementOrderArray,
		setElementMap,
		setElementOrderArray,
	} = useGetState()

	const { currentlyEditedElement, selectElement } = UseSelectElement()

	function addNewTextBlock() {
		const newElMap = new Map(elementMap)
		const newElArray = [...elementOrderArray]
		console.log({ arrPreAdding: newElArray })
		addNewParagraph('', newElMap, newElArray)
		console.log({ arrPostAdding: newElArray })
		setElementMap(newElMap)
		setElementOrderArray(newElArray)

		const lastIdentifier = newElArray.slice(-1)[0]
		selectElement(lastIdentifier)
	}

	function addMainTitle(title: string) {
		const newElMap = new Map(elementMap)
		const newElArray = [...elementOrderArray]

		addNewTitle(title, newElMap, newElArray)
		setElementMap(newElMap)
		setElementOrderArray(newElArray)
	}

	function addNewImageBlock() {
		const newElMap = new Map(elementMap)
		const newElArray = [...elementOrderArray]

		addNewImage('', newElMap, newElArray)
		setElementMap(newElMap)
		setElementOrderArray(newElArray)
	}

	return { addNewTextBlock, addNewImageBlock }
}

export default useAddNewBlock
