// components/CommentModal.tsx
import React from 'react';
import styles from './css/commentModal.module.css';

interface CommentModalProps {
  postId: number;
  onClose: () => void;
}

export default function CommentModal({ postId, onClose }: CommentModalProps) {
  return (
    <div className={styles.overlay}>
      <div className={styles.modal}>
        <h2>Comment on Post #{postId}</h2>
        <textarea className={styles.textarea} placeholder="Write your comment..."></textarea>
        <div className={styles.actions}>
          <button className={styles.closeBtn} onClick={onClose}>Cancel</button>
          <button className={styles.submitBtn}>Submit</button>
        </div>
      </div>
    </div>
  );
}
