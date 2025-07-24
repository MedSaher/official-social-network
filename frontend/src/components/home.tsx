'use client'
import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'

export default function RenderHomePage() {
  const router = useRouter()
  const {logout } = useAuth()
  const handleLogout = async () => {
    try {
      await logout()
      router.push('/login')
    } catch (error) {
      console.error('Failed to logout:', error)
    }
  }
  return (
    <>
      <div>
        <h1>Welcome to Home Page</h1>
      </div>
    </>
  )
}
