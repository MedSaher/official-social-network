'use client'
import { useRouter } from 'next/navigation'
import { useEffect } from 'react'
import RenderHomePage from "../components/home"
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
  return (
    <main>
      <RenderHomePage />
    </main>
  );
}
