'use client';
import { useState } from "react";
import { RegistrationFormData } from '../lib/types';
import axios from 'axios';

const RegistrationForm = () => {
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
        gender: '', // ← Added gender field
    });

    const [errors, setErrors] = useState<Partial<Record<keyof RegistrationFormData, string>>>({});
    const [loading, setLoading] = useState(false);

    const handleChange = (
        e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
    ) => {
        const { name, value, type, files } = e.target as HTMLInputElement;
        if (type === 'file' && files) {
            setFormData((prev) => ({ ...prev, [name]: files[0] }));
        } else {
            setFormData((prev) => ({ ...prev, [name]: value }));
        }
    };

    const validate = () => {
        const newErrors: typeof errors = {};
        if (!formData.email) newErrors.email = 'Email is required';
        else if (!/\S+@\S+\.\S+/.test(formData.email)) newErrors.email = 'Invalid email';
        if (!formData.password) newErrors.password = 'Password is required';
        if (!formData.firstName) newErrors.firstName = 'First name is required';
        if (!formData.lastName) newErrors.lastName = 'Last name is required';
        if (!formData.dateOfBirth) newErrors.dateOfBirth = 'Date of birth is required';
        if (!formData.gender) newErrors.gender = 'Gender is required';
        return newErrors;
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        const validationErrors = validate();
        if (Object.keys(validationErrors).length > 0) {
            setErrors(validationErrors);
            return;
        }

        setErrors({});
        setLoading(true);

        try {
            let avatarUrl: string | null = null;
            if (formData.avatar) {
                const avatarData = new FormData();
                avatarData.append('avatar', formData.avatar);
                const uploadRes = await axios.post(`http://localhost:8080/api/upload-avatar`, avatarData, {
                    headers: { 'Content-Type': 'multipart/form-data' },
                });
                avatarUrl = uploadRes.data.url;
            }

            const registrationData = {
                email: formData.email,
                password: formData.password,
                firstName: formData.firstName,
                lastName: formData.lastName,
                dateOfBirth: formData.dateOfBirth,
                aboutMe: formData.aboutMe,
                privacyStatus: formData.privacyStatus,
                gender: formData.gender,
                avatarUrl,
            };

            const res = await axios.post('http://localhost:8080/api/register', registrationData, {
                headers: { 'Content-Type': 'application/json' },
            });

            alert('Registration successful!');
            console.log(res.data);
        } catch (error: any) {
            console.error('Registration failed:', error);
            alert(error?.response?.data?.message || 'An error occurred during registration.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="max-w-md mx-auto p-6 border rounded-md shadow-md bg-white">
            <h2 className="text-2xl mb-4">Register</h2>

            {/* Email */}
            <input
                type="email"
                name="email"
                placeholder="Email"
                value={formData.email}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded"
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
                className="w-full p-2 mb-2 border border-gray-300 rounded"
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
                className="w-full p-2 mb-2 border border-gray-300 rounded"
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
                className="w-full p-2 mb-2 border border-gray-300 rounded"
                required
            />
            {errors.lastName && <span className="text-red-600 text-sm mb-2 block">{errors.lastName}</span>}

            {/* Date of Birth */}
            <input
                type="date"
                name="dateOfBirth"
                value={formData.dateOfBirth}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded"
                required
            />
            {errors.dateOfBirth && <span className="text-red-600 text-sm mb-2 block">{errors.dateOfBirth}</span>}

            {/* Gender */}
            <label className="block mb-1 font-semibold text-gray-700">Gender:</label>
            <select
                name="gender"
                value={formData.gender}
                onChange={handleChange}
                className="w-full p-2 mb-4 border border-gray-300 rounded"
                required
            >
                <option value="">Select Gender</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
            </select>
            {errors.gender && <span className="text-red-600 text-sm mb-2 block">{errors.gender}</span>}

            {/* Avatar */}
            <input
                type="file"
                name="avatar"
                accept="image/*"
                onChange={handleChange}
                className="w-full mb-4"
            />

            {/* Nickname (optional) */}
            <input
                type="text"
                name="nickname"
                placeholder="Nickname (optional)"
                value={formData.nickname}
                onChange={handleChange}
                className="w-full p-2 mb-2 border border-gray-300 rounded"
            />

            {/* About Me */}
            <textarea
                name="aboutMe"
                placeholder="About Me (optional)"
                value={formData.aboutMe}
                onChange={handleChange}
                className="w-full p-2 mb-4 border border-gray-300 rounded resize-none"
                rows={4}
            />

            {/* Privacy Status */}
            <label className="block mb-1 font-semibold text-gray-700">Profile Privacy:</label>
            <select
                name="privacyStatus"
                value={formData.privacyStatus}
                onChange={handleChange}
                className="w-full p-2 mb-4 border border-gray-300 rounded"
            >
                <option value="public">Public</option>
                <option value="private">Private</option>
                <option value="almost_private">Almost Private</option>
            </select>

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
    );
};

export default RegistrationForm;
