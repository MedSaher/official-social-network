// ProfileHeader.tsx
'use client'

import TogglePrivacyButton from './PrivacyToggle'
import FollowButton from './FollowButton'

interface ProfileHeaderProps {
  id: string
  firstName: string
  lastName: string
  nickname?: string
  avatar?: string
  aboutMe?: string
  isPrivate: boolean
  isOwnProfile: boolean
  isFollowing: boolean
}

export default function ProfileHeader({
  id,
  firstName,
  lastName,
  nickname,
  avatar,
  aboutMe,
  isPrivate,
  isOwnProfile,
  isFollowing,
}: ProfileHeaderProps) {
  return (
    <div className="flex flex-col md:flex-row items-center md:items-start gap-6 bg-white shadow-md rounded-lg p-6 max-w-4xl mx-auto">
      {/* Avatar */}
      <div className="flex-shrink-0">
        <img
          src={avatar || '/default-avatar.png'}
          alt={`${firstName} ${lastName} avatar`}
          className="w-32 h-32 rounded-full object-cover border-4 border-indigo-500"
          loading="lazy"
        />
      </div>

      {/* User Info and Actions */}
      <div className="flex flex-col flex-grow">
        <div className="flex flex-col md:flex-row md:items-center md:justify-between">
          {/* Name and nickname */}
          <div>
            <h1 className="text-3xl font-extrabold text-gray-900">
              {firstName} {lastName}
            </h1>
            {nickname && (
              <p className="text-indigo-600 text-lg font-semibold mt-1">@{nickname}</p>
            )}
          </div>

          {/* Action buttons */}
          <div className="mt-4 md:mt-0 flex gap-3">
            {isOwnProfile ? (
              <TogglePrivacyButton isPrivate={isPrivate} />
            ) : (
              <FollowButton userId={id} isFollowing={isFollowing} />
            )}
          </div>
        </div>

        {/* About me */}
        {aboutMe && (
          <p className="mt-6 text-gray-700 text-base max-w-xl leading-relaxed">
            {aboutMe}
          </p>
        )}
      </div>
    </div>
  )
}
