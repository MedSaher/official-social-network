'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useAuth } from './AuthContext'
import axios from 'axios'
import './component.css/login-form.css' // 👈 Import the CSS

export default function LoginForm() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const router = useRouter()
  const { login } = useAuth()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError('')
    let creds = {
      email: email,
      password: password
    }
    console.log(creds);

    try {
      await axios.post('http://localhost:8080/api/login',
        creds, {
        withCredentials: true // important to send and receive cookies:
      })
      router.push('/')
    } catch (err) {
      setError('Login failed: Invalid email or password')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="login-container">
      <h2 className="login-title">Login</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          value={email}
          onChange={e => setEmail(e.target.value)}
          placeholder="Email"
          required
          className="login-input"
        />
        <input
          type="password"
          value={password}
          onChange={e => setPassword(e.target.value)}
          placeholder="Password"
          required
          className="login-input"
        />
        <button
          type="submit"
          disabled={loading}
          className="login-button"
        >
          {loading ? 'Logging in...' : 'Login'}
        </button>
      </form>
      {error && <p className="login-error">{error}</p>}
      <p className="login-footer">
        Don't have an account?{' '}
        <a href="/register">Register</a>
      </p>
    </div>
  )
}
