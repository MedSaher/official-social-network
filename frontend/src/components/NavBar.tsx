'use client'
import './component.css/NavBar.css'
import Link from 'next/link'
import { useAuth } from './AuthContext'
import { useRouter } from 'next/navigation'

export default function NavBar() {
  const router = useRouter()
  const { isAuthenticated, loading, logout } = useAuth()

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
      <Link href="/" className="nav-link">Home</Link>

      {!loading && isAuthenticated && (
        <>
          <Link href="/profile" className="nav-link">Profile</Link>
          <button onClick={handleLogout} className="logout-btn">Logout</button>
        </>
      )}

      {!loading && !isAuthenticated && (
        <Link href="/login" className="nav-link">Login</Link>
      )}
    </nav>
  )
}
