// "use client"

import Link from "next/link"
import { useRouter } from "next/navigation"
import { Suspense } from "react"

export default function page() {
  // const router = useRouter()

  return (
    // <Suspense fallback={null}>
    <div className="bg-red-500">
      {/* <button onClick={() => router.back()}>back</button> */}
      <Link href={"/"}>back </Link>
      about page
    </div>
    // </Suspense>
  )
}
