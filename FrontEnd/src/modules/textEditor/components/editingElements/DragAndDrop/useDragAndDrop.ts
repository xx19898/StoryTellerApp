import { useState } from 'react'
import useGetState from '../../../hooks/useGetElementState'

export default () => {
	const [dragAndDropActive, setDragAndDropActive] = useState(false)

	const {
		elementMap,
		elementOrderArray,
		setElementMap,
		setElementOrderArray,
	} = useGetState()
}
