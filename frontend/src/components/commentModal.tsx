import React, { useEffect, useState } from 'react';
import styles from './css/CommentModal.module.css';

interface Comment {
  id: number;
  post_id: number;
  user_id: number;
  content: string;
  created_at: string;
}

interface CommentModalProps {
  postId: number;
  onClose: () => void;
}

// Replace this with your actual logged-in user ID (or get it from context/session)
const CURRENT_USER_ID = 2;

export default function CommentModal({ postId, onClose }: CommentModalProps) {
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState(true);
  const [newComment, setNewComment] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    async function fetchComments() {
      try {
        const res = await fetch(`http://localhost:8080/api/posts/fetch_comments?post_id=${postId}`);
        if (!res.ok) {
          throw new Error('Failed to fetch comments');
        }
        const data = await res.json();
        setComments(data);
      } catch (err) {
        console.error(err);
        setError('Failed to load comments');
      } finally {
        setLoading(false);
      }
    }

    fetchComments();
  }, [postId]);

  async function handleSubmit() {
    if (!newComment.trim()) {
      setError('Comment cannot be empty');
      return;
    }

    setError(null);
    setSubmitting(true);

    try {
      const res = await fetch('http://localhost:8080/api/posts/create_comment', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          post_id: postId,
          user_id: CURRENT_USER_ID,
          content: newComment.trim(),
        }),
      });

      if (!res.ok) {
        throw new Error('Failed to submit comment');
      }

      const createdComment = await res.json();

      // Append the new comment to the list
      setComments((prev) => [...prev, createdComment]);
      setNewComment('');
    } catch (err) {
      console.error(err);
      setError('Failed to submit comment');
    } finally {
      setSubmitting(false);
    }
  }

  return (
    <div className={styles.overlay}>
      <div className={styles.modal}>
        <h2>Comments on Post #{postId}</h2>

        {loading ? (
          <p>Loading comments...</p>
        ) : comments.length === 0 ? (
          <p>No comments yet.</p>
        ) : (
          <div className={styles.commentList}>
            {comments.map((c) => (
              <div key={c.id} className={styles.commentItem}>
                <div className={styles.commentMeta}>
                  User #{c.user_id} · {new Date(c.created_at).toLocaleString()}
                </div>
                <div className={styles.commentContent}>{c.content}</div>
              </div>
            ))}
          </div>
        )}

        {error && <p style={{ color: 'red', marginTop: '8px' }}>{error}</p>}

        <textarea
          className={styles.textarea}
          value={newComment}
          onChange={(e) => setNewComment(e.target.value)}
          placeholder="Write your comment..."
          disabled={submitting}
        />

        <div className={styles.actions}>
          <button className={styles.closeBtn} onClick={onClose} disabled={submitting}>
            Cancel
          </button>
          <button className={styles.submitBtn} onClick={handleSubmit} disabled={submitting}>
            {submitting ? 'Submitting...' : 'Submit'}
          </button>
        </div>
      </div>
    </div>
  );
}
