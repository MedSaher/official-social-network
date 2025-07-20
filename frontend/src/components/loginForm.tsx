'use client'

import {useState} from 'react'
import axios from 'axios'

import {useRouter} from 'next/navigation'

export default function LoginForm(){
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState('')
    const router = useRouter()

const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError('')
    try{
        await axios.post('http://localhost:8080/login', {
            email,
            password
        }, {
            withCredentials: true // important to send and receive cookies:
        })
         router.push('/')
    }catch (err) {
      setError('Login failed: Invalid email or password')
    } finally {
      setLoading(false)
    }
}
  return (
    <div className="max-w-md mx-auto p-6 border rounded-md shadow-md">
      <h2 className="text-2xl mb-4">Login</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          value={email}
          onChange={e => setEmail(e.target.value)}
          placeholder="Email"
          required
          className="w-full p-2 mb-3 border rounded"
        />
        <input
          type="password"
          value={password}
          onChange={e => setPassword(e.target.value)}
          placeholder="Password"
          required
          className="w-full p-2 mb-3 border rounded"
        />
        <button
          type="submit"
          disabled={loading}
          className="w-full bg-blue-600 text-white p-2 rounded disabled:opacity-50"
        >
          {loading ? 'Logging in...' : 'Login'}
        </button>
      </form>
      {error && <p className="mt-3 text-red-600">{error}</p>}
      <p className="mt-4">
        Don't have an account? <a href="/register" className="text-blue-600 underline">Register</a>
      </p>
    </div>
  )
}