import { IoIosSync } from 'react-icons/io'
import gsap from 'gsap'
import { useGSAP } from '@gsap/react'
import { useRef } from 'react'

export default () => {
	const spinnerRef = useRef<HTMLDivElement>(null)
	useGSAP(() => {
		gsap.to(spinnerRef.current, { rotation: 180, repeat: -1, duration: 3 })
	})

	return (
		<div ref={spinnerRef}>
			<IoIosSync size={50} />
		</div>
	)
}
