'use client'

import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'
import CreatePost from './postCreate'
import PostList from "./postFetch"  // If you want to list posts too

export default function RenderHomePage() {
  const router = useRouter()
  const { logout } = useAuth()

  const handleLogout = async () => {
    try {
      await logout()
      router.push('/login')
    } catch (error) {
      console.error('Failed to logout:', error)
    }
  }

  return (
    <>
      <div>
        <h1>Welcome to Home Page</h1>
        {/* Add the CreatePost component here */}
        <CreatePost />
        {/* Optionally show posts */}
        {/* <PostList /> */}
      </div>
    </>
  )
}
