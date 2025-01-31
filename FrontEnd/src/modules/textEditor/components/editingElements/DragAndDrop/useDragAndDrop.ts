import useGetState from '../../../hooks/useGetElementState'
import {
	draggedElementIdentifier,
	hoveredElementIdentifier,
	shadowElementArray,
} from './dragAndDropAtoms'
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

	function hoverOver(draggedElement: string) {}

	function onDrag(draggedElement: string) {
		syncShadowElementArrayWithTheRealElementArray
		setDraggedElement(draggedElement)
	}

	function syncShadowElementArrayWithTheRealElementArray() {
		setShadowArray(elementOrderArray)
	}

	function swapElementsInShadowOrderArray(firstEl: string, secEl: string) {
		const newArr = swapBlocks(firstEl, secEl, shadowArray)
		setShadowArray(elementOrderArray)
	}

	function onStopDrag() {
		setDraggedElement(undefined)
	}

	return {
		onDrag,
		draggedElement,
		onStopDrag,
		swapElementsInShadowOrderArray,
	}
}
