'use client';

import React, { useEffect, useState } from 'react';
import './user.css'; 
interface User {
  id: number;
  username: string;
  avatarUrl: string;
}

export default function AllUsers() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchAllUsers = async () => {
      try {
        const res = await fetch('/api/users');
        if (!res.ok) throw new Error('Failed to fetch users');
        const data = await res.json();
        setUsers(data);
      } catch (err) {
        console.error('Error:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchAllUsers();
  }, []);

  if (loading) return <div>Loading users...</div>;
  if (users.length === 0) return <div>No users found.</div>;

  return (
    <div className="user-list" style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
      {users.map((user) => (
        <div
          key={user.id}
          style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}
        >
          <img
            src={user.avatarUrl}
            alt={user.username}
            width={32}
            height={32}
            style={{ borderRadius: '50%' }}
          />
          <span>{user.username}</span>
        </div>
      ))}
    </div>
  );
}
