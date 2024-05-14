import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Home from './pages/Home'
import Page404 from './pages/404'
import Sake from './pages/SakePage'
import Login from './pages/Login'
import Register from './pages/Register'
import { Toaster } from 'react-hot-toast'
import Footer from './components/Footer'
import Navbar from './components/Navbar'
import PasswordReset from './pages/PasswordReset'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import UserInfo from './pages/UserInfo'
import { APIProvider } from '@vis.gl/react-google-maps'

// google map api
const googleMapAPIKey = process.env.REACT_APP_GOOGLE_MAP_API_KEY as string

// react-query setting
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5
    }
  }
})

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Toaster
          position="top-right"
          gutter={12}
          containerStyle={{ margin: '8px' }}
          toastOptions={{
            success: {
              duration: 3000
            },
            error: {
              duration: 5000
            },
            style: {
              fontSize: '16px',
              maxWidth: '500px',
              padding: '16px 24px',
              backgroundColor: '#333',
              color: '#fff'
            }
          }}
        />
        <Navbar />
        <APIProvider apiKey={googleMapAPIKey} language="zh-TW">
          <Routes>
            <Route index element={<Home />} />
            <Route path="items" element={<Sake />} />
            <Route path="login" element={<Login />} />
            <Route path="register" element={<Register />} />
            <Route path="password-reset" element={<PasswordReset />} />
            <Route path="user" element={<UserInfo />} />
            <Route path="*" element={<Page404 />} />
          </Routes>
        </APIProvider>
        <Footer />
      </BrowserRouter>
    </QueryClientProvider>
  )
}
