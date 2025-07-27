// components/groups/JoinRequestsClient.tsx
'use client';

import React, { useEffect, useState } from 'react';
import styles from './css/JoinRequestsClient.module.css';

interface JoinRequest {
  id: number;
  user_id: number;
  nick_name: string;
  user_name: string;
  created_at: string;
}

interface JoinRequestsClientProps {
  groupId: string;
}

export default function JoinRequestsClient({ groupId }: JoinRequestsClientProps) {
  const [requests, setRequests] = useState<JoinRequest[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [actionStatus, setActionStatus] = useState<string | null>(null);

  useEffect(() => {
    async function fetchPendingRequests() {
      try {
        const res = await fetch(`http://localhost:8080/api/groups/${groupId}/pending_requests`, {
          method: 'POST',
          credentials: 'include',
        });

        if (!res.ok) {
          throw new Error(`Failed to fetch: ${res.statusText}`);
        }

        const data = await res.json();
        setRequests(data);
      } catch (err: any) {
        setError(err.message || 'Unknown error');
      } finally {
        setLoading(false);
      }
    }

    fetchPendingRequests();
  }, [groupId]);

  async function handleRequestResponse(requestId: number, accept: boolean) {
    setActionStatus(null);
    try {
      const res = await fetch(`http://localhost:8080/api/groups/join_request/respond`, {
        method: 'POST',
        credentials: 'include',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          request_id: requestId,
          accept,
        }),
      });

      if (!res.ok) {
        throw new Error(`Failed to ${accept ? 'accept' : 'decline'} request: ${res.statusText}`);
      }

      setRequests((prev) => prev.filter((req) => req.id !== requestId));
      setActionStatus(`Request ${accept ? 'accepted' : 'declined'} successfully.`);
    } catch (err: any) {
      setActionStatus(err.message || 'Unknown error');
    }
  }

  if (loading) return <p>Loading requests...</p>;
  if (error) return <p style={{ color: 'red' }}>{error}</p>;
  if (requests.length === 0) return <p>No pending requests.</p>;

  return (
    <>
      <ul className={styles.list}>
        {requests.map((req) => (
          <li key={req.id} className={styles.listItem}>
            <div>
              <span className={styles.userInfo}>
                {req.nick_name} (@{req.user_name})
              </span>
              <span className={styles.timestamp}>
                {' '}
                requested at {new Date(req.created_at).toLocaleString()}
              </span>
            </div>
            <div className={styles.buttons}>
              <button
                className={`${styles.button} ${styles.acceptBtn}`}
                onClick={() => handleRequestResponse(req.id, true)}
              >
                Accept
              </button>
              <button
                className={`${styles.button} ${styles.declineBtn}`}
                onClick={() => handleRequestResponse(req.id, false)}
              >
                Decline
              </button>
            </div>
          </li>
        ))}
      </ul>
      {actionStatus && <p className={styles.statusMessage}>{actionStatus}</p>}
    </>
  );
}
