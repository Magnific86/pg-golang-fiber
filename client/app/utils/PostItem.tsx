import { FC } from "react"
import Link from "next/link"
import { deletePost } from "./functions"
import { BASE_URL } from "./base_url"
import Image from "next/image"
import { Blob } from "buffer"

interface IPostProps {
  id: string
  title: string
  content: string
  file: Blob
}

export const PostItem: FC<IPostProps> = ({ id, title, content, file }) => {
  const deletePostHandler = () => {
    deletePost(id).then(result => {
      if (result.status == 200 && window) {
        window.location.href = `${BASE_URL}/posts`
      }
    })
  }

  return (
    <div className="post_container">
      <Link href={`/posts/${id}`}>{title}</Link>
      <h3>{content}</h3>
      {/* @ts-ignore */}
      <img src={() => URL.createObjectURL(file)} />
      <button onClick={deletePostHandler}>delete</button>
    </div>
  )
}
