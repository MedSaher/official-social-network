// types/PostForm.ts

export interface PostForm {
  title: string
  content: string
  image_path?: string         // optional (nullable in DB)
  privacy?: 'public' | 'almost_private' | 'private'  // defaults to 'public'
  group_id?: number           // optional (nullable)
  categories?: number[]       // optional: category IDs
}
