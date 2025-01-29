import { useCallback, useState } from 'react'
import { useDebounce } from 'use-debounce'
import _debounce from 'lodash.debounce'

interface ITitleEditField {
	setNewTitle: (newTitle: string) => Promise<{ status: string }>
	initialTitle: string
}

export default ({ initialTitle, setNewTitle }: ITitleEditField) => {
	const [title, setTitle] = useState(initialTitle)

	const changeTitleGlobally = useCallback(
		_debounce(async (currTitle: string) => {
			console.log('running debounce')
			const result = await setNewTitle(currTitle)
			console.log('CHANGED TITLE ON REMOTE')
		}, 1000),
		[]
	)

	return (
		<input
			onChange={({ target: input }) => {
				const currInputFieldValue = input.value
				console.log({ currInputFieldValue })
				setTitle(currInputFieldValue)
				changeTitleGlobally(currInputFieldValue)
			}}
		/>
	)
}
