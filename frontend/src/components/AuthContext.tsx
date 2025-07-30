// frontend/src/components/AuthContext.tsx
'use client'
import React, { createContext, useContext, useEffect, useState } from 'react'
import axios from 'axios'

const AuthContext = createContext<AuthContextType | null>(null)

interface AuthContextType {
  isAuthenticated: boolean;
  loading: boolean;
  login: () => void;
  logout: () => Promise<void>;
}

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false)
    const [loading, setLoading] = useState(true)

    const checkAuth = async () => {
        try {
            const res = await axios.get('http://localhost:8080/api/check-session', {
                withCredentials: true,
            })
            setIsAuthenticated(true)
        } catch (error: unknown) {
            setIsAuthenticated(false)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        checkAuth()
    }, [])

    const login = () => {
        setIsAuthenticated(true)
    }

    const logout = async () => {
        await axios.post('http://localhost:8080/api/logout', {}, { withCredentials: true })
        setIsAuthenticated(false)
    }

    return (
        <AuthContext.Provider value={{ isAuthenticated, loading, login, logout }}>
            {children}
        </AuthContext.Provider>
    )
}

export const useAuth = () => useContext(AuthContext) as AuthContextType
