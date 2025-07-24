"use client";

import { ProfileType, PostType, UserListType } from "@/lib/types";
import { useEffect, useState } from "react";

type FullUserData = {
  profile: ProfileType;
  posts: PostType[];
  followers?: UserListType[];   // optional if you want
  following?: UserListType[];   // optional if you want
};

export default function ProfilePage() {
  const [profile, setProfile] = useState<ProfileType | null>(null);
  const [profileData, setData] = useState<PostType[]>([]);
  const [userList, setUserList] = useState<UserListType[]>([]);
  const [showOptions, setShowOptions] = useState(false);
  const [showFollowersModal, setShowFollowersModal] = useState(false);
  const [showFollowingModal, setShowFollowingModal] = useState(false);

  // Fetch all user data in one request
  useEffect(() => {
    async function fetchFullUserData() {
      try {
        const res = await fetch('/api/getFullUserData', { credentials: 'include' });
        if (!res.ok) {
          throw new Error(`HTTP error! Status: ${res.status}`);
        }
        const data: FullUserData = await res.json();
        setProfile(data.profile);
        setData(data.posts);
        // If your API returns followers/following by default, you can set it here:
        // setUserList(data.followers || []);
      } catch (error) {
        console.error("Failed to fetch full user data:", error);
      }
    }
    fetchFullUserData();
  }, []);

  const fetchUserList = async (listType: "followers" | "following") => {
    try {
      const response = await fetch(`/api/followersList?&type=${listType}`, {
        credentials: 'include',
      });
      if (!response.ok) {
        console.error("Error:", await response.text());
        return;
      }
      const data: UserListType[] = await response.json();
      setUserList(data);
      if (listType === "followers") {
        setShowFollowersModal(true);
      } else {
        setShowFollowingModal(true);
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handleSelection = async (choise: "public" | "private") => {
    setShowOptions(false);
    try {
      const res = await fetch(`/api/updatePrivacy`, {
        method: 'POST',
        credentials: 'include',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ accountType: choise }),
      });
      if (!res.ok) {
        throw new Error(`http error here : ${res.status}`);
      }
      setProfile(prev => prev ? { ...prev, accountType: choise } : prev);
    } catch (error) {
      console.error("Failed to update privacy:", error);
    }
  };

  return (
    <div>
      {/* Render your profile, posts, modals, options here */}
    </div>
  );
}
