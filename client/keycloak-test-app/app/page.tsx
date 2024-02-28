'use client'
import { getSession, signIn, signOut } from "next-auth/react"

export default function Home() {

  // const session = await getSession()
  // console.log(session)


  return (
    <div className="p-4 inline-block">
      <h1 className="text-3xl font-bold pb-2">
        KeyCloak Testing Ground
      </h1>
      <hr/>
    </div>
  )
}
