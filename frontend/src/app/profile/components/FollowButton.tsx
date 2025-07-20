// /app/profile/[id]/components/FollowButton.tsx
'use client'

import { useState } from 'react'

interface FollowButtonProps {
  userId: string
  isFollowing: boolean
}

/**
 * Follow/Unfollow button component.
 * Handles sending follow/unfollow requests to backend.
 */
export default function FollowButton({ userId, isFollowing }: FollowButtonProps) {
  const [following, setFollowing] = useState(isFollowing)
  const [loading, setLoading] = useState(false)

  const handleClick = async () => {
    if (loading) return

    setLoading(true)

    try {
      const method = following ? 'DELETE' : 'POST'
      const url = `/api/profile/${userId}/${following ? 'unfollow' : 'follow'}`

      const res = await fetch(url, { method })

      if (!res.ok) {
        throw new Error(`Failed to ${following ? 'unfollow' : 'follow'} user: ${res.statusText}`)
      }

      setFollowing(prev => !prev)
    } catch (error) {
      console.error(error)
      alert('Error processing follow/unfollow action.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <button
      onClick={handleClick}
      disabled={loading}
      className={`px-4 py-1 rounded text-white ${
        following ? 'bg-gray-600 hover:bg-gray-700' : 'bg-blue-600 hover:bg-blue-700'
      }`}
    >
      {loading ? 'Processing...' : following ? 'Unfollow' : 'Follow'}
    </button>
  )
}
