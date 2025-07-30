'use client'
import React, { createContext, useContext, useEffect, useState } from 'react'
import axios from 'axios'

// Define user type
interface User {
  id: number;
  username: string;
  firstName: string;
  lastName: string;
  email: string;
  avatarUrl?: string;
}

// Update AuthContextType to include user
interface AuthContextType {
  isAuthenticated: boolean;
  loading: boolean;
  user: User | null;
  login: (userData: User) => void;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | null>(null)

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false)
    const [loading, setLoading] = useState(true)
    const [user, setUser] = useState<User | null>(null)

    const checkAuth = async () => {
        try {
            const res = await axios.get('http://localhost:8080/api/check-session', {
                withCredentials: true,
            })
            setIsAuthenticated(true)
            setUser(res.data.user) // Assuming your API returns user data
        } catch (error: unknown) {
            setIsAuthenticated(false)
            setUser(null)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        checkAuth()
    }, [])

    const login = (userData: User) => {
        setIsAuthenticated(true)
        setUser(userData)
    }

    const logout = async () => {
        try {
            await axios.post('http://localhost:8080/api/logout', {}, { withCredentials: true })
        } finally {
            setIsAuthenticated(false)
            setUser(null)
        }
    }

    return (
        <AuthContext.Provider value={{ isAuthenticated, loading, user, login, logout }}>
            {children}
        </AuthContext.Provider>
    )
}

export const useAuth = () => {
    const context = useContext(AuthContext)
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider')
    }
    return context
}