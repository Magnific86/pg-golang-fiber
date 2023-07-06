"use client"

import { useEffect, useState } from "react"
import { IPostWithId } from "../utils/types"
import { getPosts } from "../utils/functions"
import { PostItem } from "../utils/PostItem"

const AllPostsPage = () => {
  const [posts, setPosts] = useState<IPostWithId[]>([])

  const getPostsHandler = () => {
    getPosts().then(result => {
      if (result.status === 200) {
        setPosts(result.data)
      }
    })
  }

  useEffect(() => {
    getPostsHandler()
  }, [])

  return (
    <div>
      {posts &&
        posts?.length > 0 &&
        posts.map(({ id, title, content }, index) => <PostItem id={id.toString()} key={index} title={title} content={content} />)}
    </div>
  )
}
export default AllPostsPage
