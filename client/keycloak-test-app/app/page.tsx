/* eslint-disable @next/next/no-async-client-component */
import { getToken } from "next-auth/jwt";
import MainPage from "./components/mainPage";
import { cookies } from 'next/headers'
import { getSession } from "next-auth/react";
import { getServerSession } from "next-auth";
// import { getToken } from "next-auth/jwt"
// import { getServerSession } from "next-auth";
// import { getSession } from "next-auth/react";
import { authOptions } from "./api/auth/[...nextauth]/route";

export default async function Home() {

  // Treat this as server component

  // Works in server component
  // Need to get access token this way, can't do it through cookies unless you append all session token values
  const session = await getServerSession(authOptions)

  const serverData = await fetch("http://localhost:5050/", {
    headers: {
      Cookie: `token=${session?.access_token};`
    },
    cache: "no-cache",
  });
  //const serverDataJson = await serverData.json()
  const serverDataJson = await serverData.text()
  console.log(serverDataJson)

  return (
    <MainPage dataProp={serverDataJson} />
  )
}


// const [serverResponse, setServerResponse] = useState("");


// // Grab Nasa data (backend limits to 50 entries for demo)
// useEffect(()=>{
//   const fetchServerData = async () => {
//     const serverData = await fetch("http://localhost:5050/", {
//       credentials: "include",
//     });
//     //const serverDataJson = await serverData.json()
//     const serverDataJson = await serverData.text()
//     console.log(serverDataJson)
//     setServerResponse(serverDataJson)
//   }

//   fetchServerData()

// }, [])


// return (
//   <div className="p-4 inline-block">
//     <h1 className="text-3xl font-bold pb-2">
//       KeyCloak Testing Ground
//     </h1>
//     <hr/>
//     <div>
//       Data from Server: {serverResponse}
//     </div>
//   </div>
// )