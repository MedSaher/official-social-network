'use client'
import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'
import CreatePost from './postCreate'
import PostList from "./postFetch"
import CreateGroupForm from './groups/CreateGroupForm'
import GroupList from '@/components/groups/GroupList';

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
        <CreateGroupForm />
        <GroupList />
      </div>
    </>
  )
}
