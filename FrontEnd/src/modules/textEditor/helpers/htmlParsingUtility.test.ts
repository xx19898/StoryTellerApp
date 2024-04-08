import { extractTypeAndContentOfHtmlElement } from './HtmlParsingElementUtilities'
import {addHtmlElementIdentifier, addNewParagraph, buildHtmlString, deleteEmptyElements, processHtmlString, setNewTitle} from './HtmlParsingUtilities'

import { expect, test, describe } from 'vitest'


describe('Testing html string parsing utility', () => {
  test('Element, type and contents get correctly extracted from html element string', () => {
    const testHtmlString = '<h2>Test Title</h2>'
    const  {
      contents,
      element,
      elementType
    } = extractTypeAndContentOfHtmlElement(testHtmlString)

    expect(contents).toBe('Test Title')
    expect(element).toBe('<h2>Test Title</h2>')
    expect(elementType).toBe('title')
  })

  test('Html string gets partitioned into individual strings correctly',() => {
    const testString = '<h2>Test Title</h2><p>Test Paragraph</p><img src="image_1"></img>'

    const {htmlElementMap,htmlOrderArray} = processHtmlString(testString)

    const title = htmlElementMap.get(htmlOrderArray[0])

    const paragraph = htmlElementMap.get(htmlOrderArray[1])

    const image = htmlElementMap.get(htmlOrderArray[2])

    expect(title).toBe('<h2>Test Title</h2>')
    expect(paragraph).toBe('<p>Test Paragraph</p>')
    expect(image).toBe('<img src="image_1"></img>')

  })

  test('html element identifier gets joined with content properly',() => {
    const testTitle = 'Test Title'
    const testParagraph = 'Test Paragraph'

    const titleElement = addHtmlElementIdentifier('h2',testTitle)
    expect(titleElement).toBe('<h2>Test Title</h2>')

    const paragraphElement = addHtmlElementIdentifier('p',testParagraph)

    expect(paragraphElement).toBe('<p>Test Paragraph</p>')
  })

  test('add new paragraph functions as it should',() => {
    const testParagraph = 'Test Paragraph'

    const map = new Map()
    const arr = ['']

    addNewParagraph(testParagraph,map,arr)

    expect(arr.length).toBe(2)

    const paragraphElement = map.get(arr[1])
    expect(paragraphElement).toBe('<p>Test Paragraph</p>')
  })

  test('Setting new title works',() => {
    const testTitle = 'Test Title'

    const map = new Map()
    const arr = ['']

    setNewTitle(testTitle,map,arr)

    expect(arr.length).toBe(1)

    const titleElement = map.get(arr[0])
    expect(titleElement).toBe('<h2>Test Title</h2>')
  })

  test('Html string gets built correctly(title and paragraphs)',() => {
    const testTitle = 'Test Title'
    const testParagraph = 'Test Paragraph'

    const map = new Map()
    const arr = ['']

    addNewParagraph(testParagraph,map,arr)
    setNewTitle(testTitle,map,arr)

    const titleElement = map.get(arr[0])
    const paragraphElement = map.get(arr[1])

    const properString = '<h2>Test Title</h2><p>Test Paragraph</p>'

    const resultingString = buildHtmlString(map,arr)

    expect(resultingString).toBe(properString)
  })

  test('Function for deleting empty html elements(paragraphs,titles,images with zero content)', () => {
    const testTitle = 'Test Title'
    const testParagraph = 'Test Paragraph'
    const emptyParagraph = ''

    const map = new Map()
    const arr = ['']

    addNewParagraph(testParagraph,map,arr)
    addNewParagraph(emptyParagraph,map,arr)
    setNewTitle(testTitle,map,arr)
    
    const arrWithNoEmptyElements = deleteEmptyElements(map,arr)
    expect(arrWithNoEmptyElements.length).toEqual(2)
  })
})
