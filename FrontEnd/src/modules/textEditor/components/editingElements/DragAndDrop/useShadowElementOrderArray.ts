import { useState } from 'react'
import useGetState from '../../../hooks/useGetElementState'
import { swapBlocks } from './dragAndDropHelpers'

// used when user drags one element over some other element but does not let go of the mousekey. In this way he sees how the changed piece looks
const useShadowElementOrderArray = () => {
	const { elementOrderArray } = useGetState()
	const [shadowElements, setShadowElements] = useState(elementOrderArray)

	function swapElementsInShadowOrderArray(firstEl: string, secEl: string) {
		const newArr = swapBlocks(firstEl, secEl, shadowElements)
		setShadowElements(newArr)
	}

	return { shadowElements, setShadowElements }
}

export default useShadowElementOrderArray
