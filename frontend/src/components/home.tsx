'use client'

import axios from "axios"
import { useRouter } from 'next/navigation'
import CreatePost  from './postCreate'
import PostList from "./postFetch"

export default function RenderHomePage() {
    let router = useRouter()
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
      <CreatePost/>
      <PostList/>
    </div>
  )
}