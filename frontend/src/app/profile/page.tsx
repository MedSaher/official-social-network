"use client";

import { useEffect, useState } from "react";
import "./profile.css";
import CreateGroupForm from "@/components/groups/CreateGroupForm";
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

  const [showOptions, setShowOptions] = useState(false);
  const [showFollowersModal, setShowFollowersModal] = useState(false);
  const [showFollowingModal, setShowFollowingModal] = useState(false);
  const [showBigAvatar, setShowBigAvatar] = useState(false); // 🔥 NEW STATE

  useEffect(() => {
    async function fetchFullProfile() {
      try {
        const res = await fetch("http://localhost:8080/api/profile", {
          credentials: "include",
        });
        if (!res.ok) throw new Error(`HTTP error! Status: ${res.status}`);

        const data: FullProfileResponse = await res.json();
        console.log("Fetched profile data:", data);

        setProfile(data.user);
        setFollowers(data.followers);
        setFollowing(data.following);
        setFollowersCount(data.followers_count);
        setFollowingCount(data.following_count);
      } catch (error) {
        console.error("Failed to fetch profile:", error);
      }
    }
    fetchFullProfile();
  }, []);

  return (
    <div className="profile-container">
      {profile && (
        <div className="profile-card">
          <h2>{profile.username}</h2>
          <p>{profile.email}</p>
          <p>
            {profile.firstName} {profile.lastName}
          </p>
          <p>{profile.privacyStatus}</p>
          <p>Gender: {profile.gender}</p>
          <p className="joined">Joined: {profile.createdAt}</p>
          {!profile.avatarUrl ? (
            <div className="avatar-placeholder">No Avatar</div>
          ) : (
            <>
            <CreateGroupForm />
              <img
                src={`http://localhost:8080/${profile.avatarUrl}`}
                alt="User Avatar"
                className="avatar"
                onClick={() => setShowBigAvatar(true)} // 👆 Click to open modal
              />
              {showBigAvatar && (
                <div
                  className="avatar-modal"
                  onClick={() => setShowBigAvatar(false)}
                >
                  <span className="close-avatar">&times;</span>
                  <img
                    src={`http://localhost:8080/${profile.avatarUrl}`}
                    alt="Full Size Avatar"
                    className="avatar-large"
                  />
                </div>
              )}
            </>
          )}
        </div>
      )}

      <div className="buttons">
        <button onClick={() => setShowFollowersModal(true)}>
          Followers: {followersCount}
        </button>
        <button onClick={() => setShowFollowingModal(true)}>
          Following: {followingCount}
        </button>
      </div>

      {showFollowersModal && (
        <div className="modal">
          <h3>Followers</h3>
          {followers.map((f) => (
            <div key={f.id} className="follower">
              <p>{f.username}</p>
              {f.avatarUrl && <img src={f.avatarUrl} alt="follower avatar" />}
            </div>
          ))}
        </div>
      )}

      {showFollowingModal && (
        <div className="modal">
          <h3>Following</h3>
          {following.map((f) => (
            <div key={f.id} className="follower">
              <p>{f.username}</p>
              {f.avatarUrl && <img src={f.avatarUrl} alt="following avatar" />}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
