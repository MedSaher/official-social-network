import ProfileHeader from './components/ProfileHeader'

interface Profile {
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

async function getProfile(id: string): Promise<Profile> {
  const res = await fetch(`http://localhost:8080/api/profile/${id}`, {
    cache: 'no-store',
  })

  if (!res.ok) {
    throw new Error('Failed to fetch profile')
  }

  return res.json()
}

export default async function ProfilePage({
  params,
}: {
  params: { id: string }
}) {
  const profile = await getProfile(params.id)

  return (
    <div className="p-4">
      <ProfileHeader {...profile} />
    </div>
  )
}
