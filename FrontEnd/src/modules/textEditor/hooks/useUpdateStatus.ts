import { useAtom } from 'jotai'
import { updateStatusAtom } from '../textEditorState'

export default () => {
	const [updateStatus, setUpdateStatus] = useAtom(updateStatusAtom)

	function onUpdateSuccessful(successMessage: string) {
		setUpdateStatus({ status: 'SUCCESS', successMessage })
	}
	function onUpdateFailure(failureMessage: string) {
		setUpdateStatus({ status: 'FAILURE', errorMessage: failureMessage })
	}
	function onUpdatePending() {
		setUpdateStatus({ status: 'PENDING' })
	}
	function onNullifyUpdate() {
		setUpdateStatus({ status: 'NO_UPDATE' })
	}

	return {
		updateStatus,
		onUpdateSuccessful,
		onUpdateFailure,
		onUpdatePending,
		onNullifyUpdate,
	}
}
