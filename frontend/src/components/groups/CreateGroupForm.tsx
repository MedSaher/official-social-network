'use client';

import React, { useState } from 'react';
import styles from './CreateGroupForm.module.css';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus, faTimes } from '@fortawesome/free-solid-svg-icons'; // You can use more if needed

export default function CreateGroupForm() {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const [showForm, setShowForm] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();

    if (!title.trim() || !description.trim()) {
      setError('Both title and description are required.');
      return;
    }

    setError('');
    setMessage('Submitting...');

    try {
      const res = await fetch('http://localhost:8080/api/groups/create_group', {
        method: 'POST',
        credentials: 'include',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, description }),
      });

      if (!res.ok) throw new Error('Group creation failed.');

      setMessage('Group created successfully!');
      setTitle('');
      setDescription('');
    } catch (err) {
      setError('Error: could not create group.');
      setMessage('');
    }
  }

  const toggleFormVisibility = () => {
    setShowForm(!showForm);
  };

  return (
    <div>
      <button onClick={toggleFormVisibility} className={styles.createGroupButton}>
        <FontAwesomeIcon icon={faPlus} className="text-white text-lg" />
        {/* You can add text next to the icon if needed */}
      </button>

      {showForm && (
        <div className={styles.formPopupOverlay}>
          <div className={styles.formPopup}>
            <h2>Create a New Group</h2>
            <form onSubmit={handleSubmit} className={styles.form}>
              <input
                type="text"
                placeholder="Group Title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                className={styles.input}
              />
              <textarea
                placeholder="Group Description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                className={styles.textarea}
              />
              <button type="submit" className={styles.submitBtn}>
                Create Group
              </button>

              {message && <p className={styles.success}>{message}</p>}
              {error && <p className={styles.error}>{error}</p>}
            </form>

            <button
              onClick={toggleFormVisibility}
              className={styles.closeButton}
            >
              <FontAwesomeIcon icon={faTimes} /> Close
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
