// types/PostForm.ts


export interface PostFormData {
  title: string;
  content: string;
  imageUrl: string | null; // URL from separate upload
  privacy: 'public' | 'almost_private' | 'private';
  group_id?: number | '';
  categories: number[]; // multiple category ids
}

export interface Props {
  categoriesList: { id: number; name: string }[]; // categories options to select from
  groupsList?: { id: number; name: string }[]; // optional groups list
}
