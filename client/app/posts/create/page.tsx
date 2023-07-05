"use client"

import { createPost } from "@/app/utils/functions"
import { IPost } from "@/app/utils/types"
import { ChangeEvent, FormEvent, useState } from "react"
import { toast } from "react-toastify"

const CreatePost = () => {
  const emptyPost = {
    id: "",
    title: "",
    content: "",
    file: null,
  }

  //   const getPostsHandler = () => {
  //     getPosts().then(result => {
  //       if (result.status === 200) {
  //         setPosts(result.data.data)
  //       }
  //     })
  //   }

  const [currentPost, setCurrentPost] = useState<IPost>(emptyPost)
  const [file, setFile] = useState<File | null>(null)

  const changedataHandler = (key: string, value: string) => {
    setCurrentPost({ ...currentPost, [key]: value })
  }

  const createPostHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    createPost({ title: currentPost.title, content: currentPost.content, file }).then(result => {
      if (result.status == 200) {
        // getPostsHandler()
        setCurrentPost(emptyPost)
      } else {
        toast.error("filed to create new post")
      }
    })
  }

  const onUploadFileHandler = (e: ChangeEvent<HTMLInputElement>) => {
    const files = (e.target as HTMLInputElement).files
    if (files) {
      console.log("onUploadFileHandler: ", files[0])
      setFile(files[0])
    }
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
        <input type="file" name="file" onChange={e => onUploadFileHandler(e)} accept="image/*" />
        <input type="submit" value="submit" />
      </form>
    </div>
  )
}
export default CreatePost
