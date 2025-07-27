'use client';

import React, { useEffect, useState } from 'react';
import styles from './GroupList.module.css';
import { useRouter } from 'next/navigation';

type Group = {
  id: number;
  creator_id: number;
  title: string;
  description: string;
  created_at: string;
  updated_at: string;
  is_creator: boolean;
  is_member: boolean;
};

export default function GroupList() {
  const [groups, setGroups] = useState<Group[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    async function fetchGroups() {
      try {
        const res = await fetch('http://localhost:8080/api/groups/fetch_groups', {
          credentials: 'include',
        });

        if (!res.ok) {
          throw new Error('Failed to fetch groups');
        }

        const data = await res.json();
        setGroups(data);
      } catch (err: any) {
        console.error(err);
        setError(err.message || 'Error loading groups');
      } finally {
        setLoading(false);
      }
    }

    fetchGroups();
  }, []);

  const handleJoin = async (groupId: number) => {
    try {
      const res = await fetch(`http://localhost:8080/api/groups/${groupId}/join`, {
        method: 'POST',
        credentials: 'include',
      });

      if (!res.ok) {
        alert(res.status)
        throw new Error('Failed to join group');
      }

      alert('Join request sent');
      // Optionally re-fetch or update UI
    } catch (err) {
      console.error(err);
      alert('Failed to join group');
    }
  };

  if (loading) return <p className={styles.message}>Loading groups...</p>;
  if (error) return <p className={styles.error}>{error}</p>;

  return (
    <div className={styles.groupList}>
      <h2 className={styles.heading}>All Groups</h2>

      {groups.length === 0 ? (
        <p className={styles.message}>No groups found.</p>
      ) : (
        groups.map((group) => (
          <div key={group.id} className={styles.groupCard}>
            <div className={styles.groupTitle}>{group.title}</div>
            <div className={styles.groupMeta}>
              Created by User #{group.creator_id} ·{' '}
              {new Date(group.created_at).toLocaleString()}
            </div>
            <div className={styles.groupDescription}>{group.description}</div>

            <div className={styles.actions}>
              {group.is_member ? (
                <button
                  className={styles.viewBtn}
                  onClick={() => router.push(`/groups/${group.id}`)}
                >
                  View Group
                </button>
              ) : (
                <button
                  className={styles.joinBtn}
                  onClick={() => handleJoin(group.id)}
                >
                  Join
                </button>
              )}
            </div>
          </div>
        ))
      )}
    </div>
  );
}
