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
export const elementOrderArrayAtom = atom<string[]>([])
export const elementMapAtom = atom<Map<string, string>>(
	new Map([['maiasdadsdas', 'titles']])
)
export const testMap = atom<Map<string, string>>(new Map())
