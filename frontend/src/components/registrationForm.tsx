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
        <form onSubmit={handleSubmit}>
            {/* Inputs as before */}
            <input
                type="email"
                name="email"
                placeholder="Email"
                value={formData.email}
                onChange={handleChange}
            />
            {errors.email && <span>{errors.email}</span>}

            <input
                type="password"
                name="password"
                placeholder="Password"
                value={formData.password}
                onChange={handleChange}
            />
            {errors.password && <span>{errors.password}</span>}

            <input
                type="text"
                name="firstName"
                placeholder="First Name"
                value={formData.firstName}
                onChange={handleChange}
            />
            {errors.firstName && <span>{errors.firstName}</span>}

            <input
                type="text"
                name="lastName"
                placeholder="Last Name"
                value={formData.lastName}
                onChange={handleChange}
            />
            {errors.lastName && <span>{errors.lastName}</span>}

            <input
                type="date"
                name="dateOfBirth"
                value={formData.dateOfBirth}
                onChange={handleChange}
            />
            {errors.dateOfBirth && <span>{errors.dateOfBirth}</span>}

            <input type="file" name="avatar" accept="image/*" onChange={handleChange} />

            <input
                type="text"
                name="nickname"
                placeholder="Nickname (optional)"
                value={formData.nickname}
                onChange={handleChange}
            />

            <textarea
                name="aboutMe"
                placeholder="About Me (optional)"
                value={formData.aboutMe}
                onChange={handleChange}
            />

            <label>Profile Privacy:</label>
            <select
                name="privacyStatus"
                value={formData.privacyStatus}
                onChange={handleChange}
            >
                <option value="public">Public</option>
                <option value="private">Private</option>
            </select>

            <button type="submit" disabled={loading}>
                {loading ? 'Submitting...' : 'Register'}
            </button>
        </form>
    )

}

export default RegistrationForm;