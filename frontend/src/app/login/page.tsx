'use client'
import React, { useEffect } from 'react'
import LoginForm from '../../components/loginForm'
import { useRouter } from 'next/navigation'

export default function LoginPage() {
  const router = useRouter()

  useEffect(() => {
    fetch('http://localhost:8080/api/check-session', {
      method: 'GET',
      credentials: 'include',
    })
      .then((res) => {
        if (res.status === 200) {
          // ✅ User is already logged in — redirect to home
          router.push('/')
        }
      })
      .catch(() => {
        // ❌ On error, we assume no session — stay on login
        console.log("stay in the login page.");
        
      })
  }, [])

  return (
    <main>
      <LoginForm />
    </main>
  )
}
