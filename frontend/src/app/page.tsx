'use client'
import { useEffect } from 'react'
import { useRouter } from 'next/navigation'

export default function HomePage() {
  const router = useRouter()

  // 🔐 Check for active session
  useEffect(() => {
    fetch('http://localhost:8080/api/check-session', {
      method: 'GET',
      credentials: 'include', // send cookie!
    })
      .then((res) => {
        if (res.status !== 200) {
          router.push('/login')
        }
      })
      .catch(() => router.push('/login'))
  }, [])

  // 🔓 Logout: clear session in DB + browser
  const handleLogout = async () => {
    try {
      await fetch('http://localhost:8080/api/logout', {
        method: 'POST',
        credentials: 'include', // includes session cookie
      })

      localStorage.removeItem('userId') // if you use it
      router.push('/login')
    } catch (err) {
      console.error('Logout failed:', err)
    }
  }

  return (
    <div>
      <h1>Welcome to Home Page</h1>
      <button onClick={handleLogout}>Logout</button>
    </div>
  )
}

