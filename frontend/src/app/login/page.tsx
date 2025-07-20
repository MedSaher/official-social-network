// app/register/page.tsx
'use client'; // only if you're using useState, useEffect, or other client-side hooks
import React from "react";
import LoginForm from '../../components/loginForm'
import useEffect from ''


export default function RegisterPage() {
   useEffect(() => {
      fetch('http://localhost:8080/api/check-session', {
        method: 'GET',
        credentials: 'include', // 🔥 This is CRUCIAL for cookies!
      })
        .then((res) => {
          if (res.status !== 200) {
            router.push('/')
          }
        })
        .catch(() => router.push('/login'))
    }, [])
  
  return (
    <main>
      <LoginForm />
    </main>
  );
}