'use client';

import { useState } from 'react';
import styles from './css/CreatePost.module.css'; // Adjust if needed

interface PostCreateProps {
  groupId?: string;
}

export default function CreatePost({ groupId }: PostCreateProps) {
  const [content, setContent] = useState('');
  const [privacy, setPrivacy] = useState('public');
  const [image, setImage] = useState<File | null>(null);
  const [responseMsg, setResponseMsg] = useState('');
  const [responseColor, setResponseColor] = useState('green');

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();

    if (!content.trim()) {
      setResponseColor('red');
      setResponseMsg('Content is required');
      return;
    }

    const formData = new FormData();
    formData.append('content', content);
    formData.append('privacy', privacy);
    if (groupId) formData.append('group_id', groupId);
    if (image) formData.append('image', image);

    try {
      const res = await fetch('http://localhost:8080/api/posts/create_post', {
        method: 'POST',
        credentials: 'include',
        body: formData,
      });

      const data = await res.json();

      if (res.ok) {
        console.log('Post created successfully:', data);
        
        setResponseColor('green');
        setResponseMsg('✅ Post created successfully!');
        setContent('');
        setPrivacy('public');
        setImage(null);
      } else {
        setResponseColor('red');
        setResponseMsg(`❌ Error: ${data.error || 'Unknown error'}`);
      }
    } catch (err: unknown) {
      if (err instanceof Error) {
        setResponseColor('red');
        setResponseMsg(`❌ Network error: ${err.message || 'Unknown error'}`);
      }
    }
  }

  return (
    <div className={styles.formContainer}>
      <h2 className={styles.title}>Create a New Post</h2>

      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <label htmlFor="content" className={styles.label}>Post Content</label>
        <textarea
          id="content"
          name="content"
          rows={4}
          required
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className={styles.textarea}
        />

        <label htmlFor="privacy" className={styles.label}>Privacy</label>
        <select
          id="privacy"
          name="privacy"
          value={privacy}
          onChange={(e) => setPrivacy(e.target.value)}
          className={styles.select}
        >
          <option value="public">Public</option>
          <option value="almost_private">Almost Private</option>
          <option value="private">Private</option>
        </select>

        <label htmlFor="image" className={styles.label}>Upload Image</label>
        <input
          type="file"
          id="image"
          name="image"
          accept=".jpg, .jpeg, .png, .gif"
          onChange={(e) =>
            setImage(e.target.files?.[0] || null)
          }
          className={styles.fileInput}
        />

        <button type="submit" className={styles.button}>Post</button>

        <div className={styles.response} style={{ color: responseColor }}>
          {responseMsg}
        </div>
      </form>
    </div>
  );
}
