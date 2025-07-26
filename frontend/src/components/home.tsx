'use client'
import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'
import CreatePost  from './postCreate'
import PostList from "./postFetch"

export default function RenderHomePage() {
  const router = useRouter()
  const {logout } = useAuth()
  const handleLogout = async () => {
    try {
      await logout()
      router.push('/login')
    } catch (error) {
      console.error('Failed to logout:', error)
    }
  }
  return (
<<<<<<< HEAD
    <div>
      <h1>Welcome to Home Page</h1>
      <button onClick={handleLogout}>Logout</button>
      <CreatePost/>
      <PostList/>
    </div>
=======
    <>
      <div>
        <h1>Welcome to Home Page</h1>
      </div>
    </>
>>>>>>> origin/fix-frontend-errs
  )
}
