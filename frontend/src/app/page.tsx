'use client'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'

export default function HomePage() {
  const router = useRouter()
  const [isLoading, setIsLoading] = useState(true)
  

  useEffect(() => {
    const checkSession = async () => {
      try {
        const res = await fetch('http://localhost:8080/api/check-session', {
          method: 'GET',
          credentials: 'include',
        })

        if (res.status !== 200) {
          router.push('/login')
        } else {
          setIsLoading(false)
        }
      } catch (err) {
        router.push('/login')
      }
    }

    checkSession()
  }, [])

  const handleLogout = async () => {
    try {
      await fetch('http://localhost:8080/api/logout', {
        method: 'POST',
        credentials: 'include',
      })
      localStorage.removeItem('userId')
      router.push('/login')
    } catch (err) {
      console.error('Logout failed:', err)
    }
  }

  if (isLoading) {
    return <div>Loading...</div>
  }

  return (
    <div>
      <h1>Welcome to Home Page</h1>
      <button onClick={handleLogout}>Logout</button>
    </div>
  )
}
