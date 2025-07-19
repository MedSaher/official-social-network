// app/layout.tsx
import './globals.css'
import { ReactNode } from 'react'

export const metadata = {
  title: 'My App',
  description: 'A simple auth app',
}

// the rendered layout that holds evergything:
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <nav>
          <a href="/">Home</a> | <a href="/login">Login</a> | <a href="/register">Register</a>
        </nav>
        <hr />
        <main>{children}</main>
      </body>
    </html>
  )
}

