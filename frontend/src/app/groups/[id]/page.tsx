// app/groups/[id]/page.tsx
import React from 'react';
import styles from './css/GroupPage.module.css';
import GroupAdminWrapper from './GroupAdmin';

export default function GroupPage({ params }: { params: { id: string } }) {
  const { id } = params;
  
  return (
    <div className={styles.container}>
      <aside className={styles.sidebar}>
        <h1 className={styles.title}>Group #{id}</h1>
        <nav className={styles.nav}>
          <a href="#requests">Requests</a>
          <a href="#posts">Posts</a>
          <a href="#events">Events</a>
          <a href="#chat">Chat</a>
        </nav>
      </aside>

      <main className={styles.main}>

        <GroupAdminWrapper groupId={id} />

        <section id="posts" className={styles.card}>
          <h2>Group Posts</h2>
          <p className={styles.placeholder}>Posts will appear here.</p>
        </section>

        <section id="events" className={styles.card}>
          <h2>Group Events</h2>
          <p className={styles.placeholder}>Events will appear here.</p>
        </section>

        <section id="chat" className={styles.card}>
          <h2>Group Chat</h2>
          <p className={styles.placeholder}>Chat will appear here.</p>
        </section>
      </main>
    </div>
  );
}
