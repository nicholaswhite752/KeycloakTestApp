import { JWT, getToken } from 'next-auth/jwt';
import { NextRequest, NextResponse } from "next/server";


export async function middleware(req: NextRequest) {
    // validate the user is authenticated
    const verifiedToken = await getToken({req})
  
    // redirect if the token is invalid
    if (!verifiedToken) {
      return NextResponse.redirect('http://localhost:3000/api/auth/signin');
    }

    return NextResponse.next();
}

// TODO: Should add a route matcher to not match to login page
export const config = {
    matcher: ['/:path?'],
  }