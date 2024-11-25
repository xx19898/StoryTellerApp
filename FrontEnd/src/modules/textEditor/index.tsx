import StoryEditor from './components/storyVisualisation/StoryEditor'

export const TextEditor = () => {
	return (
		<div
			className='w-full py-2 px-4 min-h-screen h-auto
            flex flex-col gap-[3.5em] justify-start items-center
            text-white'
		>
			<StoryEditor />
		</div>
	)
}
