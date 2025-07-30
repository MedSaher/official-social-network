'use client';

import { useEffect, useState } from 'react';
import styles from './ProfilePage.module.css';
import CreatePost from '@/components/postCreate';
import PostList from '@/components/postFetch';

type UserProfileDTO = {
  id: number;
  username: string;
  firstName: string;
  lastName: string;
  avatarUrl?: string | null;
  email: string;
  aboutMe?: string | null;
  privacyStatus: string;
  gender: string;
  createdAt: string;
};

type FollowerInfo = {
  id: number;
  username: string;
  avatarUrl?: string | null;
};

type FullProfileResponse = {
  user: UserProfileDTO;
  followers_count: number;
  following_count: number;
  followers: FollowerInfo[];
  following: FollowerInfo[];
};

export default function ProfilePage() {
  const [profile, setProfile] = useState<UserProfileDTO | null>(null);
  const [followers, setFollowers] = useState<FollowerInfo[]>([]);
  const [following, setFollowing] = useState<FollowerInfo[]>([]);
  const [followersCount, setFollowersCount] = useState(0);
  const [followingCount, setFollowingCount] = useState(0);
  const [activeTab, setActiveTab] = useState('posts');
  const [showFollowersModal, setShowFollowersModal] = useState(false);
  const [showFollowingModal, setShowFollowingModal] = useState(false);
  const [showBigAvatar, setShowBigAvatar] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function fetchFullProfile() {
      try {
        const res = await fetch("http://localhost:8080/api/profile", {
          credentials: "include",
        });
        if (!res.ok) throw new Error(`HTTP error! Status: ${res.status}`);

        const data: FullProfileResponse = await res.json();
        setProfile(data.user);
        setFollowers(data.followers);
        setFollowing(data.following);
        setFollowersCount(data.followers_count);
        setFollowingCount(data.following_count);
      } catch (error) {
        console.error("Failed to fetch profile:", error);
        setError(error instanceof Error ? error.message : 'Failed to load profile');
      } finally {
        setLoading(false);
      }
    }
    fetchFullProfile();
  }, []);

  const renderTabContent = () => {
    switch (activeTab) {
      case 'posts':
        return <PostList />;
      case 'about':
        return (
          <div className={styles.aboutSection}>
            <h3>About Me</h3>
            <p>{profile?.aboutMe || 'No information provided'}</p>
            
            <div className={styles.detailsGrid}>
              <div className={styles.detailItem}>
                <span className={styles.detailLabel}>Email:</span>
                <span>{profile?.email}</span>
              </div>
              <div className={styles.detailItem}>
                <span className={styles.detailLabel}>Gender:</span>
                <span>{profile?.gender}</span>
              </div>
              <div className={styles.detailItem}>
                <span className={styles.detailLabel}>Privacy:</span>
                <span>{profile?.privacyStatus}</span>
              </div>
              <div className={styles.detailItem}>
                <span className={styles.detailLabel}>Member Since:</span>
                <span>{profile ? new Date(profile.createdAt).toLocaleDateString() : ''}</span>
              </div>
            </div>
          </div>
        );
      case 'friends':
        return (
          <div className={styles.friendsSection}>
            <div className={styles.friendsStats}>
              <button 
                className={styles.friendsStatItem}
                onClick={() => setShowFollowersModal(true)}
              >
                <span className={styles.statNumber}>{followersCount}</span>
                <span>Followers</span>
              </button>
              <button 
                className={styles.friendsStatItem}
                onClick={() => setShowFollowingModal(true)}
              >
                <span className={styles.statNumber}>{followingCount}</span>
                <span>Following</span>
              </button>
            </div>
          </div>
        );
      default:
        return null;
    }
  };

  if (loading) {
    return (
      <div className={styles.loadingContainer}>
        <div className={styles.spinner}></div>
        <p>Loading profile...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className={styles.errorContainer}>
        <p className={styles.errorMessage}>{error}</p>
      </div>
    );
  }

  if (!profile) {
    return (
      <div className={styles.errorContainer}>
        <p className={styles.errorMessage}>Profile not found</p>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      {/* Header Section */}
      <div className={styles.header}>
        <div className={styles.coverPhoto}></div>
        
        <div className={styles.profileHeaderContent}>
          <div className={styles.avatarContainer}>
            {profile.avatarUrl ? (
              <>
                <img
                  src={`http://localhost:8080/${profile.avatarUrl}`}
                  alt="User Avatar"
                  className={styles.avatar}
                  onClick={() => setShowBigAvatar(true)}
                />
                {showBigAvatar && (
                  <div className={styles.avatarModal} onClick={() => setShowBigAvatar(false)}>
                    <span className={styles.closeAvatar}>&times;</span>
                    <img
                      src={`http://localhost:8080/${profile.avatarUrl}`}
                      alt="Full Size Avatar"
                      className={styles.avatarLarge}
                    />
                  </div>
                )}
              </>
            ) : (
              <div className={styles.avatarPlaceholder}>
                {profile.firstName.charAt(0)}{profile.lastName.charAt(0)}
              </div>
            )}
          </div>

          <div className={styles.profileInfo}>
            <h1 className={styles.username}>{profile.username}</h1>
            <p className={styles.name}>{profile.firstName} {profile.lastName}</p>
            
            <div className={styles.profileActions}>
              <button className={styles.editProfileBtn}>Edit Profile</button>
              <button className={styles.privacyBtn}>
                {profile.privacyStatus === 'public' ? 'Public' : 'Private'}
              </button>
            </div>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className={styles.mainContent}>
        {/* Sidebar Navigation */}
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
              className={`${styles.navItem} ${activeTab === 'about' ? styles.active : ''}`}
              onClick={() => setActiveTab('about')}
            >
              <span className={styles.navIcon}>ℹ️</span>
              About
            </button>
            <button
              className={`${styles.navItem} ${activeTab === 'friends' ? styles.active : ''}`}
              onClick={() => setActiveTab('friends')}
            >
              <span className={styles.navIcon}>👥</span>
              Friends
            </button>
          </nav>
        </div>

        {/* Content Area */}
        <div className={styles.content}>
          {activeTab === 'posts' && <CreatePost />}
          {renderTabContent()}
        </div>
      </div>

      {/* Modals */}
      {showFollowersModal && (
        <div className={styles.modalOverlay} onClick={() => setShowFollowersModal(false)}>
          <div className={styles.modalContent} onClick={e => e.stopPropagation()}>
            <h3>Followers ({followersCount})</h3>
            <div className={styles.followersList}>
              {followers.map((follower) => (
                <div key={follower.id} className={styles.followerItem}>
                  {follower.avatarUrl ? (
                    <img 
                      src={`http://localhost:8080/${follower.avatarUrl}`} 
                      alt={follower.username} 
                      className={styles.followerAvatar}
                    />
                  ) : (
                    <div className={styles.followerAvatarPlaceholder}>
                      {follower.username.charAt(0)}
                    </div>
                  )}
                  <span className={styles.followerUsername}>{follower.username}</span>
                </div>
              ))}
            </div>
            <button 
              className={styles.closeModalBtn}
              onClick={() => setShowFollowersModal(false)}
            >
              Close
            </button>
          </div>
        </div>
      )}

      {showFollowingModal && (
        <div className={styles.modalOverlay} onClick={() => setShowFollowingModal(false)}>
          <div className={styles.modalContent} onClick={e => e.stopPropagation()}>
            <h3>Following ({followingCount})</h3>
            <div className={styles.followersList}>
              {following.map((user) => (
                <div key={user.id} className={styles.followerItem}>
                  {user.avatarUrl ? (
                    <img 
                      src={`http://localhost:8080/${user.avatarUrl}`} 
                      alt={user.username} 
                      className={styles.followerAvatar}
                    />
                  ) : (
                    <div className={styles.followerAvatarPlaceholder}>
                      {user.username.charAt(0)}
                    </div>
                  )}
                  <span className={styles.followerUsername}>{user.username}</span>
                </div>
              ))}
            </div>
            <button 
              className={styles.closeModalBtn}
              onClick={() => setShowFollowingModal(false)}
            >
              Close
            </button>
          </div>
        </div>
      )}
    </div>
  );
}