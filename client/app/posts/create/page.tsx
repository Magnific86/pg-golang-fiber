"use client"

import { createPost } from "@/app/utils/functions"
import { IPost } from "@/app/utils/types"
import { FormEvent, useState } from "react"

const CreatePost = () => {
  const emptyPost = {
    id: "",
    title: "",
    content: "",
  }

  //   const getPostsHandler = () => {
  //     getPosts().then(result => {
  //       if (result.status === 200) {
  //         setPosts(result.data.data)
  //       }
  //     })
  //   }

  const [currentPost, setCurrentPost] = useState<IPost>(emptyPost)

  const changedataHandler = (key: string, value: string) => {
    setCurrentPost({ ...currentPost, [key]: value })
  }

  const createPostHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    createPost({title: currentPost.title, content: currentPost.content}).then(result => {
      if (result.status == 200) {
        // getPostsHandler()
        setCurrentPost(emptyPost)
      }
    })
  }

  return (
    <div className="">
      <form onSubmit={createPostHandler}>
        <label>
          title
          <input className="my_input" type="text" value={currentPost.title} onChange={e => changedataHandler("title", e.target.value)} />
        </label>
        <label>
          content
          <input
            className="my_input"
            type="text"
            value={currentPost.content}
            onChange={e => changedataHandler("content", e.target.value)}
          />
        </label>
        <input type="submit" value="submit" />
      </form>
    </div>
  )
}
export default CreatePost
