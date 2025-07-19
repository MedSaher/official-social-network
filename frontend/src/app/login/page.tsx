'use client'
import { useState } from 'react'
import axios from 'axios'
import { useRouter } from 'next/navigation'

export default function LoginPage() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const router = useRouter()

  const handleLogin = async () => {
    try {
      const res = await axios.post('http://localhost:8080/login', {
        email,
        password
      })
      localStorage.setItem('userId', res.data.userId)
      router.push('/')
    } catch (err) {
      alert('Login failed')
    }
  }

  return (
    <div>
      <h2>Login</h2>
      <input value={email} onChange={e => setEmail(e.target.value)} placeholder="Email" /><br />
      <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" /><br />
      <button onClick={handleLogin}>Login</button>
      <p>Don&apos;t have an account? <a href="/register">Register</a></p>
    </div>
  )
}
