/* eslint-disable @next/next/no-async-client-component */
'use client'

import { getSession } from "next-auth/react";
import { useEffect, useState } from "react";

export default function Home() {


  const [serverResponse, setServerResponse] = useState("");


  // Grab Nasa data (backend limits to 50 entries for demo)
  useEffect(()=>{
    const fetchServerData = async () => {
      const serverData = await fetch("http://localhost:5050/", {
        credentials: "include",
      });
      //const serverDataJson = await serverData.json()
      const serverDataJson = await serverData.text()
      console.log(serverDataJson)
      setServerResponse(serverDataJson)
    }

    fetchServerData()

  }, [])


  return (
    <div className="p-4 inline-block">
      <h1 className="text-3xl font-bold pb-2">
        KeyCloak Testing Ground
      </h1>
      <hr/>
      <div>
        Data from Server: {serverResponse}
      </div>
    </div>
  )
}
