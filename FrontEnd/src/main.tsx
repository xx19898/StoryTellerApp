import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import {RouterProvider} from 'react-router-dom'
import router from './router'
import { QueryClient,QueryClientProvider} from 'react-query'
import { Provider } from 'jotai'
import { myStore } from './atomStore'

const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Provider store={myStore}>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router}/>
    </QueryClientProvider>
    </Provider>
  </React.StrictMode>,
)
