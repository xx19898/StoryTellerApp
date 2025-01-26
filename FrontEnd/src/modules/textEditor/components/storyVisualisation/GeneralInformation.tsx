// TODO: create component for showing the general information on the article: title osv
import { IoMdPerson } from 'react-icons/io'
import { ReactNode } from 'react'

interface IGeneralInformation {
	children: ReactNode
	creationDate: Date
	lastChangedDate: Date
	author: string
}

export default ({
	children,
	author,
	creationDate,
	lastChangedDate,
}: IGeneralInformation) => {
	return (
		<div className='relative flex w-full align-center justify-center h-auto min-h-5 p-0'>
			<div className='absolute right-0 bg-darkerSecondary flex items-center align-center flex-col p-4 opacity-60 hover:opacity-100 mr-4 mt-4'>
				<p>
					<b>Author:{author}</b>
				</p>
				<div className='bg-darkerPrimary rounded-[50%] w-20 h-20 flex justify-center items-center'>
					<IoMdPerson size={40} />
				</div>
				<p>Creation date: {creationDate.getFullYear()}</p>
				<p>Last modified: {lastChangedDate.getDate()}</p>
			</div>
			{children}
		</div>
	)
}
