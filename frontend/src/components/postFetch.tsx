// components/PostList.tsx
'use client';

import React, { useEffect, useState } from 'react';
import styles from './css/PostFetch.module.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCommentDots } from '@fortawesome/free-solid-svg-icons';
import CommentModal from './commentModal';

type Post = {
  id: number;
  user_id: number;
  group_id?: number;
  content: string;
  image_path?: string;
  privacy: 'public' | 'almost_private' | 'private';
  created_at: string;
  updated_at: string;
};

export default function PostList() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [activePostId, setActivePostId] = useState<number | null>(null);

  useEffect(() => {
    async function fetchPosts() {
      try {
        const res = await fetch('http://localhost:8080/api/posts/fetch_posts');
        if (!res.ok) {
          throw new Error(`Failed to fetch posts: ${res.status}`);
        }
        const data = await res.json();
        setPosts(data);
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message || 'Something went wrong');
        }
      } finally {
        setLoading(false);
      }
    }

    fetchPosts();
  }, []);

  if (loading) return <p className={styles.message}>Loading posts...</p>;
  if (error) return <p className={styles.error}>{error}</p>;
  if (posts === null) return <p className={styles.message}>No posts found.</p>;

  return (
    <div className={styles.postList}>
      {posts.map((post) => (
        <div key={post.id} className={styles.postCard}>
          <div className={styles.header}>
            <span className={styles.privacy}>
              Post #{post.id} · <strong>{post.privacy}</strong>
            </span>
            <span className={styles.timestamp}>
              {new Date(post.created_at).toLocaleString()}
            </span>
          </div>
          <p className={styles.content}>{post.content}</p>
          {post.image_path && (
            <img
              src={`http://localhost:8080/${post.image_path}`}
              alt="Post image"
              className={styles.image}
            />
          )}
          <div className={styles.iconRow}>
            <FontAwesomeIcon
              icon={faCommentDots}
              className={styles.commentIcon}
              onClick={() => setActivePostId(post.id)}
              title="Comment"
            />
          </div>
        </div>
      ))}

      {activePostId !== null && (
        <CommentModal postId={activePostId} onClose={() => setActivePostId(null)} />
      )}
    </div>
  );
}
