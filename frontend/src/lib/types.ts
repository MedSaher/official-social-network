export interface RegistrationFormData {
    email: string;
    password: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string; // ISO format: YYYY-MM-DD
    avatar?: File | null;
    nickname?: string;
    aboutMe?: string;
    privacyStatus: 'public' | 'private';
}