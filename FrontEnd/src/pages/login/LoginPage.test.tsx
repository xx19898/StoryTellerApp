import { expect, test, describe,vi,afterEach } from 'vitest'
import {cleanup, fireEvent, render,screen, waitFor} from '@testing-library/react'
import LoginPage from './loginPage'

describe('Testing the Login page', () => {
    afterEach(() => {
        cleanup()
    })
    test('Username and Password inputs are defined',async () => {
        render(<LoginPage error='' isLoading={false} login={() => console.log('')} loginSuccess={false} onSuccess={() => console.log('')}/>)

        const usernameInput = await waitFor(() => screen.findByTestId('username-input'))
        const passwordInput = await waitFor(() => screen.findByTestId('password-input'))

        expect(usernameInput).toBeDefined()
        expect(passwordInput).toBeDefined()
    })

    test('Login func gets called with proper parameters', async () => {
        const mockFn = vi.fn((username:string,password:string) => 0)

        render(<LoginPage error='' isLoading={false} login={(username:string,password:string) => mockFn(username,password)} loginSuccess={false} onSuccess={() => console.log('')}/>)

        const usernameInput = await waitFor(() => screen.findByTestId('username-input'))
        const passwordInput = await waitFor(() => screen.findByTestId('password-input'))

        expect(usernameInput).toBeDefined()
        expect(passwordInput).toBeDefined()

        fireEvent.change(usernameInput,{target:{value:'Username'}})
        fireEvent.change(passwordInput,{target:{value:'Password'}})

        expect(usernameInput.value).toEqual('Username')
        expect(passwordInput.value).toEqual('Password')

        const signInButton = await waitFor(() => screen.findByText('Sign In'))

        expect(signInButton).toBeDefined()

        fireEvent(signInButton, new MouseEvent('click',{
            bubbles: true,
            cancelable: true,
        }))

        expect(mockFn).toHaveBeenCalledOnce()
        expect(mockFn).toHaveBeenCalledWith('Username','Password')
    })
})
