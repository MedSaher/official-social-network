// app/layout.tsx
import './globals.css'
import { ReactNode } from 'react'
import { AuthProvider } from '../components/AuthContext'
import NavBar from '../components/NavBar'

export const metadata = {
  title: 'My App',
  description: 'A simple auth app',
}

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-green-50 min-h-screen flex flex-col items-center">
        <AuthProvider>
          <NavBar />
          <hr className="w-full border-green-300" />
          <main className="flex-grow w-full max-w-3xl p-6">
            {children}
          </main>
        </AuthProvider>
      </body>
    </html>
  )
}
