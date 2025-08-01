'use client';

import React, { useEffect, useState, use } from 'react';
import styles from './css/GroupPage.module.css';
import GroupAdminWrapper from './GroupAdmin';
import CreatePost from '@/components/postCreate';
import CreateEventModal from './CreatEvent';

interface Post {
  id: number;
  content: string;
  created_at: string;
  user_id: number;
  image_path: string | null;
  user_name: string;
  user_avatar: string | null;
}

interface Event {
  id: number;
  title: string;
  description: string;
  event_date: string;
  created_at: string;
  creator_name: string;
}

interface GroupInfo {
  id: number;
  title: string;
  description: string;
  creator_id: number;
  posts_count: number;
  events_count: number;
}

export default function GroupPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const [activeTab, setActiveTab] = useState('posts');
  const [posts, setPosts] = useState<Post[]>([]);
  const [events, setEvents] = useState<Event[]>([]);
  const [groupInfo, setGroupInfo] = useState<GroupInfo | null>(null);
  const [loadingPosts, setLoadingPosts] = useState(true);
  const [loadingEvents, setLoadingEvents] = useState(true);
  const [loadingGroupInfo, setLoadingGroupInfo] = useState(true);
  const [error, setError] = useState<string | null>(null);

   // Move these hooks to the top level
  const [showCreateEventModal, setShowCreateEventModal] = useState(false);
  const [showCreatePostModal, setShowCreatePostModal] = useState(false);

  useEffect(() => {
    async function fetchGroupInfo() {
      try {
        const res = await fetch(`http://localhost:8080/api/groups/${id}/group_info`, {
          credentials: 'include'
        });
        if (!res.ok) throw new Error('Failed to fetch group info');
        const data = await res.json();
        setGroupInfo(data);
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message);
        }
      } finally {
        setLoadingGroupInfo(false);
      }
    }

    async function fetchPosts() {
      try {
        const res = await fetch(`http://localhost:8080/api/groups/${id}/posts`, { credentials: 'include' });
        if (!res.ok) throw new Error('Failed to fetch posts');
        const data = await res.json();
        setPosts(data);
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message);
        }
      } finally {
        setLoadingPosts(false);
      }
    }

    async function fetchEvents() {
      try {
        const res = await fetch(`http://localhost:8080/api/groups/${id}/events`, { credentials: 'include' });
        if (!res.ok) throw new Error('Failed to fetch events');
        const data = await res.json();
        setEvents(data);
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message);
        }
      } finally {
        setLoadingEvents(false);
      }
    }

    fetchGroupInfo();
    fetchPosts();
    fetchEvents();
  }, [id]);

  const renderTabContent = () => {
    switch (activeTab) {
      case 'posts':
        return (
          <div className={styles.tabContent}>
            <CreatePost groupId={id} />
            {loadingPosts ? (
              <div className={styles.loading}>
                <div className={styles.spinner}></div>
                <p>Loading posts...</p>
              </div>
            ) : error ? (
              <p className={styles.error}>{error}</p>
            ) : posts === null ? (
              <p className={styles.placeholder}>No posts yet. Be the first to share something!</p>
            ) : (
              <div className={styles.postsGrid}>
                {posts.map(post => (
                  <div key={post.id} className={styles.postCard}>
                    <div className={styles.postHeader}>
                      {post.user_avatar ? (
                        <img src={`http://localhost:8080/${post.user_avatar}`} alt="User" className={styles.userAvatar} />
                      ) : (
                        <div className={styles.avatarPlaceholder}>{post.user_name.charAt(0)}</div>
                      )}
                      <div>
                        <p className={styles.userName}>{post.user_name}</p>
                        <p className={styles.postTime}>{new Date(post.created_at).toLocaleString()}</p>
                      </div>
                    </div>
                    <p className={styles.postContent}>{post.content}</p>
                    {post.image_path && (
                      <div className={styles.postImageContainer}>
                        <img
                          src={`http://localhost:8080/${post.image_path}`}
                          alt="Post"
                          className={styles.postImage}
                        />
                      </div>
                    )}
                  </div>
                ))}
              </div>
            )}
          </div>
        );
       case 'events':
        return (
          <div className={styles.tabContent}>
            <button
              className={styles.createEventBtn}
              onClick={() => setShowCreateEventModal(true)}
            >
              + Create Event
            </button>

            {showCreateEventModal && (
              <CreateEventModal
                groupId={id}
                onClose={() => setShowCreateEventModal(false)}
              />
            )}

            {loadingEvents ? (
              <div className={styles.loading}>
                <div className={styles.spinner}></div>
                <p>Loading events...</p>
              </div>
            ) : error ? (
              <p className={styles.error}>{error}</p>
            ) : events === null ? (
              <p className={styles.placeholder}>No upcoming events. Create one to get started!</p>
            ) : (
              <div className={styles.eventsGrid}>
                {events.map(event => (
                  <div key={event.id} className={styles.eventCard}>
                    <div className={styles.eventDate}>
                      <span className={styles.eventDay}>
                        {new Date(event.event_date).getDate()}
                      </span>
                      <span className={styles.eventMonth}>
                        {new Date(event.event_date).toLocaleString('default', { month: 'short' })}
                      </span>
                    </div>
                    <div className={styles.eventDetails}>
                      <h4 className={styles.eventTitle}>{event.title}</h4>
                      <p className={styles.eventCreator}>Created by {event.creator_name}</p>
                      <p className={styles.eventDescription}>{event.description}</p>
                      <div className={styles.eventActions}>
                        <button className={styles.goingBtn}>Going</button>
                        <button className={styles.notGoingBtn}>Not Going</button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        );
      case 'members':
        return (
          <div className={styles.tabContent}>
            <p className={styles.placeholder}>Member list will appear here.</p>
          </div>
        );
      case 'chat':
        return (
          <div className={styles.tabContent}>
            <div className={styles.chatContainer}>
              <div className={styles.chatMessages}>
                <p className={styles.placeholder}>Group chat will appear here.</p>
              </div>
              <div className={styles.chatInput}>
                <input type="text" placeholder="Type a message..." />
                <button className={styles.sendBtn}>Send</button>
              </div>
            </div>
          </div>
        );
      default:
        return null;
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div className={styles.headerContent}>
          {loadingGroupInfo ? (
            <div className={styles.loading}>
              <div className={styles.spinner}></div>
              <p>Loading group info...</p>
            </div>
          ) : error ? (
            <p className={styles.error}>{error}</p>
          ) : groupInfo ? (
            <>
              <h1 className={styles.groupTitle}>{groupInfo.title}</h1>
              <p className={styles.groupDescription}>{groupInfo.description}</p>
              <div className={styles.groupStats}>
                <span>{groupInfo.posts_count} Posts</span>
                <span>{groupInfo.events_count} Events</span>
              </div>
            </>
          ) : (
            <p className={styles.error}>Failed to load group information</p>
          )}
        </div>
        <div className={styles.headerActions}>
          <button className={styles.inviteBtn}>Invite People</button>
          <button className={styles.settingsBtn}>Group Settings</button>
        </div>
      </div>

      <div className={styles.mainLayout}>
        <div className={styles.sidebar}>
          <nav className={styles.nav}>
            <button
              className={`${styles.navItem} ${activeTab === 'posts' ? styles.active : ''}`}
              onClick={() => setActiveTab('posts')}
            >
              <span className={styles.navIcon}>📝</span>
              Posts
            </button>
            <button
              className={`${styles.navItem} ${activeTab === 'events' ? styles.active : ''}`}
              onClick={() => setActiveTab('events')}
            >
              <span className={styles.navIcon}>📅</span>
              Events
            </button>
            <button
              className={`${styles.navItem} ${activeTab === 'members' ? styles.active : ''}`}
              onClick={() => setActiveTab('members')}
            >
              <span className={styles.navIcon}>👥</span>
              Members
            </button>
            <button
              className={`${styles.navItem} ${activeTab === 'chat' ? styles.active : ''}`}
              onClick={() => setActiveTab('chat')}
            >
              <span className={styles.navIcon}>💬</span>
              Chat
            </button>
          </nav>
          <GroupAdminWrapper groupId={id} />
        </div>

        <div className={styles.content}>
          {renderTabContent()}
        </div>
      </div>
    </div>
  );
}