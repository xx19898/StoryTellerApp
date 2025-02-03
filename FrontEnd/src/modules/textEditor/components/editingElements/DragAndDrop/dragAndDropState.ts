import { atom } from 'jotai'

export const draggedElementIdentifier = atom<undefined | string>(undefined)
export const hoveredElementIdentifier = atom<undefined | string>(undefined)

export const shadowElementArray = atom<string[]>([])
