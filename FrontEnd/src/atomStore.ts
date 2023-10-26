import {createStore,atom} from 'jotai'

export const myStore = createStore()

export const authTokenAtom = atom<string | undefined>(undefined)