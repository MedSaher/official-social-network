'use client'
import { useRouter } from 'next/navigation'
import { useEffect } from 'react'
import axios from 'axios'

export default function HomePage() {
  const router = useRouter()

  const checkSession = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/check-session', {
        withCredentials: true, // ✅ Send session_token cookie
      })
      console.log('Session valid:', response.data)
    } catch (error) {
      console.error('Failed to check session:', error)
      router.push('/login')
    }
  }

  useEffect(() => {
    checkSession()
  }, [])

  const handleLogout = async () => {
    try {
      await axios.post('http://localhost:8080/api/logout', {}, {
        withCredentials: true, // ✅ Include session_token in logout request
      })
    } catch (error) {
      console.error('Failed to logout:', error)
    } finally {
      router.push('/login')
    }
  }

  return (
    <div>
      <h1>Welcome to Home Page</h1>
      <button onClick={handleLogout}>Logout</button>
    </div>
  )
}
