// lib/types.ts
export interface RegistrationFormData {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  dateOfBirth: string;
  avatar: File | null;
  nickname: string;
  aboutMe: string;
  privacyStatus: 'public' | 'private' | 'almost_private'; 
  gender: 'male' | 'female' | 'other' | '';
}

// frontend/src/lib/types.ts

export interface ProfileType {
  id: string;
  avatar: string;
  firstName: string;
  lastName: string;
  nickname: string;
  about: string;
  email: string;
  followers: number;
  following: number;
  accountType: "public" | "private";
  dateOfBirth: string;
}

export interface PostType {
  id: string;
  content: string;
  createdAt: string;
  // Add more fields as needed from your API
}

export interface UserListType {
  id: string;
  firstName: string;
  lastName: string;
  avatar: string;
}
