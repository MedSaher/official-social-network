'use client'
import { useEffect } from 'react'
import { useRouter } from 'next/navigation'

export default function HomePage() {
  const router = useRouter()

  useEffect(() => {
    fetch('http://localhost:8080/api/check-session', {
      method: 'GET',
      credentials: 'include', // 🔥 This is CRUCIAL for cookies!
    })
      .then((res) => {
        if (res.status !== 200) {
          router.push('/login')
        }
      })
      .catch(() => router.push('/login'))
  }, [])

  const handleLogout = () => {
    localStorage.removeItem('userId') // Optional if you're also storing it
    router.push('/login')
  }

  return (
    <div>
      <h1>Welcome to Home Page</h1>
      <button onClick={handleLogout}>Logout</button>
    </div>
  )
}
