// app/layout.tsx
import './globals.css'
import { ReactNode } from 'react'

export const metadata = {
  title: 'My App',
  description: 'A simple auth app',
}

// the rendered layout that holds evergything:
// app/layout.tsx


export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-green-50 min-h-screen flex flex-col items-center">
        <nav className="w-full bg-white shadow-md py-4 px-6 flex justify-center space-x-6 sticky top-0 z-10">
          <a href="/" className="text-green-800 hover:text-green-600 font-semibold">Home</a>
          <a href="/login" className="text-green-800 hover:text-green-600 font-semibold">Login</a>
          <a href="/register" className="text-green-800 hover:text-green-600 font-semibold">Register</a>
        </nav>
        <hr className="w-full border-green-300" />
        <main className="flex-grow w-full max-w-3xl p-6">
          {children}
        </main>
      </body>
    </html>
  )
}


