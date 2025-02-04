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
		<li
			key={identifier}
			draggable={true}
			onDragEnd={(e) => {
				e.preventDefault()
				console.log('drag stop')
				onStopDrag()
			}}
			onMouseUp={(e) => 'mouse released'}
			onDragStart={(e) => {
				//hack for some browsers, otherwise drag behaviour is inconsistent due to the drag event changing html/css/(eg brave browser)
				e.dataTransfer.setData('text/plain', identifier)
				setTimeout(() => {
					onDrag(identifier)
				}, 0)
			}}
			className='w-[80%] flex flex-wrap justify-center items-center'
			onDrop={(e) => console.log('drop')}
			style={{
				background: '#2A1E5C',
				padding: draggedElement ? '2em' : 0,
			}}
			onDragOver={() => {
				if (draggedElement && draggedElement != identifier)
					swapElementsInShadowOrderArray(draggedElement, identifier)
			}}
		>
			{children}
		</li>
	)

	//TODO: make styling depend on whether some element is dragged or not
}
