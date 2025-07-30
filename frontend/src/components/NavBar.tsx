'use client'
import styles from './css/HomePage.module.css'
import Link from 'next/link'
import { useState, useEffect } from 'react'
import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'

interface User {
  id: string;
  username: string;
  firstName?: string;
  lastName?: string;
  avatarUrl?: string;
}

export default function NavBar() {
  const router = useRouter()
  const { isAuthenticated, loading, logout, user } = useAuth()
  const [searchQuery, setSearchQuery] = useState('')
  const [searchResults, setSearchResults] = useState<User[]>([])

  useEffect(() => {
    const fetchUsers = async () => {
      if (searchQuery.length > 0) {
        try {
          const response = await fetch(`/api/search_users?q=${searchQuery}`)
          if (response.ok) {
            const data = await response.json()
            setSearchResults(data)
          } else {
            console.error('Failed to fetch users:', response.statusText)
            setSearchResults([])
          }
        } catch (error) {
          console.error('Error fetching users:', error)
          setSearchResults([])
        }
      } else {
        setSearchResults([])
      }
    }

    const delayDebounceFn = setTimeout(() => {
      fetchUsers()
    }, 300)

    return () => clearTimeout(delayDebounceFn)
  }, [searchQuery])

  const handleLogout = async () => {
    try {
      await logout()
      router.push('/login')
    } catch (error) {
      console.error('Failed to logout:', error)
    }
  }

  if (loading) return null

  return (
    <header className={styles.header}>
      <Link href="/" className={styles.logo}>
        SocialCircle
      </Link>

      {isAuthenticated && (
        <>
          <div className={styles.searchBar}>
            <input
              type="text"
              placeholder="Search posts, people, or groups..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
            <button className={styles.searchButton}>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                <path d="M15.5 14h-.79l-.28-.27a6.5 6.5 0 001.48-5.34c-.47-2.78-2.79-5-5.59-5.34a6.505 6.505 0 00-7.27 7.27c.34 2.8 2.56 5.12 5.34 5.59a6.5 6.5 0 005.34-1.48l.27.28v.79l4.25 4.25c.41.41 1.08.41 1.49 0 .41-.41.41-1.08 0-1.49L15.5 14zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z" />
              </svg>
            </button>

            {searchResults.length > 0 && (
              <div className={styles.searchResults}>
                {searchResults.map((user) => (
                  <Link
                    key={user.id}
                    href={`/profile/${user.username}`}
                    className={styles.searchResultItem}
                  >
                    {user.avatarUrl ? (
                      <img src={user.avatarUrl} alt={user.username} className={styles.searchResultAvatar} />
                    ) : (
                      <div className={styles.searchResultInitials}>
                        {user.username.charAt(0).toUpperCase()}
                      </div>
                    )}
                    <span>{user.username}</span>
                  </Link>
                ))}
              </div>
            )}
          </div>

          <div className={styles.userControls}>
            <button className={styles.notificationButton}>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 22c1.1 0 2-.9 2-2h-4c0 1.1.89 2 2 2zm6-6v-5c0-3.07-1.64-5.64-4.5-6.32V4c0-.83-.67-1.5-1.5-1.5s-1.5.67-1.5 1.5v.68C7.63 5.36 6 7.92 6 11v5l-2 2v1h16v-1l-2-2z" />
              </svg>
              <span className={styles.notificationBadge}>3</span>
            </button>

            <button className={styles.profileButton} onClick={() => router.push('/profile')}>
              {user?.avatarUrl ? (
                <img src={user.avatarUrl} alt="Profile" className={styles.profileImage} />
              ) : (
                <div className={styles.profileInitials}>
                  {user?.firstName?.charAt(0)}{user?.lastName?.charAt(0)}
                </div>
              )}
            </button>

            <button className={styles.logoutButton} onClick={handleLogout}>
              Logout
            </button>
          </div>
        </>
      )}

      {!isAuthenticated && (
        <div className={styles.authButtons}>
          <button
            className={styles.loginButton}
            onClick={() => router.push('/login')}
          >
            Login
          </button>
          <button
            className={styles.signupButton}
            onClick={() => router.push('/signup')}
          >
            Sign Up
          </button>
        </div>
      )}
    </header>
  )
}