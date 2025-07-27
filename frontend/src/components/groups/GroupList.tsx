'use client';

import React, { useEffect, useState } from 'react';
import styles from './GroupList.module.css';

type Group = {
  id: number;
  creator_id: number;
  title: string;
  description: string;
  created_at: string;
};

export default function GroupList() {
  const [groups, setGroups] = useState<Group[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function fetchGroups() {
      try {
        const res = await fetch('http://localhost:8080/api/groups/fetch_groups', {
          method: 'GET',
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
          </div>
        ))
      )}
    </div>
  );
}
