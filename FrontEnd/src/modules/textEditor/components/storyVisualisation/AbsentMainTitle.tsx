import useAddNewBlock from '../../hooks/useAddNewBlock'

const AbsentMainTitle = () => {
	const { addNewTextBlock } = useAddNewBlock()

	return <div>Please set your title</div>
}

export default AbsentMainTitle
