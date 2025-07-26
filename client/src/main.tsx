import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { Toaster } from 'sonner'


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
    <Toaster
      theme="dark"
      position="top-right"
      expand={false}
      visibleToasts={1}
      richColors
      closeButton
    />
  </React.StrictMode>,
)