'use client';

import { useAuth } from './AuthContext';
import { useRouter } from 'next/navigation';
import CreatePost from './postCreate';
import PostList from "./postFetch";
import GroupList from '@/components/groups/GroupList';
import AllUsers from './user/fetchuser';
import { useState } from 'react';
import styles from './css/HomePage.module.css';

export default function HomePage() {
  const router = useRouter();
  const { logout, user } = useAuth();
  const [activeTab, setActiveTab] = useState('feed');
  const [showCreatePostModal, setShowCreatePostModal] = useState(false);
  const [refreshKey, setRefreshKey] = useState(0);

  const handleLogout = async () => {
    try {
      await logout();
      router.push('/login');
    } catch (error) {
      console.error('Failed to logout:', error);
    }
  };


  return (
    <div className={styles.container}>
      {/* Main Content */}
      <div className={styles.mainContent}>
        {/* Left Sidebar */}
        <aside className={styles.leftSidebar}>
          <nav className={styles.navigation}>
            <button 
              className={`${styles.navItem} ${activeTab === 'feed' ? styles.active : ''}`}
              onClick={() => setActiveTab('feed')}
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z"/>
              </svg>
              <span>Home Feed</span>
            </button>
            <button 
              className={`${styles.navItem} ${activeTab === 'groups' ? styles.active : ''}`}
              onClick={() => setActiveTab('groups')}
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                <path d="M16 11c1.66 0 2.99-1.34 2.99-3S17.66 5 16 5c-1.66 0-3 1.34-3 3s1.34 3 3 3zm-8 0c1.66 0 2.99-1.34 2.99-3S9.66 5 8 5C6.34 5 5 6.34 5 8s1.34 3 3 3zm0 2c-2.33 0-7 1.17-7 3.5V19h14v-2.5c0-2.33-4.67-3.5-7-3.5zm8 0c-.29 0-.62.02-.97.05 1.16.84 1.97 1.97 1.97 3.45V19h6v-2.5c0-2.33-4.67-3.5-7-3.5z"/>
              </svg>
              <span>Groups</span>
            </button>
            <button 
              className={`${styles.navItem} ${activeTab === 'chat' ? styles.active : ''}`}
              onClick={() => setActiveTab('chat')}
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                <path d="M20 2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zM6 9h12v2H6V9zm8 5H6v-2h8v2zm4-6H6V6h12v2z"/>
              </svg>
              <span>Messages</span>
            </button>
            <button 
              className={`${styles.navItem} ${activeTab === 'events' ? styles.active : ''}`}
              onClick={() => setActiveTab('events')}
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                <path d="M17 12h-5v5h5v-5zM16 1v2H8V1H6v2H5c-1.11 0-1.99.9-1.99 2L3 19c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2h-1V1h-2zm3 18H5V8h14v11z"/>
              </svg>
              <span>Events</span>
            </button>
          </nav>

          <div className={styles.createGroupSection}>
            <h3>Create Something New</h3>
            <button 
              className={styles.createButton}
              onClick={() => setShowCreatePostModal(true)}
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
                <path d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/>
              </svg>
              New Post
            </button>
          </div>
        </aside>

        {/* Main Feed */}
        <main className={styles.feed}>
          {activeTab === 'feed' && (
            <>
              <CreatePost/>
              <PostList key={refreshKey} />
            </>
          )}
          {activeTab === 'groups' && <GroupList />}
          {activeTab === 'chat' && <AllUsers />}
          {activeTab === 'events' && (
            <div className={styles.comingSoon}>
              <h2>Events Coming Soon!</h2>
              <p>We're working on bringing you the best event experience.</p>
            </div>
          )}
        </main>

        {/* Right Sidebar */}
        <aside className={styles.rightSidebar}>
          <div className={styles.upcomingEvents}>
            <h3>Upcoming Events</h3>
            <div className={styles.eventCard}>
              <div className={styles.eventDate}>
                <span className={styles.eventDay}>15</span>
                <span className={styles.eventMonth}>Jun</span>
              </div>
              <div className={styles.eventDetails}>
                <h4>Team Meetup</h4>
                <p>Virtual Hangout</p>
              </div>
            </div>
          </div>

          <div className={styles.suggestedGroups}>
            <h3>Suggested Groups</h3>
            <div className={styles.groupCard}>
              <div className={styles.groupIcon}>📚</div>
              <div className={styles.groupDetails}>
                <h4>Book Lovers</h4>
                <p>125 members</p>
              </div>
            </div>
          </div>
        </aside>
      </div>

      {/* Create Post Modal */}
      {showCreatePostModal && (
        <div className={styles.modalOverlay} onClick={() => setShowCreatePostModal(false)}>
          <div className={styles.modalContent} onClick={e => e.stopPropagation()}>
            <button 
              className={styles.closeModal}
              onClick={() => setShowCreatePostModal(false)}
            >
              &times;
            </button>
            <CreatePost/>
          </div>
        </div>
      )}
    </div>
  );
}