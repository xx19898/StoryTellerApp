import { atom } from 'jotai'

interface IEditedArticle {
	title: string
	rawData: string
	created: Date
}

export const generalInfo = atom<IEditedArticle>({
	title: 'none',
	rawData: '',
	created: new Date('2027/12/02'),
})
export const elementOrderArrayAtom = atom<string[]>([
	'first',
	'second',
	'third',
])
export const elementMapAtom = atom<Map<string, string>>(
	new Map([
		[
			'first',
			'<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean mollis nulla eu sem gravida sodales. Suspendisse quis velit eleifend, ornare enim quis, elementum arcu. Nunc rhoncus est et mauris interdum ornare. Mauris vestibulum iaculis ex vel facilisis. In eu malesuada ipsum. Duis non dui auctor, tempus nibh in, laoreet augue. Quisque porttitor enim diam, nec volutpat metus interdum porta. Proin orci lorem, tincidunt non odio nec, vulputate dictum nibh. Vivamus id ipsum dui. Vivamus vitae tempor felis. In fringilla quam eget gravida volutpat. Aenean porta finibus erat vel convallis. Integer.</p>',
		],
		[
			'second',
			'<p>Donec fringilla lobortis ligula feugiat commodo. Nullam auctor molestie leo. Aliquam tincidunt nulla id est aliquet sodales. In hac habitasse platea dictumst. Sed tempor volutpat urna ut commodo. Sed cursus magna mi, eget suscipit orci condimentum vitae. Praesent id ante id purus condimentum tempus. Proin non magna nec ante consectetur posuere. Maecenas hendrerit vehicula ligula, ut iaculis arcu suscipit non.</p>',
		],
		['third', '<p>third paragraph</p>'],
	])
)
export const testMap = atom<Map<string, string>>(new Map())
