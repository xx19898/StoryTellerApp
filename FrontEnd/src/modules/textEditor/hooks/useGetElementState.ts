import { useAtom } from 'jotai'
import { elementMapAtom, elementOrderArrayAtom } from '../textEditorState'

const useGetState = () => {
	const [elementOrderArray, setElementOrderArray] = useAtom(
		elementOrderArrayAtom
	)
	const [elementMap, setElementMap] = useAtom(elementMapAtom)

	return {
		elementOrderArray,
		setElementOrderArray,
		elementMap,
		setElementMap,
	}
}

export default useGetState
