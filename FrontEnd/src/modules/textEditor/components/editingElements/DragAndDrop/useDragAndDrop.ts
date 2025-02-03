import useGetState from '../../../hooks/useGetElementState'
import {
	draggedElementIdentifier,
	hoveredElementIdentifier,
	shadowElementArray,
} from './dragAndDropState'
import { useAtom } from 'jotai'
import { swapBlocks } from './dragAndDropHelpers'

export default () => {
	const [draggedElement, setDraggedElement] = useAtom(
		draggedElementIdentifier
	)
	const [hoveredElement, setHoveredElement] = useAtom(
		hoveredElementIdentifier
	)

	const { elementOrderArray } = useGetState()
	const [shadowArray, setShadowArray] = useAtom(shadowElementArray)

	function onDrag(draggedElement: string) {
		syncShadowElementArrayWithTheRealElementArray()
		setDraggedElement(draggedElement)
	}

	function syncShadowElementArrayWithTheRealElementArray() {
		setShadowArray(elementOrderArray)
	}

	function swapElementsInShadowOrderArray(firstEl: string, secEl: string) {
		const newArr = swapBlocks(firstEl, secEl, shadowArray)
		console.log({ newArr })
		setShadowArray(newArr)
	}

	function onStopDrag() {
		console.log('drag stops')
		setDraggedElement(undefined)
	}

	return {
		onDrag,
		draggedElement,
		onStopDrag,
		swapElementsInShadowOrderArray,
	}
}
