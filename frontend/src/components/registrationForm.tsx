'use client';
import { useState } from "react";
import { RegistrationFormData } from '../lib/types'
import axios from 'axios'

const RegistrationForm = () => {
    // a use state for managing the registration form:
    const [formData, setFormData] = useState<RegistrationFormData>({
        email: '',
        password: '',
        firstName: '',
        lastName: '',
        dateOfBirth: '',
        avatar: null,
        nickname: '',
        aboutMe: '',
        privacyStatus: 'public',
    })

    // a use state to handle the error in relation to keys of the form data object as optional:
    const [errors, setErrors] = useState<Partial<Record<keyof RegistrationFormData, string>>>({})

    // making a signal when the an async operation is happening:
    const [loading, setLoading] = useState(false)

    //  A function to handle change in the input fields:
    const handleChange = (
        e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
    ) => {
        const { name, value, type } = e.target
        if (type === 'file') {
            setFormData((prev) => ({ ...prev, [name]: File }));
        } else {
            setFormData((prev) => ({ ...prev, [name]: value }))
        }
    };

    // ensure all required fields are filled out and valid:
    const validate = () => {
        const newErrors: typeof errors = {};
        if (!formData.email) newErrors.email = 'Email is required';
        else if (!/\S+@\S+\.\S+/.test(formData.email)) newErrors.email = 'Invalid email';
        if (!formData.password) newErrors.password = 'Password is required';
        if (!formData.firstName) newErrors.firstName = 'First name is required';
        if (!formData.lastName) newErrors.lastName = 'Last name is required';
        if (!formData.dateOfBirth) newErrors.dateOfBirth = 'Date of birth is required';
        return newErrors;
    }

    // Handle the form submition:
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        const validationErrors = validate();
        if (Object.keys(validationErrors).length > 0) {
            setErrors(validationErrors);
            return
        }

        setErrors({})
        setLoading(true)
        // encapsulate form data:
        // step1: upload avatar if exists:
        try {
            let avatarUrl: string | null = null
            if (formData.avatar) {
                const avataData = new FormData();
                avataData.append('avatar', formData.avatar)
                // ease the communication with back end using axios:
                const uploadsRes = await axios.post(`http://localhost:8080/api/upload-avatar`, avataData, {
                    headers: { 'Content-Type': 'multipart/form-data' },
                });

                avatarUrl = uploadsRes.data.url // assuming the backend returns {url: ...}
            }

            // step2 : send json registration data with avatar url or (null):
            const registrationData = {
                email: formData.email,
                password: formData.password,
                firstName: formData.firstName,
                lastName: formData.lastName,
                dateOfBirth: formData.dateOfBirth,
                nickname: formData.nickname,
                aboutMe: formData.aboutMe,
                privacyStatus: formData.privacyStatus,
                avatarUrl,
            };
            const res = await axios.post('http://localhost:8080/api/register', registrationData, {
                headers: { 'Content-Type': 'application/json' },
            })
            alert('Registration successful!');
            console.log(res.data);
        } catch (error: any) {
            console.error('Registration failed:', error);
            alert(
                error?.response?.data?.message ||
                'An error occurred during registration.'
            );
        } finally {
            setLoading(false);
        }
    };

    // return the componet to display on the browser:
    return (
        <form onSubmit={handleSubmit} className="max-w-md mx-auto p-6 border rounded-md shadow-md bg-white">
            {/* Email */}
            <input
                type="email"
                name="email"
                placeholder="Email"
                value={formData.email}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
            />
            {errors.email && <span className="text-red-600 text-sm mb-2 block">{errors.email}</span>}

            {/* Password */}
            <input
                type="password"
                name="password"
                placeholder="Password"
                value={formData.password}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
            />
            {errors.password && <span className="text-red-600 text-sm mb-2 block">{errors.password}</span>}

            {/* First Name */}
            <input
                type="text"
                name="firstName"
                placeholder="First Name"
                value={formData.firstName}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
            />
            {errors.firstName && <span className="text-red-600 text-sm mb-2 block">{errors.firstName}</span>}

            {/* Last Name */}
            <input
                type="text"
                name="lastName"
                placeholder="Last Name"
                value={formData.lastName}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
            />
            {errors.lastName && <span className="text-red-600 text-sm mb-2 block">{errors.lastName}</span>}

            {/* Date of Birth */}
            <input
                type="date"
                name="dateOfBirth"
                value={formData.dateOfBirth}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
            />
            {errors.dateOfBirth && <span className="text-red-600 text-sm mb-2 block">{errors.dateOfBirth}</span>}

            {/* Avatar */}
            <input
                type="file"
                name="avatar"
                accept="image/*"
                onChange={handleChange}
                className="w-full mb-4"
            />

            {/* Nickname */}
            <input
                type="text"
                name="nickname"
                placeholder="Nickname (optional)"
                value={formData.nickname}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            />

            {/* About Me */}
            <textarea
                name="aboutMe"
                placeholder="About Me (optional)"
                value={formData.aboutMe}
                onChange={handleChange}
                className="w-full p-2 mb-4 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none"
                rows={4}
            />

            {/* Privacy Status */}
            <label className="block mb-1 font-semibold text-gray-700">Profile Privacy:</label>
            <select
                name="privacyStatus"
                value={formData.privacyStatus}
                onChange={handleChange}
                className="w-full p-2 mb-4 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
                <option value="public">Public</option>
                <option value="private">Private</option>
            </select>

            {/* Submit Button */}
            <button
                type="submit"
                disabled={loading}
                className="w-full bg-blue-600 text-white p-2 rounded hover:bg-blue-700 disabled:opacity-50"
            >
                {loading ? 'Submitting...' : 'Register'}
            </button>
            <p className="mt-4">
                You have an account? <a href="/login" className="text-blue-600 underline">Login</a>
            </p>
        </form>

    )

}

export default RegistrationForm;