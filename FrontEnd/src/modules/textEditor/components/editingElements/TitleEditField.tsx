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
		(currTitle: string) => {
			_debounce(async () => {
				const result = await setNewTitle(currTitle)
			}, 400)
		},
		[setNewTitle, title]
	)

	return (
		<input
			onChange={({ target: input }) => {
				const currInputFieldValue = input.value
				setTitle(currInputFieldValue)
				changeTitleGlobally(currInputFieldValue)
			}}
		/>
	)
}
