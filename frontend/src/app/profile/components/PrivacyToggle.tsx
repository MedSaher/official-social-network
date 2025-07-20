// /app/profile/[id]/components/TogglePrivacyButton.tsx
'use client'

import { useState } from 'react'

interface TogglePrivacyButtonProps {
  isPrivate: boolean
}

export default function TogglePrivacyButton({ isPrivate }: TogglePrivacyButtonProps) {
  const [privateProfile, setPrivateProfile] = useState(isPrivate)
  const [loading, setLoading] = useState(false)

  /**
   * Handles toggling the privacy status.
   * Sends a POST request to backend to toggle privacy.
   */
  const togglePrivacy = async () => {
    if (loading) return

    setLoading(true)

    try {
      const res = await fetch(`/api/profile/toggle-privacy`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
      })

      if (!res.ok) {
        throw new Error(`Failed to toggle privacy: ${res.statusText}`)
      }

      // Flip local state only if server confirms success
      setPrivateProfile(prev => !prev)
    } catch (error) {
      console.error(error)
      alert('Error toggling profile privacy.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <button
      onClick={togglePrivacy}
      disabled={loading}
      className={`px-4 py-1 rounded text-white ${
        privateProfile ? 'bg-red-600 hover:bg-red-700' : 'bg-green-600 hover:bg-green-700'
      }`}
    >
      {loading ? 'Updating...' : privateProfile ? 'Make Public' : 'Make Private'}
    </button>
  )
}
