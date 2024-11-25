import { useState } from 'react'
import { deleteEmptyElements } from '../helpers/HtmlParsingUtilities'
import useGetState from './useGetElementState'

function useRemoveEmptyElements() {
	const {
		elementMap,
		elementOrderArray,
		setElementMap,
		setElementOrderArray,
	} = useGetState()

	//console.log('%c useRemoveEmptyElements updated','color: violet;')
	//console.log({ elementMap })

	function removeEmptyElements() {
		//console.log({ originalElementMap: elementMap })
		const newElementMap = new Map(elementMap)
		//console.log({ copyOfOriginalMap: newElementMap })
		const newArr = [...elementOrderArray]
		const purifiedArr = deleteEmptyElements(newElementMap, newArr)
		setElementOrderArray(purifiedArr)
		setElementMap(newElementMap)
	}

	return { removeEmptyElements }
}

export default useRemoveEmptyElements
