import { expect, test, describe } from 'vitest'
import {fireEvent, render,screen, waitFor} from '@testing-library/react'
import LoginPage from './loginPage'

describe('Testing the Login page', () => {
    test.only('Username and Password labels exist',async () => {
        render(<LoginPage error='' isLoading={false} login={() => console.log('')} loginSuccess={false} onSuccess={() => console.log('')}/>)

        const usernameLabel = await waitFor(() => screen.findByText('Username'))
        const passwordLabel = await waitFor(() => screen.findByText('Password'))

        expect(usernameLabel).toBeDefined()
        expect(passwordLabel).toBeDefined()
    })

    test('Login func gets called with proper parameters', async () => {
        const mockCallback = jest.fn((username:string,password:string) => username + password)

        render(<LoginPage error='' isLoading={false} login={mockCallback} loginSuccess={false} onSuccess={() => console.log('')}/>)

        const usernameInput = await waitFor(() => screen.findByLabelText('Username'))
        const passwordInput = await waitFor(() => screen.findByLabelText('Password'))

        expect(usernameInput).toBeDefined()
        expect(passwordInput).toBeDefined()

        fireEvent.change(usernameInput,{target:{value:'User'}})
        fireEvent.change(passwordInput,{target:{value:'Password'}})

        const signInButton = await waitFor(() => screen.findByText('Sign In'))

        fireEvent(signInButton, new MouseEvent('click'))

        expect(mockCallback).toHaveBeenCalledWith({Username:'User',Password:'Password'})
    })
})
