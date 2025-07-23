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

