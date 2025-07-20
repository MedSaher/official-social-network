'use client'
import {useState} from 'react'
import axios from 'axios'

export default function PostForm() {
  const [formData, setFormData] = useState({
    title: '',
    content: '',
    privacy: 'public',
  image: null as File | null
 })

 const [message, setMessage] = useState('')

 // handle input change:
const handleChange = (
  e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
) => {
  const { name, value, type, files } = e.target as HTMLInputElement

  if (type === 'file') {
    setFormData((prev) => ({ ...prev, [name]: files?.[0] || null }))
  } else {
    setFormData((prev) => ({ ...prev, [name]: value }))
  }
}

// handle form submit:
const handleSubmit = async (e: React.FormEvent) =>{
e.preventDefault()
const data = new FormData()
    data.append('title', formData.title)
    data.append('content', formData.content)
    data.append('privacy', formData.privacy)
    if (formData.image) data.append('image', formData.image)
        try {
    const response = await axios.post('http://localhost:8080/api/post', data, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })

    if (response.status === 200 || response.status === 201) {
      setMessage('✅ Post created successfully!')
      setFormData({ title: '', content: '', privacy: 'public', image: null })
    } else {
      setMessage('❌ Failed to create post.')
    }
  } catch (err) {
    console.error('Error submitting form:', err)
    setMessage('❌ Error while submitting post.')
  }
}
  return (
    <div className="max-w-lg mx-auto p-6 bg-white rounded shadow">
      <h2 className="text-2xl mb-4">Create Post</h2>
      <form onSubmit={handleSubmit} className="space-y-4" encType="multipart/form-data">
        <input
          type="text"
          name="title"
          value={formData.title}
          onChange={handleChange}
          required
          placeholder="Title"
          className="w-full p-2 border rounded"
        />
        <textarea
          name="content"
          value={formData.content}
          onChange={handleChange}
          required
          placeholder="What's on your mind?"
          className="w-full p-2 border rounded h-32"
        />
        <select
          name="privacy"
          value={formData.privacy}
          onChange={handleChange}
          className="w-full p-2 border rounded"
        >
          <option value="public">Public</option>
          <option value="almost_private">Almost Private</option>
          <option value="private">Private</option>
        </select>
        <input
          type="file"
          name="image"
          accept="image/*"
          onChange={handleChange}
          className="w-full"
        />
        <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded">
          Submit
        </button>
      </form>
      {message && <p className="mt-4 text-center">{message}</p>}
    </div>
  )
}