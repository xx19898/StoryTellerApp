import {extractTypeAndContentOfHtmlElement} from './htmlParsingUtilities'

import { expect, test, describe } from 'vitest'


describe('Testing html string parsing utility', () => {
  test('Element, type and contents get correctly extracted from html string', () => {
    const testHtmlString = '<h1>Test Title</h1>'
    const  {
      contents,
      element,
      elementType
    } = extractTypeAndContentOfHtmlElement(testHtmlString)

    console.log({contents,element,elementType})

    expect(contents).toBe('Test Title')
    expect(element).toBe('<h1>Test Title</h1>')
    expect(elementType).toBe('h1')
  })
})
