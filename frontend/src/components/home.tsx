'use client'
import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'
import CreatePost from './postCreate'
import PostList from "./postFetch"
import CreateGroupForm from './groups/CreateGroupForm'
import GroupList from '@/components/groups/GroupList';
import styles from './css/home.module.css';
import AllUsers from './user/fetchuser';
import { useState } from 'react';

export default function RenderHomePage() {
  const router = useRouter()
  const { logout } = useAuth()
  const [showCreatePostModal, setShowCreatePostModal] = useState(false);
  const [refreshPosts, setRefreshPosts] = useState(false);

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
      <div className={styles.main}>
        <div className={styles.groups}>
          <GroupList />
        </div>
        <div className='posts'>
          <CreatePost />
          <PostList />

        </div>
        <div className={styles.chat}>
          <AllUsers />
        </div>
      </div>
    </>
  )
}
