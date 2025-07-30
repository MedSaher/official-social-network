'use client'
import './component.css/NavBar.css'
import Link from 'next/link'
import { useState, useEffect } from 'react'
import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'

interface User {
  id: string;
  username: string;
}

export default function NavBar() {
  const router = useRouter()
  const { isAuthenticated, loading, logout } = useAuth()
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

  return (
    <nav className="navbar">

      {!loading && isAuthenticated && (
        <>
          <Link href="/" className="nav-link">Home</Link>
          <Link href="/profile" className="nav-link">Profile</Link>
          <Link href="/groups" className="nav-link">Groups</Link>
          <Link href="/chat" className="nav-link">Chat</Link>
          <div className="search-bar">
            <input
              type="text"
              placeholder="Search users..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
            {searchResults.length > 0 && (
              <ul className="search-results">
                {searchResults.map((user) => (
                  <li key={user.id}>
                    <Link href={`/profile/${user.username}`}>
                      {user.username}
                    </Link>
                  </li>
                ))}
              </ul>
            )}
          </div>
          <button onClick={handleLogout} className="logout-btn">Logout</button>

        </>
      )}

      {!loading && !isAuthenticated && (
        <Link href="/login" className="nav-link">Login</Link>
      )}
    </nav>
  )
}
