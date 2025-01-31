import { ReactNode } from 'react'
import useDragAndDrop from './useDragAndDrop'

export default ({
	identifier,
	children,
}: {
	identifier: string
	children: ReactNode
}) => {
	const {
		onDrag,
		onStopDrag,
		draggedElement,
		swapElementsInShadowOrderArray,
	} = useDragAndDrop()

	return (
		<div
			onDragEnd={() => onStopDrag()}
			onDrag={(e) => onDrag(identifier)}
			className='w-[80%] p-2 bg-secSpecial flex flex-wrap justify-center items-center'
			onMouseOver={(e) => {
				if (draggedElement && draggedElement != identifier)
					swapElementsInShadowOrderArray(draggedElement, identifier)
			}}
		>
			{children}
		</div>
	)
}
