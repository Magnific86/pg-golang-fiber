import { Blob } from "buffer"

export interface IPost {
  id: string
  title: string
  content: string
  file: Blob | null
}
