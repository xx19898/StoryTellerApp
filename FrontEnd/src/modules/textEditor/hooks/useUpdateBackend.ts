import { buildHtmlString } from '../helpers/HtmlParsingUtilities'
import useGetState from './useGetElementState'
import useUpdateStatus from './useUpdateStatus'

export default () => {
	const { elementMap, elementOrderArray } = useGetState()
	const { onUpdatePending, onNullifyUpdate } = useUpdateStatus()

	async function updateBackend() {
		const updatedStory = buildHtmlString(elementMap, elementOrderArray)
		onUpdatePending()
		setTimeout(() => {
			console.log({ updatedStory })
			onNullifyUpdate()
		}, 5000)
	}

	return { updateBackend }
}
