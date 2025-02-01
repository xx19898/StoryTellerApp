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
			key={identifier}
			draggable={true}
			onDragEnd={() => onStopDrag()}
			onDragStart={(e) => {
				setTimeout(() => {
					onDrag(identifier)
				}, 0)
			}}
			className='w-[80%] flex flex-wrap justify-center items-center'
			style={{
				background: '#2A1E5C',
				padding: draggedElement ? '2em' : 0,
			}}
			onDragOver={(e) => {
				if (draggedElement && draggedElement != identifier)
					swapElementsInShadowOrderArray(draggedElement, identifier)
			}}
		>
			{children}
		</div>
	)

	//TODO: make styling depend on whether some element is dragged or not
}
