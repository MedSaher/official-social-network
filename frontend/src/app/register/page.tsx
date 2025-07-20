// app/register/page.tsx
'use client'; // only if you're using useState, useEffect, or other client-side hooks
import React from "react";
import RegistrationForm from '../../components/registrationForm'


export default function RegisterPage() {
  return (
    <main>
      <h1>Register</h1>
      <RegistrationForm />
    </main>
  );
}
