'use client';

import React, { useEffect, useState } from 'react';
import JoinRequestsClient from './JoinRequestsClient';

interface Props {
  groupId: string;
}

export default function GroupAdminWrapper({ groupId }: Props) {
  const [role, setRole] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function checkRole() {
      try {
        const res = await fetch(`http://localhost:8080/api/groups/${groupId}/member_role`, {
          method: 'POST',
          credentials: 'include',
        });
        if (!res.ok) throw new Error('Unauthorized or not a member');

        const data = await res.json();
        setRole(data.role); // 'creator', 'admin', 'member', or 'none'
      } catch (err) {
        setRole('none');
      } finally {
        setLoading(false);
      }
    }

    checkRole();
  }, [groupId]);

  if (loading) return <p>Checking permissions...</p>;
  if (role !== 'creator' && role !== 'admin') return null;

  return (
    <section id="requests">
      <h2>Join Requests</h2>
      <JoinRequestsClient groupId={groupId} />
    </section>
  );
}
