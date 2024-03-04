'use client'

import { useEffect, useState } from "react";

export default function MainPage(props : any){

    const [serverResponse, setServerResponse] = useState("");

    useEffect(()=>{
        setServerResponse(props.dataProp)
    }, [props])
  
  
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