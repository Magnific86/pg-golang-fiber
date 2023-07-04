import { NextApiRequest, NextApiResponse } from "next"
import { NextResponse } from "next/server"
import { headers } from "next/headers"
import NextCors from "nextjs-cors"

const respData = {
  name: "Name",
  age: "20",
}

// export async function GET(req: NextApiRequest, res: NextApiResponse) {
//   const headerList = headers()
//   // await NextCors(req, res, {
//   //   methods: ["GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"],
//   //   origin: "*",
//   //   optionsSuccessStatus: 200,
//   // })
//   return await NextResponse.json({ respData })
// }

export async function GET(request: Request) {
  return new Response(JSON.stringify(respData), {
    status: 200,
    headers: {
      "Access-Control-Allow-Origin": "http://localhost:4000",
      "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type, Authorization",
    },
  })
}
