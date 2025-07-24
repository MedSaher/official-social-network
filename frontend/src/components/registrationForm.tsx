'use client';
import { useState } from "react";
import { RegistrationFormData } from '../lib/types';
import axios from 'axios';
import './component.css/RegistrationForm.css'; // 👈 Import the CSS file

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
        gender: '',
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
        <form onSubmit={handleSubmit} className="form-container">
            <h2>Register</h2>

            <input type="email" name="email" placeholder="Email" value={formData.email}
                onChange={handleChange} className="form-input" required />
            {errors.email && <span className="form-error">{errors.email}</span>}

            <input type="password" name="password" placeholder="Password" value={formData.password}
                onChange={handleChange} className="form-input" required />
            {errors.password && <span className="form-error">{errors.password}</span>}

            <input type="text" name="firstName" placeholder="First Name" value={formData.firstName}
                onChange={handleChange} className="form-input" required />
            {errors.firstName && <span className="form-error">{errors.firstName}</span>}

            <input type="text" name="lastName" placeholder="Last Name" value={formData.lastName}
                onChange={handleChange} className="form-input" required />
            {errors.lastName && <span className="form-error">{errors.lastName}</span>}

            <input type="date" name="dateOfBirth" value={formData.dateOfBirth}
                onChange={handleChange} className="form-input" required />
            {errors.dateOfBirth && <span className="form-error">{errors.dateOfBirth}</span>}

            <label className="form-label">Gender:</label>
            <select name="gender" value={formData.gender}
                onChange={handleChange} className="form-select" required>
                <option value="">Select Gender</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
            </select>
            {errors.gender && <span className="form-error">{errors.gender}</span>}

            <input type="file" name="avatar" accept="image/*"
                onChange={handleChange} className="form-input" />

            <input type="text" name="nickname" placeholder="Nickname (optional)"
                value={formData.nickname} onChange={handleChange} className="form-input" />

            <textarea name="aboutMe" placeholder="About Me (optional)" value={formData.aboutMe}
                onChange={handleChange} className="form-textarea" rows={4} />

            <label className="form-label">Profile Privacy:</label>
            <select name="privacyStatus" value={formData.privacyStatus}
                onChange={handleChange} className="form-select">
                <option value="public">Public</option>
                <option value="private">Private</option>
                <option value="almost_private">Almost Private</option>
            </select>

            <button type="submit" disabled={loading} className="form-button">
                {loading ? 'Submitting...' : 'Register'}
            </button>

            <p className="login-link">
                You have an account? <a href="/login">Login</a>
            </p>
        </form>
    );
};

export default RegistrationForm;
