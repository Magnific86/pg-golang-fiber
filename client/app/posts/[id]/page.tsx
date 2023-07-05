"use client"

import { PostItem } from "@/app/utils/PostItem"
import { getCurrentPost } from "@/app/utils/functions"
import { IPost } from "@/app/utils/types"
import { useParams } from "next/navigation"
import React, { useEffect, useState } from "react"

const PostItemPage = () => {
  const [post, setPost] = useState<IPost>()
  const params = useParams()

  useEffect(() => {
    getCurrentPost(params?.id).then(result => {
      if (result.status == 200) {
        setPost(result.data.data)
      }
    })
  }, [])

  console.log("params: ", params)

  return (
    <>
      {post?.content && post.title && params.id ? (
        <PostItem id={params.id} title={post?.title} content={post?.content} file={post.file} />
      ) : (
        <p>failed to load post</p>
      )}
    </>
  )
}
export default PostItemPage
