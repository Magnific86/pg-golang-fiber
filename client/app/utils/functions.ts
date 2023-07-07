import axios from "axios"
import { IPost } from "./types"
import { API_BASE_URL } from "./base_url"

export const getPosts = async () => {
  try {
    const resp = await axios.get<IPost[]>(`${API_BASE_URL}/posts`)
    console.log("getData response: ", resp)
    return resp
  } catch (e) {
    return e
  }
}

export const createPost = async obj => {
  const config = {
    headers: {
      // "Content-Type": "multipart/form-data",
      "Content-Type": "application/json",
    },
  }
  // const formData = new FormData()
  // formData.append("title", obj.title)
  // formData.append("content", obj.content)
  // formData.append("file", obj.file)

  const body = JSON.stringify(obj)

  try {
    return await axios.post(`${API_BASE_URL}/create_post`, body, config)
  } catch (e) {
    return e
  }
}

export const getCurrentPost = async (id: string) => {
  try {
    return await axios.get<IPost>(`${API_BASE_URL}/posts/${id}`)
  } catch (e) {
    return e
  }
}

export const deletePost = async (id: string) => {
  const config = {
    headers: {
      // "Access-Control-Allow-Origin' ": "*",
      "Content-Type": "application/json",
      // "Access-Control-Request-Method": "DELETE",
      // "Access-Control-Allow-Methods": "GET, PUT, POST, DELETE",
      // "Access-Control-Allow-Headers": "X-Requested-With, content-type",
    },
  }
  try {
    return await axios.delete(`${API_BASE_URL}/delete_post/${id}`, config)
  } catch (e) {
    return e
  }
}

export const clearFileInput = ctrl => {
  try {
    ctrl.value = null
  } catch (ex) {}
  if (ctrl.value) {
    ctrl.parentNode.replaceChild(ctrl.cloneNode(true), ctrl)
  }
}
