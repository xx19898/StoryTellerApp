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
export const elementOrderArrayAtom = atom<string[]>(['first', 'second'])
export const elementMapAtom = atom<Map<string, string>>(
	new Map([
		['first', '<p>first<p>'],
		['second', '<p>second</p>'],
	])
)
export const testMap = atom<Map<string, string>>(new Map())
