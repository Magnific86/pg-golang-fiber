import Link from "next/link"
import "./globals.css"
import { ToastContainer } from "react-toastify"

export const meta = <title>my title name</title>

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>{meta}</head>
      <body>
        <nav className="gap-20">
          <Link href={"/"}>Main</Link>
          <Link href={"/about"}>About</Link>
          <Link href={"/posts"}>All posts</Link>
          <Link href={"/posts/create"}>Create new post</Link>
        </nav>
        {children}
        <ToastContainer />
      </body>
    </html>
  )
}
